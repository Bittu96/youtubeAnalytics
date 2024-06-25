-- Active: 1710401727370@@127.0.0.1@5432@postgres@youtubeanalytics
CREATE SCHEMA youtubeAnalytics;

CREATE TABLE youtubeAnalytics.channel (
    channel_id VARCHAR(64) PRIMARY KEY,
    details JSONB,
    row_created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE youtubeAnalytics.playlist (
    playlist_id VARCHAR(64) PRIMARY KEY,
    channel_id VARCHAR(64) NOT NULL,
    details JSONB,
    row_created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE youtubeAnalytics.video (
    video_id varchar(64) PRIMARY KEY,
    channel_id VARCHAR(64) NOT NULL,
    details JSONB,
    row_created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE youtubeAnalytics.videoStatistics (
    video_statistics_id varchar(64) PRIMARY KEY,
    video_id varchar(64),
    view_count BIGINT,
    like_count BIGINT,
    comment_count BIGINT,
    row_created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE youtubeAnalytics.channelStatistics (
    channel_statistics_id VARCHAR(64) PRIMARY KEY,
    channel_id varchar(64),
    subscriber_count BIGINT,
    view_count BIGINT,
    video_count BIGINT,
    row_created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    row_updated TIMESTAMP
    -- CONSTRAINT fk_channel_statistics_channel FOREIGN KEY (channel_id) REFERENCES youtubeAnalytics.channel (channel_id) ON DELETE CASCADE
);

select * from youtubeanalytics.channelStatistics;
select channel_id, details->>'title' from youtubeanalytics.channel c;

-- DROP TABLE youtubeAnalytics.channel;
-- DROP TABLE youtubeAnalytics.playlist;
-- DROP TABLE youtubeAnalytics.video;
-- DROP TABLE youtubeAnalytics.videoStatistics;
-- DROP TABLE youtubeAnalytics.channelStatistics;