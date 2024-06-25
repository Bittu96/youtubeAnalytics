package main

import (
	"youtubeAnalytics/services"
)

func main() {
	// systemctl start rabbitmq-server
	// rmq.RMQPublisherClient = rmq.New(configs.RMQURL, configs.QueueName)
	// if err := rmq.RMQPublisherClient.Connect(); err != nil {
	// 	log.Fatal(err)
	// 	return
	// } else {
	// 	defer rmq.RMQPublisherClient.Close()
	// }

	// targetChannels := []string{s
	// 	"UC5OrDvL9DscpcAstz7JnQGA",
	// 	"UC70pKToywlxOGdgIvz8gYqA",
	// }

	services.ExtractKeywords(`{"tags": ["Dance Pop", "Energetic", "Feeling Good", "Happy", "Love Never Felt So Good", "Michael Jackson", "Romantic", "Upbeat"], "title": "Michael Jackson - Love Never Felt So Good (Official Video)", "channelId": "UCulYu1HEIa7f70L2lYZWHOw", "localized": {"title": "Michael Jackson - Love Never Felt So Good (Official Video)", "description": "The first track on Michael Jackson’s album, XSCAPE,  is “Love Never Felt So Good” which was “contemporized” by producer (and Michael Jackson Estate Co-Executor) John McClain.  This version of the video for the song celebrates Michael and a recording which makes everyone who hears it want to get up and dance just like all of the dancers featured in the video.  \n\nDownload Xscape on iTunes Now: http://smarturl.it/xscape?IQid=youtube\nDownload Xscape on Amazon Now: http://smarturl.it/xscape-amazonmp3\nAudio stream also available now at Music Unlimited Music Unlimited\n\nXSCAPE is an album of previously unreleased Michael Jackson songs. The album is produced and curated by Epic Records Chairman and CEO L.A. Reid, who retooled the production to add a fresh, contemporary sound that retains Jackson's essence and integrity. It's a process Reid calls \"contemporizing.\" The list of producers include global hitmakers Timbaland, Rodney Jerkins, Stargate, and John McClain.\n\nFor more information on the album, go to www.michaeljackson.com.\n\nFacebook.com/MichaelJackson\nTwitter: @MichaelJackson\nInstagram: @MichaelJackson\nGoogle+: +MichaelJackson\n\nMusic video by Michael Jackson performing Love Never Felt So Good. (C) 2014 MJJ Productions, Inc."}, "categoryId": "10", "description": "The first track on Michael Jackson’s album, XSCAPE,  is “Love Never Felt So Good” which was “contemporized” by producer (and Michael Jackson Estate Co-Executor) John McClain.  This version of the video for the song celebrates Michael and a recording which makes everyone who hears it want to get up and dance just like all of the dancers featured in the video.  \n\nDownload Xscape on iTunes Now: http://smarturl.it/xscape?IQid=youtube\nDownload Xscape on Amazon Now: http://smarturl.it/xscape-amazonmp3\nAudio stream also available now at Music Unlimited Music Unlimited\n\nXSCAPE is an album of previously unreleased Michael Jackson songs. The album is produced and curated by Epic Records Chairman and CEO L.A. Reid, who retooled the production to add a fresh, contemporary sound that retains Jackson's essence and integrity. It's a process Reid calls \"contemporizing.\" The list of producers include global hitmakers Timbaland, Rodney Jerkins, Stargate, and John McClain.\n\nFor more information on the album, go to www.michaeljackson.com.\n\nFacebook.com/MichaelJackson\nTwitter: @MichaelJackson\nInstagram: @MichaelJackson\nGoogle+: +MichaelJackson\n\nMusic video by Michael Jackson performing Love Never Felt So Good. (C) 2014 MJJ Productions, Inc.", "publishedAt": "2014-06-19T23:00:02Z", "channelTitle": "michaeljacksonVEVO", "defaultAudioLanguage": "en-US", "liveBroadcastContent": "none"}`)
	// services.GetAllVideos(targetChannels)
}
