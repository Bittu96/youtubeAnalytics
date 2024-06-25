package playlist

import (
	"log"
	"sync"
	"youtubeAnalytics/models/video"
	"youtubeAnalytics/pkg/rmq"
	playlistItemAPI "youtubeAnalytics/pkg/youtube/apis/playlistItem"
)

type Playlist struct {
	PlaylistID string                  `json:"playlist_id"`
	ChannelID  string                  `json:"channel_id"`
	VideoCount int64                   `json:"video_count"`
	Details    playlistItemAPI.Snippet `json:"details"`
	Videos     []video.Video           `json:"videos"`
}

func New(playlistId string) Playlist {
	return Playlist{
		PlaylistID: playlistId,
	}
}

func (p Playlist) Load() (Playlist, error) {
	var nextPageToken string
	playlistItemAPIResponse, err := playlistItemAPI.Request(p.PlaylistID).MakeAPICall()
	if err != nil {
		return Playlist{}, err
	} else if len(playlistItemAPIResponse.Items) == 0 {
		return Playlist{}, nil
	} else {
		nextPageToken = playlistItemAPIResponse.NextPageToken
		p.VideoCount = playlistItemAPIResponse.PageInfo.TotalResults
		p.Details = playlistItemAPIResponse.Items[0].Snippet
		p.ChannelID = playlistItemAPIResponse.Items[0].Snippet.ChannelID
		if err := rmq.RMQPublisherClient.Publish("playlist", p); err != nil {
			log.Println(err)
		}
		p.loadAllVideos(playlistItemAPIResponse.Items)
	}

	for nextPageToken != "" {
		playlistItemAPIResponse, err := playlistItemAPI.Request(p.PlaylistID).SetPage(nextPageToken).MakeAPICall()
		if err != nil {
			return Playlist{}, err
		} else {
			nextPageToken = playlistItemAPIResponse.NextPageToken
			p.loadAllVideos(playlistItemAPIResponse.Items)
		}
	}

	if p.VideoCount != int64(len(p.Videos)) {
		log.Println("playlist couldn't fetch all videos", p.VideoCount, len(p.Videos))
	}

	return p, nil
}

func (p *Playlist) loadAllVideos(playlistItems []playlistItemAPI.Item) {
	wg := &sync.WaitGroup{}
	mux := &sync.Mutex{}
	for _, playlistVideoData := range playlistItems {
		wg.Add(1)
		go p.videoDownloader(wg, mux, playlistVideoData.ContentDetails.VideoID)
	}
	wg.Wait()
}

func (p *Playlist) videoDownloader(wg *sync.WaitGroup, mux *sync.Mutex, videoId string) {
	defer func() {
		wg.Done()
	}()

	videoData, err := video.New(videoId).Load()
	if err != nil {
		return
	} else {
		mux.Lock()
		p.Videos = append(p.Videos, videoData)
		mux.Unlock()
	}
}
