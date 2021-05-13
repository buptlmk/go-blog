package redis

import (
	"blog/log"
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var RDB *redis.Client
var ctxRedis = context.Background()

func Close() error {
	err := RDB.Close()
	return err
}

func InitRedis() error {

	//RDB = redis.NewClient(&redis.Options{
	//	Addr: config.Settings.RedisSet.Addr,
	//	//Addr: "10.106.128.113:6379",
	//	Password: config.Settings.RedisSet.Password,
	//	//Password: "",
	//	DB: config.Settings.RedisSet.DB,
	//	//DB:0,
	//})

	RDB = redis.NewClient(&redis.Options{
		Addr:     "10.106.128.113:6379",
		Password: "",
		DB:       0,
	})

	_, err := RDB.Ping(ctxRedis).Result()
	//log.Logger.Error(err)
	return err
}

// 注意使用json.Marshal 的struct 属性必须是可以导出
func Set(key string, value interface{}, expiration time.Duration) error {
	statusCmd := RDB.Set(ctxRedis, key, value, expiration)
	log.Logger.Info(statusCmd.String())
	return statusCmd.Err()
}

// 取struct 利用json.unmarshal转化
func Get(key string) (string, error) {
	statusCmd := RDB.Get(ctxRedis, key)
	log.Logger.Info(statusCmd.String())
	if statusCmd == nil {
		return "", errors.New("data is no exist")
	}
	return statusCmd.Result()
}

// TODO: add hash set
func RPush(key string, value ...interface{}) error {
	cmd := RDB.RPush(ctxRedis, key, value)
	err := cmd.Err()
	if err != nil {
		log.Logger.Error(cmd.String())
	}
	log.Logger.Info(cmd.String())
	return err
}

func LPop(key string) (string, error) {

	cmd := RDB.LPop(ctxRedis, key)
	err := cmd.Err()
	if err != nil {
		log.Logger.Error(cmd.String())
	}
	log.Logger.Info(cmd.String())
	return cmd.Result()
}

func ZAdd(key string, value ...ZStruct) error {
	z := make([]*redis.Z, len(value))
	for i := 0; i < len(value); i++ {
		z[i] = &redis.Z{
			Score:  value[i].Score,
			Member: value[i].Member,
		}

	}
	cmd := RDB.ZAdd(ctxRedis, key, z...)
	log.Logger.Info(cmd.String())
	return cmd.Err()
}
func ZRevRangeWithScores(key string, start int, stop int) ([]ZStruct, error) {
	cmd := RDB.ZRevRangeWithScores(ctxRedis, key, int64(start), int64(stop))
	log.Logger.Info(cmd.String())
	z, err := cmd.Result()
	if err != nil {
		return nil, err
	}
	res := make([]ZStruct, len(z))
	for i := 0; i < len(z); i++ {
		res[i] = ZStruct(z[i])
	}
	return res, nil
}

func SetExpiration(key string, duration time.Duration) error {
	cmd := RDB.Expire(ctxRedis, key, duration)

	log.Logger.Info(cmd.String())
	return cmd.Err()
}

// TODO:分布式锁

// TODO:事务
// redis的事务不具有原子性
func TxPipeline() {
}

// 测试需要改动InitRedis的配置，因为没有加载config.json
func test() {
	err := InitRedis()
	if err != nil {
		log.Logger.Error(err)
		panic(err)
	}
	defer Close()

	err = Set("key", 22, 0)
	s, err := Get("2233")
	fmt.Println(s, err)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
