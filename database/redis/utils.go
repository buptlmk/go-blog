package redis

import (
	"blog/mq"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var local = LocalSpike{
	LocalInStock:     150,
	LocalSalesVolume: 0,
}
var remote = RemoteSpikeKeys{
	SpikeOrderHashKey:  "ticket_hash_key",
	TotalInventorKey:   "ticket_total_nums",
	QuantityOfOrderKey: "ticket_sold_nums",
}

var done = make(chan struct{}, 1)

func InitTicket(number int64) {

	local.LocalInStock = number
	local.LocalSalesVolume = 0
	done <- struct{}{}
	mq.NewRabbitMq("ticket", "", "")

	go consume()
}

func Pur(orderHash string) error {
	LogMsg := ""
	flag := false
	<-done
	if local.LocalDeductionStock() && remote.RemoteDeductionStock(orderHash) {
		LogMsg = "result:1,localSales:" + strconv.FormatInt(local.LocalSalesVolume, 10)
		flag = true
	} else {
		LogMsg = "result:0,localSales:" + strconv.FormatInt(local.LocalSalesVolume, 10)

	}
	done <- struct{}{}
	if flag {
		go push(orderHash)
	} else {
		fmt.Println("已售空或者订单已存在")
		return errors.New("failed")
	}
	go writeLog(LogMsg, "./stat.log")

	return nil

}

func pur(done chan struct{}, orderHash string) {
	LogMsg := ""
	flag := false
	<-done
	if local.LocalDeductionStock() && remote.RemoteDeductionStock(orderHash) {
		LogMsg = "result:1,localSales:" + strconv.FormatInt(local.LocalSalesVolume, 10)
		flag = true
	} else {
		LogMsg = "result:0,localSales:" + strconv.FormatInt(local.LocalSalesVolume, 10)

	}
	done <- struct{}{}
	if flag {
		push(orderHash)
	} else {
		fmt.Println("已售空或者订单已存在")
	}

	writeLog(LogMsg, "./stat.log")
}

func writeLog(msg string, logPath string) {
	fd, _ := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer fd.Close()
	content := strings.Join([]string{msg, "\r\n"}, "")
	buf := []byte(content)
	fd.Write(buf)
}

func push(order string) {
	//mq.PushDirect([]byte(order),"ticket")
	mq.RabMq.PushDirect([]byte(order))
}

func consume() {
	ticker := time.Tick(1 * time.Second)
	for range ticker {
		mq.RabMq.ConsumeDirect()
		//mq.ConsumeDirect("ticket")
	}
}
