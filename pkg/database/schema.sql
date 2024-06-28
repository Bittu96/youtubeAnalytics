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

CREATE TABLE youtubeAnalytics.channelInsights (
    channel_insights_id VARCHAR(64) PRIMARY KEY,
    channel_id varchar(64),
    subscriber_count_inc BIGINT,
    view_count_inc BIGINT,
    top_keywords JSONB,
    row_created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    -- CONSTRAINT fk_channel_statistics_channel FOREIGN KEY (channel_id) REFERENCES youtubeAnalytics.channel (channel_id) ON DELETE CASCADE
);

CREATE TABLE youtubeAnalytics.videoInsights (
    video_insights_id varchar(64) PRIMARY KEY,
    video_id varchar(64),
    view_count_inc BIGINT,
    view_count_inc_perc FLOAT,
    like_count_inc BIGINT,
    like_count_inc_perc FLOAT,
    comment_count_inc BIGINT,
    comment_count_inc_perc FLOAT,
    total_impressions BIGINT,
    total_impressions_perc FLOAT,
    row_created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- DROP TABLE youtubeAnalytics.channel;
-- DROP TABLE youtubeAnalytics.playlist;
-- DROP TABLE youtubeAnalytics.video;
-- DROP TABLE youtubeAnalytics.videoStatistics;
-- DROP TABLE youtubeAnalytics.channelStatistics;
-- DROP TABLE youtubeAnalytics.videoInsights;

-- select * from youtubeanalytics.channelStatistics;
-- select channel_id, details->>'title' from youtubeanalytics.channel c;
-- select video_id, last_value(view_count), first_value(view_count), percent from videostatistics where row_created > now() - interval '24 hours' group by video_id;