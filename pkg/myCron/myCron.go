package myCron

import (
	"fmt"
	"time"
)

func Start(interval time.Duration, task func()) {
	ticker := time.NewTicker(interval)

	for t := range ticker.C {
		fmt.Println("Tick at", t)
		task()
	}
}
