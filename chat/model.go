package chat

import (
	//"blog/chat/map"
	"context"
	"fmt"
	"time"
)

const (
	JoinEvent = iota
	MessageEvent
	QuitEvent
)

type Event struct {
	Type      int       `json:"type"`
	Timestamp time.Time `json:"timestamp"`
	User      string    `json:"user"`
	UserId    string    `json:"user_id"`
	Text      string    `json:"text"`
}

func NewEvent(infoType int, user string, userId string, text string) Event {

	return Event{
		Type:      infoType,
		Timestamp: time.Now(),
		User:      user,
		UserId:    userId,
		Text:      text,
	}
}

type Person struct {
	Name        string        `json:"name"`
	ID          id            `json:"id"`
	ReceiveChan <-chan Event  `json:"receive_chan"`
	SendChan    chan<- *Event `json:"send_chan"`
}

func (p *Person) Leave() {
	event := Pool.Get().(*Event)
	event.Type = QuitEvent
	event.User = p.Name
	event.UserId = p.ID
	event.Timestamp = time.Now()
	event.Text = p.Name + " leave this room"
	p.SendChan <- event
}

func (p *Person) Say(text string) {

	event := Pool.Get().(*Event)
	event.Type = MessageEvent
	event.User = p.Name
	event.UserId = p.ID
	event.Timestamp = time.Now()
	event.Text = text
	p.SendChan <- event
}

func (p *Person) Get(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case e := <-p.ReceiveChan:

			s := fmt.Sprintf("%s--->%s说:%s", e.Timestamp.Format("2006-01-02 15:04:05"), e.User, e.Text)
			//log.Logger.WithField(p.Name,s)
			fmt.Println(p.Name, ":: ", s)

		}
	}
}
