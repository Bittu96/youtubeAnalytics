package rmq

import (
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RMQ struct {
	url           string
	queueName     string
	rmqConnection *amqp.Connection
	rmqChannel    *amqp.Channel
	queue         amqp.Queue
}

const (
	maxRetryCount = 5
	retryInterval = time.Second
)

func New(url, queueName string) *RMQ {
	return &RMQ{url: url, queueName: queueName}
}

func (r *RMQ) Connect() (err error) {
	r.rmqConnection, err = amqp.Dial(r.url)
	for i := 0; i < maxRetryCount && shouldRetryConnection(r.rmqConnection, err); i++ {
		log.Println("retrying RMQConnection...")
		time.Sleep(retryInterval)
		r.rmqConnection, err = amqp.Dial(r.url)
	}
	if err != nil {
		return err
	}

	r.rmqChannel, err = r.rmqConnection.Channel()
	for i := 0; i < maxRetryCount && shouldRetryConnection(r.rmqChannel, err); i++ {
		log.Println("retrying RMQChannel...")
		time.Sleep(retryInterval)
		r.rmqChannel, err = r.rmqConnection.Channel()
	}
	if err != nil {
		return err
	}

	r.queue, err = r.rmqChannel.QueueDeclare(r.queueName, false, false, false, false, nil)
	for i := 0; i < maxRetryCount && shouldRetryConnection(r.queue, err); i++ {
		log.Println("retrying QueueDeclare...")
		time.Sleep(retryInterval)
		r.queue, err = r.rmqChannel.QueueDeclare(r.queueName, false, false, false, false, nil)
	}
	if err != nil {
		return err
	}

	return nil
}

func (r *RMQ) Close() {
	r.rmqConnection.Close()
}

func shouldRetryConnection(conn interface{}, err error) bool {
	if err != nil || conn == nil {
		return true
	}
	return false
}
