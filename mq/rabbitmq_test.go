package mq

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestA(t *testing.T) {
	err := NewRabbitMq("test111", "", "")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("okkk")
	go func() {
		for i := 0; i < 10; i++ {

			//s := SS{
			//	"lmk",
			//}
			//ee,err := json.Marshal(s)
			//err = RabMq.PublishFanOut(ee)
			err = RabMq.PushDirect([]byte("lmk"))
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println("ok" + strconv.Itoa(i))
			time.Sleep(1 * time.Second)
		}
	}()

	//for i:=0;i<10;i++{
	//	err = RabMq.PublishFanOut("lmk")
	//	err = RabMq.PushDirect("lmk")
	//	if err!=nil{
	//		fmt.Println(err.Error())
	//	}
	//	fmt.Println("ok"+strconv.Itoa(i))
	//}

	//err = RabMq.PushDirect("lmk")

	//go func() {
	//	RabMq.ConsumeFanOut()
	//}()

	fmt.Println("ok")
	RabMq.ConsumeDirect()
	//RabMq.ConsumeFanOut()
	time.Sleep(10 * time.Second)
}

func TestB(t *testing.T) {
	p, err := NewProducer("lmk")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		c, err := NewConsumer(strconv.Itoa(i), p.Exchange)
		if err != nil {
			fmt.Println("new consumer err")
			continue
		}

		go c.Consume(p.Exchange)
	}

	for i := 0; i < 10; i++ {
		p.Publish("asf" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
	}
	wg.Wait()
}
