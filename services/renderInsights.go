package services

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"youtubeAnalytics/pkg/chart"
	"youtubeAnalytics/pkg/database"

	"github.com/go-echarts/go-echarts/v2/opts"
)

func RenderVideoInsights() {
	insightPlots, err := getVideoInsightPlots()
	if err != nil {
		panic(err)
	}

	chart.RenderInsights("insights_demo.html", insightPlots)
}

func getVideoInsightPlots() (chart.InsightPlots, error) {
	videoInsightsRes := make([]videoInsights, 0)
	insightPlots := chart.InsightPlots{}

	queryFilePath := filepath.Join("pkg", "database", "queries", "getVideoInsightsQuery.sql")
	q, err := os.ReadFile(queryFilePath)
	if err != nil {
		return chart.InsightPlots{}, err
	}

	videoInsightsQuery := string(q)
	fmt.Println(videoInsightsQuery)

	rows, err := database.GetClient().Query(videoInsightsQuery)
	if err != nil {
		return chart.InsightPlots{}, err
	}

	for rows.Next() {
		var res videoInsights
		if err := rows.Scan(&res.videoInsightsId, &res.videoId, &res.viewCountInc, &res.viewCountIncPerc, &res.likeCountInc, &res.likeCountIncPerc, &res.commentCountInc, &res.commentCountIncPerc, &res.totalImpressionsCount, &res.totalImpressionsCountPerc, &res.addedDate); err != nil {
			return chart.InsightPlots{}, err
		}
		if res.totalImpressionsCountPerc > 10 {
			videoInsightsRes = append(videoInsightsRes, res)
			fmt.Println(res)
		}
	}

	sort.Slice(videoInsightsRes, func(i, j int) bool {
		return videoInsightsRes[i].totalImpressionsCountPerc > videoInsightsRes[j].totalImpressionsCountPerc
	})

	// dummy padding
	// videoInsightsRes = append([]videoInsights{{}}, videoInsightsRes...)
	// videoInsightsRes = append(videoInsightsRes, videoInsights{})

	for _, res := range videoInsightsRes {
		insightPlots.VideoIds = append(insightPlots.VideoIds, res.videoId)
		insightPlots.Views = append(insightPlots.Views, opts.BarData{Value: res.viewCountIncPerc})
		insightPlots.Likes = append(insightPlots.Likes, opts.BarData{Value: res.likeCountIncPerc})
		insightPlots.Comments = append(insightPlots.Comments, opts.BarData{Value: res.commentCountIncPerc})
		insightPlots.Impressions = append(insightPlots.Impressions, opts.BarData{Value: res.totalImpressionsCountPerc})
	}

	return insightPlots, nil
}
