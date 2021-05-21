package mq

import (
	"blog/config"
	"blog/log"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"os"
	"strings"
)

var rabbitmqCfg = config.Settings.RabbitMq

var mqUrl = fmt.Sprintf("amqp://%s:%s@%s/%s", rabbitmqCfg.User, rabbitmqCfg.Password, rabbitmqCfg.Addr, rabbitmqCfg.VirtualHost)

type Rabbitmq struct {
	conn    *amqp.Connection
	channel *amqp.Channel

	queueName string
	exchange  string
	key       string
	mqUrl     string
}

type SS struct {
	Name string
}

var RabMq *Rabbitmq

func NewRabbitMq(queue, exchange, key string) error {
	RabMq = &Rabbitmq{
		queueName: queue,
		exchange:  exchange,
		key:       key,
		mqUrl:     mqUrl,
	}
	var err error
	RabMq.conn, err = amqp.Dial(RabMq.mqUrl)
	//RabMq.conn.Config.ChannelMax=200
	if err != nil {
		log.Logger.Error(err.Error())
		return err
	}
	RabMq.channel, err = RabMq.conn.Channel()
	if err != nil {
		log.Logger.Error(err.Error())
		return err
	}
	_, err = RabMq.channel.QueueDeclare(
		RabMq.queueName,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Logger.Error(err.Error())
	}

	return nil
}

func Close() (err error) {
	err = RabMq.channel.Close()
	err = RabMq.conn.Close()
	return err
}

// direct

func (r *Rabbitmq) PushDirect(message []byte) (err error) {

	// 创建队列，有则跳过

	//ch,err := r.conn.Channel()
	//if err!=nil{
	//	log.Logger.Error(err.Error())
	//	return err
	//}
	// send message
	err = r.channel.Publish(
		r.exchange,
		r.queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		},
	)
	//err = ch.Publish(
	//	r.exchange,
	//	r.queueName,
	//	false,
	//	false,
	//	amqp.Publishing{
	//		ContentType: "text/plain",
	//		Body: message,
	//	},
	//)
	if err != nil {
		log.Logger.Error(err.Error())
	}
	return err
}

func PushDirect(message []byte, queueName string) (err error) {
	conn, err := amqp.Dial(mqUrl)
	defer conn.Close()
	if err != nil {
		log.Logger.Error(err.Error())
		return err
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Logger.Error(err.Error())
		return err
	}

	_, err = ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Logger.Error(err.Error())
	}

	// send message
	err = ch.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		},
	)
	if err != nil {
		log.Logger.Error(err.Error())
	}
	return err
}

func (r *Rabbitmq) PublishFanOut(message []byte) (err error) {

	err = r.channel.ExchangeDeclare(
		r.exchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println("1---->")
	}
	err = r.channel.Publish(
		r.exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		},
	)

	if err != nil {
		fmt.Println("2--->" + err.Error())
	}
	return err
}

func (r *Rabbitmq) ConsumeFanOut() {
	err := r.channel.ExchangeDeclare(
		r.exchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println("3---->")
	}

	q, err := r.channel.QueueDeclare(
		r.queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println("4---->")
	}

	err = r.channel.QueueBind(
		q.Name,
		"",
		r.exchange,
		false,
		nil,
	)
	if err != nil {
		fmt.Println("5---->")
	}

	msg, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println("6---->")
	}

	for res := range msg {
		var v SS
		err = json.Unmarshal(res.Body, &v)
		log.Logger.Info(v)
	}
	log.Logger.Info("over")
}

func (r *Rabbitmq) ConsumeDirect() {
	// 创建队列，有则跳过
	ch, err := r.conn.Channel()
	defer ch.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	q, err := ch.QueueDeclare(
		r.queueName,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Logger.Error(err.Error())
	}

	msg, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Logger.Error(err.Error())
	}

	//r.channel.Qos(1,0,false)

	fmt.Println("----->")

	for {
		select {
		case res := <-msg:
			fmt.Println(true)
			res.Ack(false)
			writeLog(string(res.Body), "./stat.txt")
		default:
			fmt.Println("over")
			return
		}
	}

	//for res := range msg{
	//	//fmt.Println("写到数据库",string(res.Body))
	//	//log.Logger.Info(res.Body)
	//	//res.Ack(false)
	//	writeLog(string(res.Body),"./stat.txt")
	//
	//}
	//log.Logger.Info("over")
}

func ConsumeDirect(queueName string) {
	// 创建队列，有则跳过

	conn, err := amqp.Dial(mqUrl)
	ch, _ := conn.Channel()
	defer conn.Close()
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Logger.Error(err.Error())
	}

	msg, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Logger.Error(err.Error())
	}

	//r.channel.Qos()

	fmt.Println("----->")

	//for res := range msg{
	//	res.Ack(false)
	//	writeLog(string(res.Body),"./stat.txt")
	//}

	for {
		select {
		case res, ok := <-msg:
			fmt.Println(ok)
			res.Ack(false)
			fmt.Println(res.Body)
			writeLog(string(res.Body), "./stat.txt")
		default:
			fmt.Println("over")
			return
		}
	}

	//for res := range msg{
	//	//fmt.Println("写到数据库",string(res.Body))
	//	//log.Logger.Info(res.Body)
	//	//res.Ack(false)
	//	writeLog(string(res.Body),"./stat.txt")
	//
	//}
	//log.Logger.Info("over")
}

func writeLog(msg string, logPath string) {
	fd, _ := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer fd.Close()
	content := strings.Join([]string{msg, "\r\n"}, "")
	buf := []byte(content)
	fd.Write(buf)
}
