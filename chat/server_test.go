package chat

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestCloseRoom(t *testing.T) {
	room := NewRoom("chat")

	ctx, cancelFunc := context.WithCancel(context.Background())

	a, _ := room.Join("a")
	go a.Get(ctx)
	time.Sleep(1 * time.Second)

	b, _ := room.Join("b")
	go b.Get(ctx)

	time.Sleep(1 * time.Second)
	a.Say("Hello b")
	b.Say("Hello a")

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(name string) {
			defer wg.Done()

			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
			x, err := room.Join(name)
			if err != nil {
				return
			}

			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)

			x.Say(name)
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
			x.Leave()

			time.Sleep(5 * time.Second)
			fmt.Println("-----?")

		}(strconv.Itoa(i))
	}

	wg.Wait()
	time.Sleep(10 * time.Second)

	cancelFunc()
}

func TestA(t *testing.T) {
	var lock sync.RWMutex
	m1 := make(map[int]int)
	m1[0] = 1221
	m1[1] = 1221
	m1[33] = 1221
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 500; i++ {
		go func() {
			defer wg.Done()
			lock.RLock()
			for _, _ = range m1 {
				//fmt.Println(k,v)
			}
			lock.RUnlock()
		}()
	}

	for i := 0; i < 500; i++ {
		go func() {
			defer wg.Done()
			lock.Lock()
			m1[22] = 55
			lock.Unlock()
		}()
	}
	wg.Wait()
}
