package services

import (
	"log"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type Plot struct {
	X string
	Y opts.BarData
}

func getBarItems(plots []Plot) (Xdata []string, Ydata []opts.BarData) {
	xItems := make([]string, 0)
	yItems := make([]opts.BarData, 0)

	for _, plot := range plots {
		xItems = append(xItems, plot.X)
		yItems = append(yItems, plot.Y)
	}
	return xItems, yItems
}

func ViewInsights(plotName string, plots []Plot) {
	bar := charts.NewBar()

	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Youtube Analytics",
		Subtitle: "Views Counts",
	}))

	xData, yData := getBarItems(plots)
	bar.SetXAxis(xData).AddSeries("Category A", yData)

	f, _ := os.Create(plotName)
	err := bar.Render(f)
	log.Panic(err)
}
