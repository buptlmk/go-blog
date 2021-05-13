package redis

import (
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
	"time"
)

// 流程-> 生成订单，减库存，用户支付
// 保证不少卖，不超卖
// 1、下单减库存，频繁io，恶意买导致少卖
//2、支付减库存，很容易用户的订单支付不了，超卖，也容易并发io
//3、预扣库存:用户下单，扣库存，异步生成订单；用户下了单不支付设置有效期；订单的生成是异步的，放到MQ中处理

// 单机并发下，扣库存也不能对数据库操作，本地减库存

// nginx 平均单机并发量，每台机器保存一部分；若某台挂掉则少卖，所以做容错方案，本地减库存，远程也要减，redis 10W并发很适合
// 假如1000件，redis必须设置1000件，10台服务器每台可以设置100+buffer,具体多少可进一步考虑redis承受能力

type LocalSpike struct {
	LocalInStock     int64
	LocalSalesVolume int64
}

func (spike *LocalSpike) LocalDeductionStock() bool {
	spike.LocalSalesVolume = spike.LocalSalesVolume + 1
	return spike.LocalSalesVolume <= spike.LocalInStock
}

func (spike *LocalSpike) LocalAddStock() {
	spike.LocalSalesVolume -= 1
}

type RemoteSpikeKeys struct {
	SpikeOrderHashKey  string // 秒杀订单hash结构key
	TotalInventorKey   string // hash中总订单kucun
	QuantityOfOrderKey string // 已有的订单数量
}

// redis使用lua脚本保证一致性

const LuaScript = `
        local ticket_key = KEYS[1]
        local ticket_total_key = ARGV[1]
        local ticket_sold_key = ARGV[2]
		local order_hash = ARGV[3]
		local exist = redis.call('HGET',ticket_key,order_hash)
		if exist then
			return false
		end
        local ticket_total_nums = tonumber(redis.call('HGET', ticket_key, ticket_total_key))
        local ticket_sold_nums = tonumber(redis.call('HGET', ticket_key, ticket_sold_key))
		-- 查看是否还有余票,增加订单数量,返回结果值
        if(ticket_total_nums > ticket_sold_nums) then
			redis.call('HINCRBY', ticket_key, ticket_sold_key, 1)
            return redis.call('HSET', ticket_key, order_hash, 1)
        end
        return false
`

func (r *RemoteSpikeKeys) RemoteDeductionStock(orderHash string) bool {

	lua := redis.NewScript(LuaScript)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	flag, err := lua.Run(ctx, RDB, []string{r.SpikeOrderHashKey}, r.TotalInventorKey, r.QuantityOfOrderKey, orderHash).Bool()
	if err != nil {
		return false
	}

	return flag
}
