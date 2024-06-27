package main

import (
	"log"
	"youtubeAnalytics/configs"
	"youtubeAnalytics/pkg/database"
	"youtubeAnalytics/pkg/rmq"
	"youtubeAnalytics/services"
)

func main() {
	// connect to database
	dbclient, err := database.New(configs.DBHost, configs.DBPort, configs.DBUser, configs.DBPass, configs.DBName).Connect()
	if err != nil {
		log.Fatal(err)
		return
	} else {
		defer dbclient.Close()
	}

	// connect to rmq
	rmq.RMQPublisherClient = rmq.New(configs.RMQURL, configs.QueueName)
	if err := rmq.RMQPublisherClient.Connect(); err != nil {
		log.Fatal(err)
		return
	} else {
		defer rmq.RMQPublisherClient.Close()
	}

	// append target channels
	targetChannels := []string{
		"UC5OrDvL9DscpcAstz7JnQGA",
		"UC70pKToywlxOGdgIvz8gYqA",
	}

	// get video data of target channels
	services.LoadChannels(targetChannels)

	// get video insights
	if err := services.GenerateVideoInsights(); err != nil {
		panic(err)
	}
}
