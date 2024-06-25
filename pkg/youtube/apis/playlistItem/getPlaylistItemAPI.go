package playlistItemAPI

import (
	"encoding/json"
	"net/url"
	"time"
	"youtubeAnalytics/configs"
	"youtubeAnalytics/pkg/myHttp"
)

type PlaylistItemAPIRequest struct {
	Key        string `json:"key"`
	Part       string `json:"part"`
	PlaylistID string `json:"playlistId"`
	MaxResults string `json:"maxResults"`
	PageToken  string `json:"pageToken"`
}

type PlaylistItemAPIResponse struct {
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
	VideoID          string    `json:"videoId"`
	VideoPublishedAt time.Time `json:"videoPublishedAt"`
}

type Snippet struct {
	PublishedAt time.Time `json:"publishedAt"`
	ChannelID   string    `json:"channelId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	// Thumbnails             Thumbnails `json:"thumbnails"`
	ChannelTitle           string     `json:"channelTitle"`
	PlaylistID             string     `json:"playlistId"`
	Position               int64      `json:"position"`
	ResourceID             ResourceID `json:"resourceId"`
	VideoOwnerChannelTitle string     `json:"videoOwnerChannelTitle"`
	VideoOwnerChannelID    string     `json:"videoOwnerChannelId"`
}

type ResourceID struct {
	Kind    string `json:"kind"`
	VideoID string `json:"videoId"`
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

func Request(playlistId string) PlaylistItemAPIRequest {
	return PlaylistItemAPIRequest{
		Key:        configs.YoutubeDataAPIKey,
		Part:       "snippet,contentDetails",
		PlaylistID: playlistId,
		MaxResults: "50",
	}
}

func (c PlaylistItemAPIRequest) SetPage(pageToken string) PlaylistItemAPIRequest {
	c.PageToken = pageToken
	return c
}

func (c PlaylistItemAPIRequest) MakeAPICall() (response PlaylistItemAPIResponse, err error) {
	url, err := url.JoinPath(configs.YoutubeDataAPIBaseURL, configs.YouTubeDataAPIPlaylistItemsPath)
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
