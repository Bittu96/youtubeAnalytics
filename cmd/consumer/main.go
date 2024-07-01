package main

import (
	"youtubeAnalytics/configs"
	"youtubeAnalytics/pkg/database"
	"youtubeAnalytics/pkg/rmq"
)

func init() {
	// init db client
	database.New(configs.DBHost, configs.DBPort, configs.DBUser, configs.DBPass, configs.DBName)
	database.GetClient()

	// init rmq client
	rmq.New(configs.RMQURL, configs.QueueName)
	rmq.GetClient()
}

func main() {
	defer database.CloseClient()
	defer rmq.CloseClient()

	// start data consumer
	rmq.GetClient().StartConsumer()
}
