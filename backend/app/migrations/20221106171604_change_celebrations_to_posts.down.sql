ALTER TABLE posts
    RENAME COLUMN post_id TO celebration_id;

ALTER TABLE posts
    RENAME TO celebrations;

ALTER TABLE posts_media
    RENAME COLUMN post_id TO celebration_id;

ALTER TABLE posts_media
    RENAME TO celebrations_media;