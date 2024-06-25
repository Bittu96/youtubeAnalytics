package playlistCollecionAPI

import (
	"encoding/json"
	"net/url"
	"time"
	"youtubeAnalytics/configs"
	"youtubeAnalytics/pkg/myHttp"
)

type PlaylistCollecionAPIRequest struct {
	Key        string `json:"key"`
	Part       string `json:"part"`
	ChannelID  string `json:"channelId"`
	MaxResults string `json:"maxResults"`
	PageToken  string `json:"pageToken"`
}

type PlaylistCollecionAPIResponse struct {
	Kind          string   `json:"kind"`
	Etag          string   `json:"etag"`
	NextPageToken string   `json:"nextPageToken"`
	PageInfo      PageInfo `json:"pageInfo"`
	Items         []Item   `json:"items"`
	Error         ErrError `json:"error"`
}

type Item struct {
	Kind           string         `json:"kind"`
	Etag           string         `json:"etag"`
	ID             string         `json:"id"`
	Snippet        Snippet        `json:"snippet"`
	ContentDetails ContentDetails `json:"contentDetails"`
}

type ContentDetails struct {
	ItemCount int64 `json:"itemCount"`
}

type Snippet struct {
	PublishedAt time.Time `json:"publishedAt"`
	ChannelID   string    `json:"channelId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	// Thumbnails   Thumbnails `json:"thumbnails"`
	ChannelTitle string    `json:"channelTitle"`
	Localized    Localized `json:"localized"`
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

func Request(channelId string) PlaylistCollecionAPIRequest {
	return PlaylistCollecionAPIRequest{
		Key:        configs.YoutubeDataAPIKey,
		Part:       "snippet,contentDetails",
		ChannelID:  channelId,
		MaxResults: "50",
	}
}

func (c PlaylistCollecionAPIRequest) GetPage(pageToken string) PlaylistCollecionAPIRequest {
	c.PageToken = pageToken
	return c
}

func (c PlaylistCollecionAPIRequest) MakeAPICall() (response PlaylistCollecionAPIResponse, err error) {
	url, err := url.JoinPath(configs.YoutubeDataAPIBaseURL, configs.YouTubeDataAPIPlaylistCollectionPath)
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
