package rmq

import (
	"log"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RMQ struct {
	url       string
	queueName string
	conn      *amqp.Connection
	channel   *amqp.Channel
	queue     amqp.Queue
	mux       *sync.Mutex
}

const (
	maxRetryCount = 5
	retryInterval = time.Second
)

// global rmq client variable
var rmqClient *RMQ

func New(url, queueName string) {
	if rmqClient == nil {
		rmqClient = &RMQ{
			url:       url,
			queueName: queueName,
			mux:       &sync.Mutex{}}
	}
}

func (r *RMQ) Connect() (err error) {
	r.conn, err = amqp.Dial(r.url)
	// for i := 0; i < maxRetryCount && shouldRetryConnection(r.conn, err); i++ {
	// 	log.Println("retrying RMQConnection...")
	// 	time.Sleep(retryInterval)
	// 	r.conn, err = amqp.Dial(r.url)
	// }
	if err != nil {
		return err
	}

	r.channel, err = r.conn.Channel()
	// for i := 0; i < maxRetryCount && shouldRetryConnection(r.channel, err); i++ {
	// 	log.Println("retrying RMQChannel...")
	// 	time.Sleep(retryInterval)
	// 	r.channel, err = r.conn.Channel()
	// }
	if err != nil {
		return err
	}

	r.queue, err = r.channel.QueueDeclare(r.queueName, false, false, false, false, nil)
	// for i := 0; i < maxRetryCount && shouldRetryConnection(r.queue, err); i++ {
	// 	log.Println("retrying QueueDeclare...")
	// 	time.Sleep(retryInterval)
	// 	r.queue, err = r.channel.QueueDeclare(r.queueName, false, false, false, false, nil)
	// }
	if err != nil {
		return err
	}

	log.Println("rmq connection success!")
	return nil
}

// func shouldRetryConnection(conn interface{}, err error) bool {
// 	if err != nil || conn == nil {
// 		return true
// 	}
// 	return false
// }

func GetClient() *RMQ {
	if rmqClient.conn == nil || rmqClient.channel == nil || rmqClient.queue.Name == "" {
		rmqClient.mux.Lock()
		if rmqClient.conn == nil || rmqClient.channel == nil || rmqClient.queue.Name == "" {
			for rmqClient.Connect() != nil {
				time.Sleep(time.Second)
			}
		}
		rmqClient.mux.Unlock()
	}

	return rmqClient
}

func CloseClient() {
	if rmqClient.conn == nil {
		log.Println("rmq connection already closed!")
	}

	if err := rmqClient.conn.Close(); err != nil {
		log.Println("rmq connection close failed!")
	}
}
