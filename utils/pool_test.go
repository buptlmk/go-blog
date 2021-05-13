package utils

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestGoroutinePool_Put(t *testing.T) {
	InitGoroutinePool(10)

	wg := new(sync.WaitGroup)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		task := &Task{
			Handler: func(v ...interface{}) {
				wg.Done()

				time.Sleep(1 * time.Second)
				fmt.Println(v)

			},
			Parameters: []interface{}{i, i * 2, "hello"},
		}
		go GPool.Put(task)
	}

	wg.Wait()
	Close()
	time.Sleep(1 * time.Second)
	println(GPool.size)
}
