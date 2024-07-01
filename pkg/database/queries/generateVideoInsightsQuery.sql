select t.video_id, 
    min(t.views) as views, 
    max(t.views_after_24h) as views_after_24h,
    min(t.likes) as likes, 
    max(t.likes_after_24h) as likes_after_24h,
    min(t.comments) as comments, 
    max(t.comments_after_24h) as comments_after_24h
from (
    select
    distinct(video_id),
    first_value(view_count) over (PARTITION BY video_id ORDER BY row_created) as views,
    last_value(view_count) over (PARTITION BY video_id ORDER BY row_created) as views_after_24h,
    first_value(like_count) over (PARTITION BY video_id ORDER BY row_created) as likes,
    last_value(like_count) over (PARTITION BY video_id ORDER BY row_created) as likes_after_24h,
    first_value(comment_count) over (PARTITION BY video_id ORDER BY row_created) as comments,
    last_value(comment_count) over (PARTITION BY video_id ORDER BY row_created) as comments_after_24h
    from youtubeAnalytics.videostatistics
    where row_created > now() - interval '24 hours'
     ORDER BY video_id) as t
group by t.video_id;