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

	// process target channels
	services.ProcessChannels(targetChannels)
	services.RenderVideoInsights()

	log.Println("main end")
}
