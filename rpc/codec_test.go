package rpc

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"strconv"
	"sync"
	"testing"
	"time"
)

//
type EchoService struct{}

func (service *EchoService) Echo(arg string, result *string) error {
	*result = arg
	return nil
}

func RegisterAndServerOnTcp() {
	err := rpc.Register(&EchoService{})
	if err != nil {
		log.Fatal("error registering", err)
		return
	}
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":12345")
	if err != nil {
		log.Fatal("error resolving ", err)
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("error accept", err)

		} else {
			rpc.ServeCodec(NewServerCodec(conn))
		}
	}

}

func Echo(arg string) (result string, err error) {
	var c *rpc.Client
	conn, err := net.Dial("tcp", ":12345")
	c = rpc.NewClientWithCodec(NewClientCodec(conn))
	defer c.Close()

	if err != nil {
		return "", err
	}
	//err = c.Call("EchoService.Echo", arg, &result) //通过类型加方法名指定要调用的方法
	//if err != nil {
	//	return "", err
	//}
	//return result, err

	done := c.Go("EchoService.Echo", arg, &result, nil)

	select {
	case _, err := <-done.Done:
		fmt.Println(err)
	}
	return result, err
}

func TestNewServerCodec(t *testing.T) {
	go RegisterAndServerOnTcp()

	time.Sleep(1 * time.Second)
	wg := new(sync.WaitGroup)
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			s, err := Echo(strconv.Itoa(i))
			fmt.Println(s, err)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
