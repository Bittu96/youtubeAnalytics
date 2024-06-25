package channel

import (
	"strconv"
	"sync"
	"youtubeAnalytics/models/playlist"
	"youtubeAnalytics/pkg/rmq"
	channelAPI "youtubeAnalytics/pkg/youtube/apis/channel"
	playlistCollecionAPI "youtubeAnalytics/pkg/youtube/apis/playlistCollection"
)

type channel struct {
	ChannelID       string              `json:"channel_id"`
	SubscriberCount int64               `json:"subscriber_count"`
	VideoCount      int64               `json:"video_count"`
	ViewCount       int64               `json:"view_count"`
	Details         channelAPI.Snippet  `json:"details"`
	Playlists       []playlist.Playlist `json:"playlists"`
}

func New(channelId string) channel {
	return channel{
		ChannelID: channelId,
	}
}

func (c channel) Load() (channel, error) {
	channelAPIResponse, err := channelAPI.Request(c.ChannelID).MakeAPICall()
	if err != nil {
		return channel{}, err
	} else if len(channelAPIResponse.Items) == 0 {
		return channel{}, nil
	} else {
		for _, channelItem := range channelAPIResponse.Items {
			c.Details = channelItem.Snippet
			if !channelItem.Statistics.HiddenSubscriberCount {
				c.SubscriberCount, _ = strconv.ParseInt(channelItem.Statistics.SubscriberCount, 10, 64)
				c.VideoCount, _ = strconv.ParseInt(channelItem.Statistics.VideoCount, 10, 64)
				c.ViewCount, _ = strconv.ParseInt(channelItem.Statistics.ViewCount, 10, 64)
			}
			if err := rmq.RMQPublisherClient.Publish("channel", c); err != nil {
				return channel{}, err
			}
		}
	}

	if err = c.UploadsDownloader(channelAPIResponse.Items[0].ContentDetails.RelatedPlaylists.Uploads); err != nil {
		return channel{}, err
	}

	if err = c.loadAllPlaylists(); err != nil {
		return channel{}, err
	}
	return c, nil
}

func (c *channel) loadAllPlaylists() error {
	wg := &sync.WaitGroup{}
	mux := &sync.Mutex{}
	var nextPageToken string

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

	for nextPageToken != "" {
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

func (c *channel) PlaylistDownloader(wg *sync.WaitGroup, mux *sync.Mutex, playlistId string) {
	defer func() { wg.Done() }()
	newPlaylist, err := playlist.New(playlistId).Load()
	if err != nil {
		return
	}

	c.Playlists = append(c.Playlists, newPlaylist)
}

func (c *channel) UploadsDownloader(playlistId string) error {
	newPlaylist, err := playlist.New(playlistId).Load()
	if err != nil {
		return err
	}

	c.Playlists = append(c.Playlists, newPlaylist)
	return nil
}
