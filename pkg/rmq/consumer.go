package rmq

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
	"youtubeAnalytics/pkg/database"
	channelAPI "youtubeAnalytics/pkg/youtube/apis/channel"
	playlistItemAPI "youtubeAnalytics/pkg/youtube/apis/playlistItem"
	videoAPI "youtubeAnalytics/pkg/youtube/apis/video"

	"github.com/google/uuid"
	"github.com/rabbitmq/amqp091-go"
)

var (
	RMQConsumerClient *RMQ
	dataConsumerCount int = 3
)

func (r *RMQ) StartConsumer() {
	msgs, err := r.rmqChannel.Consume(r.queue.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal("Failed to register a consumer")
		return
	}

	var forever chan struct{}

	for i := 0; i < dataConsumerCount; i++ {
		go dataConsumer(i, msgs)
	}

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func dataConsumer(consumerId int, msgs <-chan amqp091.Delivery) {
	log.Println("loading consumer", consumerId)
	for d := range msgs {
		// log.Printf("[%v]Received a message: %s\n", consumerId, d.Body)
		log.Printf("[%v]Received a message type: %s\n", consumerId, d.Type)
		writeToDB(d.Type, d)
	}
}

func writeToDB(entity string, data amqp091.Delivery) {
	switch entity {
	case "channel":
		writeChannelDataToDB(data.Body)
	case "playlist":
		writePlaylistDataToDB(data.Body)
	case "video":
		writeVideoDataToDB(data.Body)
	default:
		log.Println("unknown entity")
	}
}

type channel struct {
	ChannelID       string             `json:"channel_id"`
	SubscriberCount int64              `json:"subscriber_count"`
	VideoCount      int64              `json:"video_count"`
	ViewCount       int64              `json:"view_count"`
	Details         channelAPI.Snippet `json:"details"`
}

type playlist struct {
	PlaylistID string                  `json:"playlist_id"`
	ChannelID  string                  `json:"channel_id"`
	VideoCount int64                   `json:"video_count"`
	Details    playlistItemAPI.Snippet `json:"details"`
}

type video struct {
	VideoID     string           `json:"video_id"`
	Views       int64            `json:"views"`
	Likes       int64            `json:"likes"`
	Comments    int64            `json:"comments"`
	PublishedAt time.Time        `json:"published_at"`
	Details     videoAPI.Snippet `json:"details"`
}

func writeChannelDataToDB(data []byte) {
	var c channel
	if err := json.Unmarshal(data, &c); err != nil {
		return
	}

	insertChannelQuery := fmt.Sprintf("insert into youtubeAnalytics.channel (channel_id, details) values ('%v', '%v');", c.ChannelID, jsonbEscapeString(c.Details))
	runQuery(insertChannelQuery)
	insertChannelStatisticsQuery := fmt.Sprintf("insert into youtubeAnalytics.channelStatistics (channel_statistics_id, channel_id, subscriber_count, view_count, video_count) values ('%v', '%v', %v,  %v,  %v);", getNewUUID(), c.ChannelID, c.SubscriberCount, c.ViewCount, c.VideoCount)
	runQuery(insertChannelStatisticsQuery)
}

func writePlaylistDataToDB(data []byte) {
	var p playlist
	if err := json.Unmarshal(data, &p); err != nil {
		log.Println(err)
		return
	}

	insertPlaylistQuery := fmt.Sprintf("insert into youtubeAnalytics.playlist (playlist_id, channel_id, details) values ('%v', '%v', '%v');", p.PlaylistID, p.ChannelID, jsonbEscapeString(p.Details))
	runQuery(insertPlaylistQuery)
}

func writeVideoDataToDB(data []byte) {
	var v video
	if err := json.Unmarshal(data, &v); err != nil {
		return
	}

	insertvideoQuery := fmt.Sprintf("insert into youtubeAnalytics.video (video_id, channel_id, details) values ('%v', '%v', '%v');", v.VideoID, v.Details.ChannelID, jsonbEscapeString(v.Details))
	runQuery(insertvideoQuery)
	insertVideoStatisticsQuery := fmt.Sprintf("insert into youtubeAnalytics.videoStatistics (video_statistics_id, video_id, view_count, like_count, comment_count) values ('%v', '%v', %v,  %v,  %v);", getNewUUID(), v.VideoID, v.Views, v.Likes, v.Comments)
	runQuery(insertVideoStatisticsQuery)
}

func runQuery(query string) {
	res, err := database.DBClient.Exec(query)
	if err != nil {
		// log.Println(query)
		log.Println(err)
	} else {
		log.Println("row inserted", res)
	}
}

func jsonbEscapeString(data interface{}) string {
	detailsJsonBytes, _ := json.Marshal(data)
	return strings.ReplaceAll(string(detailsJsonBytes), `'`, `''`)
}

func getNewUUID() string {
	return uuid.NewString()
}
