package services

import (
	"fmt"
	"time"
)

const INTERVAL_PERIOD time.Duration = 24 * time.Hour
const HOUR_TO_TICK int = 23
const MINUTE_TO_TICK int = 21
const SECOND_TO_TICK int = 03

type jobTicker struct {
	t *time.Timer
}

func getNextTickDuration() time.Duration {
	now := time.Now()
	nextTick := time.Date(now.Year(), now.Month(), now.Day(), HOUR_TO_TICK, MINUTE_TO_TICK, SECOND_TO_TICK, 0, time.Local)
	if nextTick.Before(now) {
		nextTick = nextTick.Add(INTERVAL_PERIOD)
	}
	fmt.Println(nextTick)

	return time.Until(time.Now())
}

func NewJobTicker() jobTicker {
	fmt.Println("new tick here")
	return jobTicker{time.NewTimer(getNextTickDuration())}
}

func (jt jobTicker) updateJobTicker() {
	fmt.Println("next tick here")
	jt.t.Reset(getNextTickDuration())
}

func NewBar() {
	jt := NewJobTicker()
	for {
		<-jt.t.C
		fmt.Println(time.Now(), "- just ticked")
		jt.updateJobTicker()
	}
}

// func Job() {
// 	services.ViewInsights()
// }

// return newChannel
// plots := make([]Plot, 0)
// for _, playlist := range newChannel.Playlists {
// 	plots = append(plots, Plot{
// 		X: string(playlist.PlaylistID[0]),
// 		Y: opts.BarData{Value: playlist.VideoCount},
// 	})
// 	// for _, video := range playlist.Videos {
// 	// 	plots = append(plots, Plot{
// 	// 		X: playlist.PlaylistID,
// 	// 		Y: opts.BarData{Value: playlist.VideoCount},
// 	// 	})
// 	// }
// }
// ViewInsights("test.html", plots)
