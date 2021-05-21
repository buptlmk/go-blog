package chat

import (
	"blog/database"
	"blog/log"
	"blog/mq"
	"encoding/json"
	"errors"
	"github.com/streadway/amqp"
	"sync"
	"time"
)

//var Rooms = make(map[string]*mq.Rabbitmq)
var Rooms = &sync.Map{}

type Event struct {
	Type      int       `json:"type" binding:"required"`
	Timestamp time.Time `json:"timestamp"`
	User      string    `json:"user" binding:"required"`
	CardId    string    `json:"cardId" binding:"required"`
	Text      string    `json:"text" binding:"required"`
}

func addRoom(name string) (err error) {

	_, exist := Rooms.Load(name)
	if exist {
		return errors.New(name + " is already existed.")
	}
	p, err := mq.NewProducer(name)
	if err != nil {
		return err
	}
	log.Logger.Info(name + " is ready.")
	Rooms.Store(name, p)
	return nil

}

func init() {
	go func() {
		time.Sleep(3 * time.Second)

		for {
			rooms, err := database.GetAllChatRoom()
			if err != nil {
				time.Sleep(1 * time.Minute)
				continue
			}

			for i := 0; i < len(rooms); i++ {
				name := rooms[i].Name
				if _, ok := Rooms.Load(name); !ok {
					p, err := mq.NewProducer(name)
					if err != nil {
						log.Logger.Info(err)
						continue
					}
					log.Logger.Info(name + " is ready.")
					Rooms.Store(name, p)
				}
			}

			time.Sleep(1 * time.Second)

		}

	}()
}

func joinRoom(roomName string, queueName string) (err error) {

	_, err = mq.NewConsumer(roomName+queueName, roomName)
	if err != nil {
		return err
	}
	log.Logger.Info(queueName + " join into chat room: " + roomName)
	return nil
}

func exitRoom(roomName string, queueName string) (err error) {

	err = mq.DelConsumer(roomName + queueName)
	if err != nil {
		log.Logger.Error(err.Error())
		return err
	}
	log.Logger.Info(queueName + " quit chat room: " + roomName)
	return nil
}

func publish(roomName string, text Event) error {

	textByte, err := json.Marshal(text)
	if err != nil {
		return err
	}

	p, exist := Rooms.Load(roomName)
	if !exist {
		return errors.New("the room is not existed")
	}
	room := p.(*mq.Producer)

	if err := room.Channel.Publish(roomName, "", true, false, amqp.Publishing{
		Timestamp:   time.Now(),
		ContentType: "application/json",
		Body:        textByte,
	}); err != nil {
		return err
	}
	return nil
}

func Consume(roomName, queueName string) (event []Event, err error) {

	p, exist := Rooms.Load(roomName)
	if !exist {
		return nil, errors.New("the room is not existed")
	}
	room := p.(*mq.Producer)
	ch, err := room.Conn.Channel()

	if err != nil {
		return nil, err
	}
	defer ch.Close()

	ch.Qos(10, 0, false)
	deliveries, err := ch.Consume(roomName+queueName, "any", false, false, false, true, nil)
	if err != nil {
		log.Logger.Error(err.Error())
		return nil, err
	}
	log.Logger.Info("pull message")
	event = make([]Event, 0, 16)
	for i := 0; i < 3; i++ {
		select {
		case v, _ := <-deliveries:
			err := v.Ack(false)
			if err != nil {
				log.Logger.Error(err.Error())
			}
			temp := Event{}
			if e := json.Unmarshal(v.Body, &temp); e != nil {
				log.Logger.Error(e.Error())
			} else {
				event = append(event, temp)
			}
			time.Sleep(200 * time.Millisecond)
		default:
			//return event,nil
			time.Sleep(200 * time.Millisecond)
		}

	}
	return event, nil
}
