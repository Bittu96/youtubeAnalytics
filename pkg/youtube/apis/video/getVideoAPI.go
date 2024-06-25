package videoAPI

import (
	"encoding/json"
	"net/url"
	"time"
	"youtubeAnalytics/configs"
	"youtubeAnalytics/pkg/myHttp"
)

type VideoAPIRequest struct {
	Key  string `json:"key"`
	Part string `json:"part"`
	ID   string `json:"id"`
}

type VideoAPIResponse struct {
	Kind     string   `json:"kind"`
	Etag     string   `json:"etag"`
	Items    []Item   `json:"items"`
	PageInfo PageInfo `json:"pageInfo"`
	Error    ErrError `json:"error"`
}

type Item struct {
	Kind           string         `json:"kind"`
	Etag           string         `json:"etag"`
	ID             string         `json:"id"`
	Snippet        Snippet        `json:"snippet"`
	ContentDetails ContentDetails `json:"contentDetails"`
	Statistics     Statistics     `json:"statistics"`
}

type ContentDetails struct {
	Duration        string        `json:"duration"`
	Dimension       string        `json:"dimension"`
	Definition      string        `json:"definition"`
	Caption         string        `json:"caption"`
	LicensedContent bool          `json:"licensedContent"`
	ContentRating   ContentRating `json:"contentRating"`
	Projection      string        `json:"projection"`
}

type ContentRating struct {
}

type Snippet struct {
	PublishedAt time.Time `json:"publishedAt"`
	ChannelID   string    `json:"channelId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	// Thumbnails           Thumbnails `json:"thumbnails"`
	ChannelTitle         string    `json:"channelTitle"`
	Tags                 []string  `json:"tags"`
	CategoryID           string    `json:"categoryId"`
	LiveBroadcastContent string    `json:"liveBroadcastContent"`
	Localized            Localized `json:"localized"`
	DefaultAudioLanguage string    `json:"defaultAudioLanguage"`
}

type Localized struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Thumbnails struct {
	Default  Default `json:"default"`
	Medium   Default `json:"medium"`
	High     Default `json:"high"`
	Standard Default `json:"standard"`
	Maxres   Default `json:"maxres"`
}

type Default struct {
	URL    string `json:"url"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
}

type Statistics struct {
	ViewCount     string `json:"viewCount"`
	LikeCount     string `json:"likeCount"`
	FavoriteCount string `json:"favoriteCount"`
	CommentCount  string `json:"commentCount"`
}

type PageInfo struct {
	TotalResults   int64 `json:"totalResults"`
	ResultsPerPage int64 `json:"resultsPerPage"`
}

type ErrError struct {
	Code    int64          `json:"code"`
	Message string         `json:"message"`
	Errors  []ErrorElement `json:"errors"`
}

type ErrorElement struct {
	Message      string `json:"message"`
	Domain       string `json:"domain"`
	Reason       string `json:"reason"`
	Location     string `json:"location"`
	LocationType string `json:"locationType"`
}

func Request(videoID string) VideoAPIRequest {
	return VideoAPIRequest{
		Key:  configs.YoutubeDataAPIKey,
		Part: "snippet,contentDetails,statistics",
		ID:   videoID,
	}
}

func (c VideoAPIRequest) MakeAPICall() (response VideoAPIResponse, err error) {
	url, err := url.JoinPath(configs.YoutubeDataAPIBaseURL, configs.YouTubeDataAPIVideosPath)
	if err != nil {
		return
	}

	body, err := myHttp.NewRequest("GET", url, c).Do()
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &response)
	return
}
