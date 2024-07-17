package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"youtubeAnalytics/configs"
	"youtubeAnalytics/pkg/database"
	"youtubeAnalytics/pkg/myCron"
	"youtubeAnalytics/pkg/rmq"
	"youtubeAnalytics/services"

	"github.com/gin-gonic/gin"
)

// append target channels here
var targetChannels = []string{
	"UC5OrDvL9DscpcAstz7JnQGA",
	"UC70pKToywlxOGdgIvz8gYqA",
}

func init() {
	gin.SetMode(gin.DebugMode)

	// init db client
	database.New(configs.DBHost, configs.DBPort, configs.DBUser, configs.DBPass, configs.DBName)
	database.GetClient()

	// init rmq client
	rmq.New(configs.RMQURL, configs.QueueName)
	rmq.GetClient()
}

func main() {
	routes := gin.Default()
	routes.Use(CORSMiddleware())

	// route for insights
	routes.GET("/", func(c *gin.Context) {
		routes.LoadHTMLFiles("data/insights_demo.html")
		c.HTML(http.StatusOK, "insights_demo.html", gin.H{
			"content": "welcome to youtube analytics",
		})
	})

	myServer := &http.Server{
		Addr:    ":8080",
		Handler: routes.Handler(),
	}

	// start my cron
	go myCron.Start(configs.CronInterval, func() {
		// process target channels
		services.ProcessChannels(targetChannels)
		services.RenderVideoInsights()
	})

	// start my server
	go func() {
		if err := myServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("listen:", err)
		}
	}()

	gracefulShutdown(func(ctx context.Context) {
		// stop server
		if err := myServer.Shutdown(ctx); err != nil {
			log.Fatal(err)
		} else {
			log.Println("http server stopped")
		}
		// close db client
		database.CloseClient()
		// close rmq client
		rmq.CloseClient()
	})
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Writer.Header().Set("Content-Type", "*; application/json; charset=utf-8; application/x-www-form-urlencoded;")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Cookie")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
		} else {
			c.Next()
		}

		c.Next()
	}
}

func gracefulShutdown(task func(ctx context.Context)) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("shutdown initated")
	ctx, cancel := context.WithTimeout(context.Background(), configs.ShutdownDelay)
	defer cancel()

	task(ctx)

	<-ctx.Done()
	log.Println("shutdown completed")
}
