package services

import (
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"time"
	"youtubeAnalytics/models/channel"
	"youtubeAnalytics/pkg/database"
	"youtubeAnalytics/pkg/keyword"

	"github.com/google/uuid"
)

func generateInsights(c channel.Channel) error {
	// extract top keys
	topKeyWords := getTopKeywordsFromChannel(c)
	fmt.Println(topKeyWords)

	// get channelInsights //todo in future
	// if err := generateChannelInsights(); err != nil {
	// 	return err
	// }

	// get video insights
	if err := GenerateVideoInsights(); err != nil {
		return err
	}

	return nil
}

// get top keywords from video titles and descriptions
func getTopKeywordsFromChannel(c channel.Channel) (topKeywords []string) {
	keywordFreqMap := make(map[string]int)

	// write frequencies of keywords to map
	for _, p := range c.Playlists {
		for _, v := range p.Videos {
			videoKeywords := keyword.Extract(v.Details.Title + v.Details.Description)
			for k := range videoKeywords {
				keywordFreqMap[k]++
			}
		}
	}

	var maxFreq int
	for _, freq := range keywordFreqMap {
		if freq > maxFreq {
			maxFreq = freq
		}
	}

	// create freq array and place keywords
	var freqSlice = make([][]string, maxFreq+1)
	for keyword, freq := range keywordFreqMap {
		freqSlice[freq] = append(freqSlice[freq], keyword)
	}

	fmt.Println(freqSlice)
	// get top 10 keywords from the
	for i := maxFreq - 1; i > 0; i-- {
		topKeywords = append(topKeywords, freqSlice[i]...)
		if len(topKeywords) > 10 {
			break
		}
	}

	return
}

// func generateChannelInsights() error { //tooo later
// 	rows, err := database.DBClient.Query("")
// 	if err != nil {
// 		return err
// 	}

// 	columns, err := rows.Columns()
// 	if err != nil {
// 		return err
// 	}

// 	for rows.Next() {
// 		values := make([]interface{}, len(columns))
// 		for i := range values {
// 			values[i] = new(interface{})
// 		}
// 		if err := rows.Scan(values...); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

func GenerateVideoInsights() error {
	queryFilePath := filepath.Join("pkg", "database", "queries", "generateVideoInsightsQuery.sql")
	c, err := os.ReadFile(queryFilePath)
	if err != nil {
		return err
	}

	videoInsightsQuery := string(c)
	fmt.Println(videoInsightsQuery)

	rows, err := database.GetClient().Query(videoInsightsQuery)
	if err != nil {
		return err
	}

	for rows.Next() {
		var res videoStats
		if err := rows.Scan(&res.videoId, &res.views, &res.viewsAfter24h, &res.likes, &res.likesAfter24h, &res.comments, &res.commentsAfter24h); err != nil {
			return err
		}
		fmt.Println(res)
		videoInsights := calculateInsights(res)
		fmt.Println(videoInsights)
		writeVideoInsightsToDB(videoInsights)
	}

	return nil
}

type videoStats struct {
	videoId          string
	views            int64
	viewsAfter24h    int64
	likes            int64
	likesAfter24h    int64
	comments         int64
	commentsAfter24h int64
}

type videoInsights struct {
	videoInsightsId           string
	videoId                   string
	viewCountInc              int64
	viewCountIncPerc          float64
	likeCountInc              int64
	likeCountIncPerc          float64
	commentCountInc           int64
	commentCountIncPerc       float64
	totalImpressionsCount     int64
	totalImpressionsCountPerc float64
	addedDate                 time.Time
}

func calculateInsights(vs videoStats) videoInsights {
	return videoInsights{
		videoInsightsId:           getNewUUID(),
		videoId:                   vs.videoId,
		viewCountInc:              vs.viewsAfter24h - vs.views,
		viewCountIncPerc:          calculatePerc(vs.viewsAfter24h, vs.views),
		likeCountInc:              vs.likesAfter24h - vs.likes,
		likeCountIncPerc:          calculatePerc(vs.likesAfter24h, vs.likes),
		commentCountInc:           vs.commentsAfter24h - vs.comments,
		commentCountIncPerc:       calculatePerc(vs.commentsAfter24h, vs.comments),
		totalImpressionsCount:     vs.viewsAfter24h + vs.likesAfter24h + vs.commentsAfter24h - vs.views - vs.likes - vs.comments,
		totalImpressionsCountPerc: calculatePerc(vs.viewsAfter24h+vs.likesAfter24h+vs.commentsAfter24h, vs.views+vs.likes+vs.comments),
	}
}

func getNewUUID() string {
	return uuid.NewString()
}

func calculatePerc(a, b int64) float64 {
	if b == 0 {
		return 0
	}
	return roundFloat(float64(a-b) / float64(b) * 100)
}

func roundFloat(val float64) float64 {
	ratio := math.Pow(10, float64(2))
	return math.Round(val*ratio) / ratio
}

func writeVideoInsightsToDB(v videoInsights) {
	insertvideoQuery := fmt.Sprintf("insert into youtubeAnalytics.videoinsights (video_insights_id, video_id, view_count_inc, like_count_inc, total_impressions, comment_count_inc, view_count_inc_perc, like_count_inc_perc, comment_count_inc_perc, total_impressions_perc) values ('%v', '%v', %v, %v, %v, %v, %v, %v, %v, %v);", v.videoInsightsId, v.videoId, v.viewCountInc, v.likeCountInc, v.commentCountInc, v.totalImpressionsCount, v.viewCountIncPerc, v.likeCountIncPerc, v.commentCountIncPerc, v.totalImpressionsCountPerc)
	runQuery(insertvideoQuery)
}

func runQuery(query string) {
	res, err := database.GetClient().Exec(query)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("row inserted", res)
	}
}
