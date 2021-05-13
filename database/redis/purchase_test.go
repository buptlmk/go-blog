package redis

//
//import (
//	"blog/mq"
//	"crypto/rand"
//	"crypto/sha256"
//	"encoding/hex"
//	"fmt"
//	"math/big"
//	"os"
//	"strconv"
//	"strings"
//	"sync"
//	"testing"
//	"time"
//)
//
//var local = LocalSpike{
//	LocalInStock: 150,
//	LocalSalesVolume: 0,
//}
//var remote = RemoteSpikeKeys{
//	SpikeOrderHashKey: "ticket_hash_key",
//	TotalInventorKey: "ticket_total_nums",
//	QuantityOfOrderKey: "ticket_sold_nums",
//}
//
//
//func TestB(t *testing.T){
//
//	fmt.Println(local)
//	mq.NewRabbitMq("ticket","","")
//	err := InitRedis()
//	if err!=nil{
//		os.Exit(1)
//	}
//
//	done := make(chan struct{},1)
//	done <- struct{}{}
//
//	var wg  sync.WaitGroup
//	wg.Add(10)
//
//	for i:=0;i<10;i++{
//		go func() {
//			defer wg.Done()
//			sha := sha256.New()
//			s,_ := rand.Int(rand.Reader,big.NewInt(1<<63-1))
//			sha.Write([]byte(s.String()))
//			hash := hex.EncodeToString(sha.Sum(nil))
//			//hash := "2020110220"
//			pur(done,hash)
//		}()
//	}
//
//	go consume()
//	wg.Wait()
//
//	time.Sleep(5*time.Second)
//
//}
//
//
//
//func pur(done chan struct{},orderHash string) {
//	LogMsg := ""
//	flag := false
//	<-done
//	if local.LocalDeductionStock()&&remote.RemoteDeductionStock(orderHash){
//		LogMsg = "result:1,localSales:" + strconv.FormatInt(local.LocalSalesVolume, 10)
//		flag = true
//	}else{
//		LogMsg = "result:0,localSales:" + strconv.FormatInt(local.LocalSalesVolume, 10)
//
//	}
//	done<- struct{}{}
//	if flag{
//		push(orderHash)
//	}else{
//		fmt.Println("已售空或者订单已存在")
//	}
//
//	writeLog(LogMsg,"./stat.log")
//}
//
//func writeLog(msg string, logPath string) {
//	fd, _ := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
//	defer fd.Close()
//	content := strings.Join([]string{msg, "\r\n"}, "")
//	buf := []byte(content)
//	fd.Write(buf)
//}
//
//func push(order string){
//	mq.RabMq.PushDirect([]byte(order))
//}
//
//func consume(){
//	mq.RabMq.ConsumeDirect()
//
//}
