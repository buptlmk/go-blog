package chat

import (
	"blog/utils"
	"context"
	"errors"
	"github.com/google/uuid"
	"sync"
	"time"
)

const BufferSize int = 100

type id = string

// Event 复用
var Pool = sync.Pool{New: func() interface{} {
	return &Event{}
}}

// 协程池
var _ = utils.InitGoroutinePool(1000)

var ctx, cancelFunc = context.WithCancel(context.Background())

type Room struct {
	locker      sync.RWMutex
	Name        string
	Number      int
	MessageChan chan *Event
	Users       *sync.Map // 保存 map[id]chan Event
	//Users map[id]chan Event	// 传给用户的是值，因为不确定什么时候用户才能读
	// 房间是否关闭
	Close bool
}

func NewRoom(name string) *Room {
	r := &Room{
		Name:        name,
		MessageChan: make(chan *Event, BufferSize),
		//MessageChan: chanEvent,
		Users: new(sync.Map),
		//Users: make(map[id]chan Event),
	}

	go r.Serve(ctx)
	return r
}

func CloseRoom(r *Room) {

	// 关闭服务
	cancelFunc()
	r.locker.Lock()
	defer r.locker.Unlock()
	r.Close = true

	// 编码时会优化成mapclear,很快，
	// 也可直接r.Users = make(map[uid]chan Event),交给GC
	r.Users.Range(func(key, value interface{}) bool {
		r.Users.Delete(key)
		return true
	})

	//for k := range r.Users {
	//	delete(r.Users, k)
	//}

	// 将Message通道清空,否则交给GC吧
	//LooP:
	//	select {
	//	case <-r.MessageChan:
	//		goto LooP
	//	default:
	//		Pool.Put(r.MessageChan)
	//	}
	return
}

func (r *Room) Join(name string) (*Person, error) {
	if r.Close {
		return nil, errors.New("room is not existed")
	}
	chanEvent := make(chan Event, BufferSize)
	// 发送事件告诉别人你上线了
	uid := uuid.New().String()
	event := Pool.Get().(*Event)
	event.UserId = uid
	event.User = name
	event.Timestamp = time.Now()
	event.Text = name + " come in"
	event.Type = JoinEvent

	r.MessageChan <- event
	if !r.Close {

		//r.Users[uid] = chanEvent
		r.Users.Store(uid, chanEvent)
		r.Number++
		//fmt.Println(uid)
	} else {
		return nil, errors.New("room is no existed")
	}

	return &Person{
		Name:        name,
		ID:          uid,
		ReceiveChan: chanEvent,
		SendChan:    r.MessageChan,
	}, nil

}

func (r *Room) Serve(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case event := <-r.MessageChan:
			task := &utils.Task{
				Handler:    broadCast2,
				Parameters: []interface{}{r, event},
			}
			//go broadCast(r, event) // 暂时这样
			utils.GPool.Put(task)
		}

	}
}

func broadCast2(v ...interface{}) {
	r := v[0].(*Room)
	e := v[1].(*Event)
	switch e.Type {
	case QuitEvent:
		//delete(r.Users, e.UserId)
		r.Users.Delete(e.UserId)

	default:
		break
	}

	r.Users.Range(func(key, value interface{}) bool {
		v := value.(chan Event)
		v <- *e
		return true
	})

	//r.locker.RLock()
	//for _, v := range r.Users {
	//	v <- *e
	//	//fmt.Println(r.Number)
	//}
	//r.locker.RUnlock()

	// 不再对e进行清空了
	Pool.Put(e)
}

func broadCast(r *Room, e *Event) {

	switch e.Type {
	case QuitEvent:
		r.Users.Delete(e.UserId)
	default:
		break
	}

	r.Users.Range(func(key, value interface{}) bool {
		v := value.(chan Event)
		v <- *e
		return true
	})
	// 不再对e进行清空了
	Pool.Put(e)
}
