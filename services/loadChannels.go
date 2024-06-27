package services

import (
	"log"
	"youtubeAnalytics/models/channel"
)

func LoadChannels(channelIds []string) {
	for _, channelID := range channelIds {
		channelData, err := channel.New(channelID).Load()
		if err != nil {
			log.Println("channel data fetch failed:", channelID, err.Error())
		} else {
			log.Println("channel data fetch success:", channelID)
			if err = generateInsights(channelData); err != nil {
				log.Println("channel insights fetch failed:", channelID, err.Error())
			} else {
				log.Println("channel insights fetch success:", channelID)
			}
		}
	}
}

// func prettyPrint(data interface{}) {
// 	jb, _ := json.MarshalIndent(data, "", " ")
// 	fmt.Println(string(jb))
// }
