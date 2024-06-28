package channel

import (
	"strconv"
	"sync"
	"youtubeAnalytics/configs"
	"youtubeAnalytics/models/playlist"
	"youtubeAnalytics/pkg/rmq"
	channelAPI "youtubeAnalytics/pkg/youtube/apis/channel"
	playlistCollecionAPI "youtubeAnalytics/pkg/youtube/apis/playlistCollection"
)

type Channel struct {
	ChannelID       string              `json:"channel_id"`
	SubscriberCount int64               `json:"subscriber_count"`
	VideoCount      int64               `json:"video_count"`
	ViewCount       int64               `json:"view_count"`
	Details         channelAPI.Snippet  `json:"details"`
	Playlists       []playlist.Playlist `json:"playlists"`
}

// create new channel
func New(channelId string) Channel {
	return Channel{
		ChannelID: channelId,
	}
}

// download channel info and all it's uploads/plylists
func (c Channel) Download() (Channel, error) {
	// calling youtube Data API
	channelAPIResponse, err := channelAPI.Request(c.ChannelID).MakeAPICall()
	if err != nil {
		return Channel{}, err
	} else if len(channelAPIResponse.Items) == 0 {
		return Channel{}, nil
	} else {
		for _, channelItem := range channelAPIResponse.Items {
			c.Details = channelItem.Snippet
			if !channelItem.Statistics.HiddenSubscriberCount {
				c.SubscriberCount, _ = strconv.ParseInt(channelItem.Statistics.SubscriberCount, 10, 64)
				c.VideoCount, _ = strconv.ParseInt(channelItem.Statistics.VideoCount, 10, 64)
				c.ViewCount, _ = strconv.ParseInt(channelItem.Statistics.ViewCount, 10, 64)
			}

			// push channel info to consumer
			if err := rmq.RMQPublisherClient.Publish("channel", c); err != nil {
				return Channel{}, err
			}
		}
	}

	// download all videos from only channel uploads section by default
	if err = c.UploadsDownloader(channelAPIResponse.Items[0].ContentDetails.RelatedPlaylists.Uploads); err != nil {
		return Channel{}, err
	}

	// download all videos from playlist sections if enabled
	if configs.LoadVideosFromPlaylists {
		if err = c.downloadAllPlaylists(); err != nil {
			return Channel{}, err
		}
	}

	return c, nil
}

func (c *Channel) downloadAllPlaylists() error {
	wg := &sync.WaitGroup{}
	mux := &sync.Mutex{}
	var nextPageToken string

	// calling youtube Data API
	playlistCollectionAPIResponse, err := playlistCollecionAPI.Request(c.ChannelID).MakeAPICall()
	if err != nil {
		return err
	} else {
		nextPageToken = playlistCollectionAPIResponse.NextPageToken
		for _, playlistItem := range playlistCollectionAPIResponse.Items {
			wg.Add(1)
			go c.PlaylistDownloader(wg, mux, playlistItem.ID)
		}
		wg.Wait()
	}

	// if multiple pages exist
	for nextPageToken != "" {
		// calling youtube Data API for next page
		playlistCollectionAPIResponse, err := playlistCollecionAPI.Request(c.ChannelID).GetPage(nextPageToken).MakeAPICall()
		if err != nil {
			return err
		} else {
			nextPageToken = playlistCollectionAPIResponse.NextPageToken
			for _, playlistItem := range playlistCollectionAPIResponse.Items {
				wg.Add(1)
				go c.PlaylistDownloader(wg, mux, playlistItem.ID)
			}
			wg.Wait()
		}
	}

	return nil
}

// download playlist data
func (c *Channel) PlaylistDownloader(wg *sync.WaitGroup, mux *sync.Mutex, playlistId string) {
	defer func() { wg.Done() }()
	newPlaylist, err := playlist.New(playlistId).Download()
	if err != nil {
		return
	}

	c.Playlists = append(c.Playlists, newPlaylist)
}

// download uploads data
func (c *Channel) UploadsDownloader(playlistId string) error {
	newPlaylist, err := playlist.New(playlistId).Download()
	if err != nil {
		return err
	}

	c.Playlists = append(c.Playlists, newPlaylist)
	return nil
}
