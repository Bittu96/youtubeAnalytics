package services

import (
	"log"
	"youtubeAnalytics/models/channel"
)

func ProcessChannels(channelIds []string) {
	for _, channelID := range channelIds {

		// get channel data from youtube APIs
		channelData, err := channel.New(channelID).Download()
		if err != nil {
			log.Printf("[%s] channel data fetch failed due to %s\n", channelID, err.Error())
		} else {
			log.Printf("[%s] channel data fetch success\n", channelID)

			// get channel insights from channel data
			if err = generateInsights(channelData); err != nil {
				log.Printf("[%s] channel insights fetch failed due to %s\n", channelID, err.Error())
			} else {
				log.Printf("[%s] channel insights fetch success\n", channelID)
			}
		}
	}
}
