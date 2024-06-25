package services

import "youtubeAnalytics/pkg/keyword"

func ExtractKeywords(text string) {
	keyword.Extract(text)
}
