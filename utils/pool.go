package utils

import (
	"context"
	"errors"
	"sync/atomic"
	"time"
)

type GoroutinePool struct {
	capacity int32
	size     int32

	//sync.Mutex
	task chan *Task
	//closeChan 	chan struct{}
}

type Task struct {
	Handler    func(v ...interface{})
	Parameters []interface{}
}

var GPool *GoroutinePool

// GPoolContext
var GPoolContext context.Context
var GPoolCancelFunc context.CancelFunc

func InitGoroutinePool(capacity int32) error {
	GPool = &GoroutinePool{
		capacity: capacity,
		size:     0,
		task:     make(chan *Task, capacity),
		//closeChan: make(chan struct{},0),
	}
	GPoolContext, GPoolCancelFunc = context.WithCancel(context.Background())

	return nil
}

func Close() {
	for len(GPool.task) != 0 {

		time.Sleep(1 * time.Millisecond)
	}
	GPoolCancelFunc()
	close(GPool.task)
}

func (p *GoroutinePool) Run() error {

	// 乐观锁，这样做而不是使用atomic.addint 避免了加锁，主要还是保持capacity不超
	for {
		oldValue := atomic.LoadInt32(&p.size)

		if oldValue == p.capacity {
			return errors.New("GG")
		}
		//fmt.Println(oldValue)
		ok := atomic.CompareAndSwapInt32(&p.size, oldValue, oldValue+1)
		//fmt.Println(ok)
		if ok {
			break
		}
	}
	//oldValue := atomic.LoadInt32(&p.size)
	//
	//if oldValue==p.capacity{
	//	return errors.New("GG")
	//}

	//fmt.Println(GPool.size)
	go func(ctx context.Context) {
		defer p.dec()
		for {
			select {
			//case _,ok:=<-p.closeChan:
			//	if !ok{
			//		fmt.Println("go over")
			//		p.dec()
			//		return
			//	}

			case <-ctx.Done():
				return
			case task, ok := <-p.task:
				if !ok {
					return
				}
				task.Handler(task.Parameters...)
			}
		}
	}(GPoolContext)

	return nil
}

func (p *GoroutinePool) Put(t *Task) error {

	err := p.Run()
	//if err!=nil{
	//	fmt.Println()
	//}
	p.task <- t
	return err
}

func (p *GoroutinePool) dec() {
	//for{
	//	oldValue := atomic.LoadInt32(&p.size)
	//	if oldValue==0{
	//		return
	//	}
	//	ok := atomic.CompareAndSwapInt32(&p.size,oldValue,oldValue-1)
	//	if ok {
	//		break
	//	}
	//}

	atomic.AddInt32(&p.size, -1)
}
