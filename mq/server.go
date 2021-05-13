package mq

import (
	"fmt"
	"github.com/streadway/amqp"
	"time"
)

//const mqUrl = "amqp://lmk:123456@10.106.128.113:5672/mq"

type Producer struct {
	Conn     *amqp.Connection
	Exchange string
	Channel  *amqp.Channel
}

type Consumer struct {
	Conn    *amqp.Connection
	Queue   string
	Channel *amqp.Channel
}

func NewProducer(name string) (*Producer, error) {

	conn, err := amqp.Dial(mqUrl)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	if err = channel.ExchangeDeclare(name, "fanout", true, false, false, true, nil); err != nil {
		return nil, err
	}

	p := &Producer{
		Conn:     conn,
		Exchange: name,
		Channel:  channel,
	}
	//
	return p, nil
}

func NewConsumer(name string, exchange string) (*Consumer, error) {

	conn, err := amqp.Dial(mqUrl)
	defer conn.Close()
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	defer channel.Close()
	if err != nil {
		return nil, err
	}

	if _, err := channel.QueueDeclare(name, true, false, false, true, nil); err != nil {

		return nil, err
	}

	if err := channel.QueueBind(name, name, exchange, true, nil); err != nil {
		return nil, err
	}

	c := &Consumer{
		Conn:    conn,
		Queue:   name,
		Channel: channel,
	}
	//

	return c, nil
}

func (p *Producer) Publish(text string) error {

	if err := p.Channel.Publish(p.Exchange, "", true, false, amqp.Publishing{
		Timestamp:   time.Now(),
		ContentType: "application/json",
		Body:        []byte(text),
	}); err != nil {
		return err
	}
	return nil
}

func (c *Consumer) Consume(exchange string) {

	deliveries, err := c.Channel.Consume(c.Queue, "any", false, false, false, true, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for {
		select {
		case v, ok := <-deliveries:
			if ok {
				if err := v.Ack(true); err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println(c.Queue + " ::: " + string(v.Body))
				}
			} else {
				fmt.Println("close")
			}

		}
	}

}
