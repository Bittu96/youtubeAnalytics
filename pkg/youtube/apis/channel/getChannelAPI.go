package channelAPI

import (
	"encoding/json"
	"net/url"
	"time"
	"youtubeAnalytics/configs"
	"youtubeAnalytics/pkg/myHttp"
)

type ChannelAPIRequest struct {
	Key  string `json:"key"`
	Part string `json:"part"`
	ID   string `json:"id"`
}

type ChannelAPIResponse struct {
	Kind     string   `json:"kind"`
	Etag     string   `json:"etag"`
	PageInfo PageInfo `json:"pageInfo"`
	Items    []Item   `json:"items"`
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
	RelatedPlaylists RelatedPlaylists `json:"relatedPlaylists"`
}

type RelatedPlaylists struct {
	Likes   string `json:"likes"`
	Uploads string `json:"uploads"`
}

type Snippet struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CustomURL   string    `json:"customUrl"`
	PublishedAt time.Time `json:"publishedAt"`
	// Thumbnails  Thumbnails `json:"thumbnails"`
	Localized Localized `json:"localized"`
}

type Localized struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Thumbnails struct {
	Default Default `json:"default"`
	Medium  Default `json:"medium"`
	High    Default `json:"high"`
}

type Default struct {
	URL    string `json:"url"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
}

type Statistics struct {
	ViewCount             string `json:"viewCount"`
	SubscriberCount       string `json:"subscriberCount"`
	HiddenSubscriberCount bool   `json:"hiddenSubscriberCount"`
	VideoCount            string `json:"videoCount"`
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

func Request(channelID string) ChannelAPIRequest {
	return ChannelAPIRequest{
		Key:  configs.YoutubeDataAPIKey,
		Part: "snippet,contentDetails,statistics",
		ID:   channelID,
	}
}

func (c ChannelAPIRequest) MakeAPICall() (response ChannelAPIResponse, err error) {
	url, err := url.JoinPath(configs.YoutubeDataAPIBaseURL, configs.YouTubeDataAPIChannelsPath)
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
