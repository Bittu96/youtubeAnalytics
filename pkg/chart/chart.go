package chart

import (
	"log"
	"os"
	"path/filepath"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type InsightPlots struct {
	VideoIds    []string
	Views       []opts.BarData
	Likes       []opts.BarData
	Comments    []opts.BarData
	Impressions []opts.BarData
}

func RenderInsights(plotName string, plots InsightPlots) {
	bar := charts.NewBar()

	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Youtube Analytics",
		Subtitle: "Todays Channel Insights",
	}), charts.WithXAxisOpts(opts.XAxis{
		AxisLabel: &opts.AxisLabel{Show: false, Rotate: 90, Interval: "0"},
	}), charts.WithTooltipOpts(opts.Tooltip{
		Show: true, Trigger: "axis", AxisPointer: &opts.AxisPointer{Type: "shadow"},
	}), charts.WithGridOpts(opts.Grid{
		// Bottom: "1%",
		// Left:   "1%",
		// Top:    "1%",
		// Right:  "1%",
		// ContainLabel: false,
		// Height: "100%",
		// Width:  "100%",
	}))

	// xData, yData := getBarItems(plots)
	bar.SetXAxis(plots.VideoIds).
		AddSeries("views", plots.Views).
		AddSeries("likes", plots.Likes).
		AddSeries("comments", plots.Comments).
		AddSeries("impressions", plots.Impressions)

	f, _ := os.Create(filepath.Join("data", plotName))
	if err := bar.Render(f); err != nil {
		log.Panic(err)
	}
}
