package main

import (
	"log"
	"youtubeAnalytics/configs"
	"youtubeAnalytics/pkg/database"
	"youtubeAnalytics/pkg/rmq"
	"youtubeAnalytics/services"
)

// append target channels here
var targetChannels = []string{
	"UC5OrDvL9DscpcAstz7JnQGA",
	"UC70pKToywlxOGdgIvz8gYqA",
}

func main() {
	// connect and init dbClient
	dbclient, err := database.New(configs.DBHost, configs.DBPort, configs.DBUser, configs.DBPass, configs.DBName).Connect()
	if err != nil {
		log.Fatal(err)
		return
	} else {
		defer dbclient.Close()
	}

	// connect and init rmq publisher client
	rmq.RMQPublisherClient = rmq.New(configs.RMQURL, configs.QueueName)
	if err := rmq.RMQPublisherClient.Connect(); err != nil {
		log.Fatal(err)
		return
	} else {
		defer rmq.RMQPublisherClient.Close()
	}

	// process target channels
	services.ProcessChannels(targetChannels)

	log.Println("main end")
}
