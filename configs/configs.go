package configs

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

const (
	// client connections
	MaxRetryCount = 10
	RetryInterval = 2 * time.Second

	// cron
	CronInterval = 1 * time.Minute

	// shutodown delay
	ShutdownDelay = 2 * time.Second
)

var (
	// youtube API creds
	YoutubeDataAPIBaseURL                = getEnv("YoutubeDataAPIBaseURL").val
	YouTubeDataAPIChannelsPath           = getEnv("YouTubeDataAPIChannelsPath").val
	YouTubeDataAPIPlaylistCollectionPath = getEnv("YouTubeDataAPIPlaylistCollectionPath").val
	YouTubeDataAPIPlaylistItemsPath      = getEnv("YouTubeDataAPIPlaylistItemsPath").val
	YouTubeDataAPIVideosPath             = getEnv("YouTubeDataAPIVideosPath").val
	YoutubeDataAPIKey                    = getEnv("YoutubeDataAPIKey").val

	// rmq creds
	RMQURL    = getEnv("RMQURL").val
	QueueName = getEnv("QueueName").val

	// db creds
	DBHost = getEnv("DBHost").val
	DBPort = getEnv("DBPort").ToInt()
	DBUser = getEnv("DBUser").val
	DBPass = getEnv("DBPass").val
	DBName = getEnv("DBName").val

	// flags
	LoadVideosFromPlaylists = getEnv("LoadVideosFromPlaylists").ToBool()
)

func LoadEnvs() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file", err)
	}
}

func getEnv(key string) (e env) {
	if value := os.Getenv(key); value == "" {
		if err := godotenv.Load(".env"); err != nil {
			log.Panic("Error loading .env file", err)
		} else {
			if value = os.Getenv(key); value == "" {
				log.Panicf("env key %v not found", key)
			} else {
				e.val = value
			}
		}
	} else {
		e.val = value
	}
	return
}

type env struct {
	val string
}

func (e env) ToInt() int {
	if i, err := strconv.Atoi(e.val); err != nil {
		panic(err)
	} else {
		return i
	}
}
func (e env) ToFloat() float64 {
	if f, err := strconv.ParseFloat(e.val, 64); err != nil {
		panic(err)
	} else {
		return f
	}
}
func (e env) ToBool() bool {
	if b, err := strconv.ParseBool(e.val); err != nil {
		panic(err)
	} else {
		return b
	}
}
