package pool

import (
	"errors"
	"fmt"
	"io"
	"sync"
)

// 资源复用池，没有maxsize 限制，没有回收功能
// 最大链接的处理需要一个queue来维持排队
// 归还链接时，如果有排队的直接给排队的等等...
// 具体方法参照sql标准库、go-redis的实现
type MultiPool struct {
	sync.Mutex
	factory  func() (io.Closer, error)
	resource chan io.Closer
	closed   bool
}

// size: max buffer
func New(f func() (io.Closer, error), size int) (*MultiPool, error) {

	return &MultiPool{
		factory:  f,
		resource: make(chan io.Closer, size),
		closed:   false,
	}, nil
}

func (p *MultiPool) Acquire() (io.Closer, error) {

	select {
	case r, ok := <-p.resource:
		if ok {
			return r, nil
		}
		return nil, errors.New("closed")
	default:
		return p.factory()
	}

}

func (p *MultiPool) Close() {
	p.Lock()
	defer p.Unlock()

	// 先判断
	if p.closed {
		return
	}

	p.closed = true
	// 关闭通道，不写了
	close(p.resource)
	for v := range p.resource {
		v.Close()
	}
}

// 释放资源，资源满的情况下直接关闭
func (p *MultiPool) Release(r io.Closer) {
	p.Lock()
	defer p.Unlock() // 好像不用枷锁
	if p.closed {
		r.Close()
		return
	}
	select {
	//
	case p.resource <- r:
		fmt.Println("success put into pool")
	default:
		fmt.Println("资源已满，直接释放")
		r.Close()
	}
}
