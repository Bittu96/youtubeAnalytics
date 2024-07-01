package video

import (
	"log"
	"strconv"
	"time"
	"youtubeAnalytics/pkg/rmq"
	videoAPI "youtubeAnalytics/pkg/youtube/apis/video"
)

type Video struct {
	VideoID     string           `json:"video_id"`
	Views       int64            `json:"views"`
	Likes       int64            `json:"likes"`
	Comments    int64            `json:"comments"`
	PublishedAt time.Time        `json:"published_at"`
	Details     videoAPI.Snippet `json:"details"`
}

// create new video
func New(videoId string) Video {
	return Video{
		VideoID: videoId,
	}
}

// download video info
func (v Video) Download() (Video, error) {
	videoAPIResponse, err := videoAPI.Request(v.VideoID).MakeAPICall()
	if err != nil {
		return Video{}, err
	} else if len(videoAPIResponse.Items) == 0 {
		return Video{}, nil
	} else {
		for _, videoData := range videoAPIResponse.Items {
			v.Views, _ = strconv.ParseInt(videoData.Statistics.ViewCount, 10, 64)
			v.Likes, _ = strconv.ParseInt(videoData.Statistics.LikeCount, 10, 64)
			v.Comments, _ = strconv.ParseInt(videoData.Statistics.CommentCount, 10, 64)
			v.PublishedAt = videoData.Snippet.PublishedAt
			v.Details = videoData.Snippet
		}
	}

	// push video info to consumer
	if err := rmq.GetClient().Publish("video", v); err != nil {
		log.Println(err)
	}
	return v, nil
}
