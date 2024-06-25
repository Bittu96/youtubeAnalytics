# Youtube Analytics

A brief description of what this project does and who it's for

## Tech Stack

**Server:** Golang, Postgres, RabbitMQ

## Run Locally

Clone the project

```bash
  git clone https://link-to-project
```

Go to the project directory

```bash
  cd my-project
```

Install dependencies

```bash
  npm install
```

Start the server

```bash
  npm run start
```

## Additional References

### To get channels list :

Get Channels list by forUserName:
https://www.googleapis.com/youtube/v3/channels?part=snippet,contentDetails,statistics&forUsername=Apple&key=

Get channels list by channel id:
https://www.googleapis.com/youtube/v3/channels/?part=snippet,contentDetails,statistics&id=UCE_M8A5yxnLfW0KghEeajjw&key=

Get Channel sections:
https://www.googleapis.com/youtube/v3/channelSections?part=snippet,contentDetails&channelId=UCE_M8A5yxnLfW0KghEeajjw&key=

### To get Playlists :

Get Playlists by Channel ID:
https://www.googleapis.com/youtube/v3/playlists?part=snippet,contentDetails&channelId=UCq-Fj5jknLsUf-MWSy4_brA&maxResults=50&key=

Get Playlists by Channel ID with pageToken:
https://www.googleapis.com/youtube/v3/playlists?part=snippet,contentDetails&channelId=UCq-Fj5jknLsUf-MWSy4_brA&maxResults=50&key=&pageToken=CDIQAA

### To get PlaylistItems :

Get PlaylistItems list by PlayListId:
https://www.googleapis.com/youtube/v3/playlistItems?part=snippet,contentDetails&maxResults=25&playlistId=PLHFlHpPjgk70Yv3kxQvkDEO5n5tMQia5I&key=

### To get videos :

Get videos list by video id:
https://www.googleapis.com/youtube/v3/videos?part=snippet,contentDetails,statistics&id=YxLCwfA1cLw&key=

Get videos list by multiple videos id:
https://www.googleapis.com/youtube/v3/videos?part=snippet,contentDetails,statistics&id=YxLCwfA1cLw,Qgy6LaO3SB0,7yPJXGO2Dcw&key=

### Get comments list

Get Comment list by video ID:
https://www.googleapis.com/youtube/v3/commentThreads?part=snippet,replies&videoId=shdhddsh

Get Comment list by channel ID:
https://www.googleapis.com/youtube/v3/commentThreads?part=snippet,replies&channelId=srshsrhdhdhd

Get Comment list by allThreadsRelatedToChannelId:
https://www.googleapis.com/youtube/v3/commentThreads?part=snippet,replies&allThreadsRelatedToChannelId=rshdhdhdh

## License

[MIT](https://choosealicense.com/licenses/mit/)