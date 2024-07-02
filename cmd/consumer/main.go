package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
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
	// start data consumer
	go rmq.GetClient().StartConsumer()

	gracefulShutdown(func(ctx context.Context) {
		// close db client
		database.CloseClient()
		// close rmq client
		rmq.CloseClient()
	})
}

func gracefulShutdown(task func(ctx context.Context)) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("shutdown initated")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	task(ctx)

	<-ctx.Done()
	log.Println("shutdown completed")
}
