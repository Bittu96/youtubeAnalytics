package services

import (
	"encoding/json"
	"fmt"
	"youtubeAnalytics/models/channel"
)

func GetAllVideos(channelIds []string) {
	for _, channelID := range channelIds {
		_, err := channel.New(channelID).Load()
		if err != nil {
			prettyPrint(err)
		} else {
			prettyPrint("fetched")
		}
	}
}

func prettyPrint(data interface{}) {
	jb, _ := json.MarshalIndent(data, "", " ")
	fmt.Println(string(jb))
}
