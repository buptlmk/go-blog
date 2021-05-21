package redis

import (
	"blog/database"
	"blog/log"
	"blog/utils"
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

type ActivityEntity struct {
	Id    int
	Name  string
	Total int
	Sold  int
	Done  chan struct{}
}

type ActivityMap struct {
	sync.RWMutex
	M map[int]*ActivityEntity
}

var activityMap = ActivityMap{
	M: make(map[int]*ActivityEntity),
}

// redis使用lua脚本保证一致性

// 2 订单已存在
// 0 出现了问题
// 1 成功
const LuaScript = `
        local ticket_key = KEYS[1]
        local ticket_total_key = ARGV[1]
        local ticket_sold_key = ARGV[2]
		local order_hash = ARGV[3]
		local exist = redis.call('HGET',ticket_key,order_hash)
		if exist then
			return 2
		end
        local ticket_total_nums = tonumber(redis.call('HGET', ticket_key, ticket_total_key))
        local ticket_sold_nums = tonumber(redis.call('HGET', ticket_key, ticket_sold_key))
		-- 查看是否还有余票,增加订单数量,返回结果值
        if(ticket_sold_nums < ticket_total_nums) then
			redis.call('HINCRBY', ticket_key, ticket_sold_key, 1)
            return redis.call('HSET', ticket_key, order_hash, 1)
        end
        return 0
`

func init() {
	go initActivityMap()
}

func initActivityMap() {
	time.Sleep(3 * time.Second)
	for {
		activities, err := database.GetAllActivity()
		if err != nil {
			log.Logger.Error(err.Error())
			time.Sleep(1 * time.Hour)
			continue
		}
		temp := make([]utils.Activity, 0, 16)
		activityMap.RLock()
		for i := 0; i < len(activities); i++ {
			if _, ok := activityMap.M[int(activities[i].Id)]; !ok {
				temp = append(temp, activities[i])
			}
		}
		activityMap.RUnlock()
		if len(temp) == 0 {
			time.Sleep(1 * time.Hour)
			continue
		}
		activityMap.Lock()
		for i := 0; i < len(temp); i++ {
			activityTemp := &ActivityEntity{
				Id:    int(temp[i].Id),
				Name:  temp[i].Name,
				Total: temp[i].Total,
				Sold:  temp[i].Total - temp[i].Res,
				Done:  make(chan struct{}, 1),
			}
			activityTemp.Done <- struct{}{}
			activityMap.M[int(temp[i].Id)] = activityTemp
		}
		activityMap.Unlock()
		log.Logger.Info(activityMap.M)
		time.Sleep(1 * time.Hour)

	}

}

func (a *ActivityEntity) Purchase() bool {
	if a.Sold < a.Total {
		a.Sold++
		return true
	}
	return false
}

func AddActivity(id int, name string, total int, sold int) (err error) {
	activity := &ActivityEntity{
		Id:    id,
		Name:  name,
		Total: total,
		Sold:  sold,
		Done:  make(chan struct{}, 1),
	}
	activity.Done <- struct{}{}

	activityMap.Lock()
	activityMap.M[id] = activity
	activityMap.Unlock()

	// 向redis 注册
	cmd := RDB.HSet(context.Background(), activity.Name, map[string]interface{}{"Total": activity.Total, "Sold": activity.Sold})
	log.Logger.Info(cmd.String())
	err = cmd.Err()
	if err != nil {
		log.Logger.Error(cmd.String())
	}
	return err
}

func DelActivity(id int) (err error) {

	activityMap.Lock()
	name := activityMap.M[id].Name
	delete(activityMap.M, id)
	activityMap.Unlock()

	// 向redis 申请删除
	//cmd := RDB.HDel(context.Background(),name,"Total","Sold")
	cmd := RDB.Del(context.Background(), name)
	log.Logger.Info(cmd.String())
	err = cmd.Err()
	if err != nil {
		log.Logger.Error(cmd.String())
	}
	return err
}

func remoteDeductionStock(orderHash string, name string) (state int, err error) {

	lua := redis.NewScript(LuaScript)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	state, err = lua.Run(ctx, RDB, []string{name}, "Total", "Sold", orderHash).Int()
	log.Logger.Info(state)
	if err != nil {
		log.Logger.Error(err.Error())
		return state, err
	}

	return state, err
}

func PurchaseTicket(id int, orderHash string) (err error) {
	activityMap.RLock()
	activity, ok := activityMap.M[id]
	log.Logger.Info(activity)
	if !ok {
		activityMap.RUnlock()
		return errors.New("activity is not exist")
	}
	activityMap.RUnlock()

	<-activity.Done

	defer func() {
		activity.Done <- struct{}{}
	}()

	if activity.Purchase() {

		state, err := remoteDeductionStock(orderHash, activity.Name)
		if err != nil {
			return err
		}

		switch state {
		case 0:
			activity.Sold--
			err = errors.New("redis failed")
		case 1:
			err = nil
		case 2:
			activity.Sold--
			err = errors.New("order is already exist")
		default:
			activity.Sold--
			err = errors.New("unknown error")
		}

		return err
	}

	return errors.New("NO ticket")
}
