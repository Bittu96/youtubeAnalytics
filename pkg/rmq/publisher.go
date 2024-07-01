package rmq

import (
	"context"
	"encoding/json"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (r *RMQ) Publish(entity string, msg interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	msgJsonBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	publishMessage := amqp.Publishing{
		ContentType: "text/json",
		Body:        msgJsonBytes,
		Type:        entity,
	}

	if err = r.channel.PublishWithContext(ctx, "", r.queueName, false, false, publishMessage); err != nil {
		log.Println(err)
	}
	for i := 0; i < maxRetryCount && shouldRetryPublisher(err); i++ {
		log.Println("retrying PublishWithContext...")
		time.Sleep(retryInterval)
		if err = r.channel.PublishWithContext(ctx, "", r.queueName, false, false, publishMessage); err != nil {
			log.Println(err)
		}
	}
	if err != nil {
		log.Println("Failed to PublishWithContext")
		return err
	}

	// log.Printf(" [x] Sent %v\n", "msg")
	return nil
}

func shouldRetryPublisher(err error) bool {
	return err != nil
}
