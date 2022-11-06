ALTER TABLE celebrations
    RENAME COLUMN celebration_id TO post_id;

ALTER TABLE celebrations
    RENAME TO posts;

ALTER TABLE celebrations_media
    RENAME COLUMN celebration_id TO post_id;

ALTER TABLE celebrations_media
    RENAME TO posts_media;
