package main

import (
	"log"
	"youtubeAnalytics/configs"
	"youtubeAnalytics/pkg/database"
	"youtubeAnalytics/pkg/rmq"
)

func main() {
	dbclient, err := database.New(configs.DBHost, configs.DBPort, configs.DBUser, configs.DBPass, configs.DBName).Connect()
	if err != nil {
		log.Fatal(err)
		return
	} else {
		defer dbclient.Close()
	}

	// systemctl start rabbitmq-server
	rmq.RMQConsumerClient = rmq.New(configs.RMQURL, configs.QueueName)
	if err = rmq.RMQConsumerClient.Connect(); err != nil {
		log.Fatal(err)
		return
	} else {
		defer rmq.RMQConsumerClient.Close()
	}

	rmq.RMQConsumerClient.StartConsumer()
}
