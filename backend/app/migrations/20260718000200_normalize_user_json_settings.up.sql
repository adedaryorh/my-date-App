UPDATE users
SET push_notification_settings = '{}'::jsonb
WHERE push_notification_settings IS NULL;

UPDATE users
SET content_settings = '{}'::jsonb
WHERE content_settings IS NULL;

ALTER TABLE users
    ALTER COLUMN push_notification_settings SET DEFAULT '{}'::jsonb,
    ALTER COLUMN push_notification_settings SET NOT NULL,
    ALTER COLUMN content_settings SET DEFAULT '{}'::jsonb,
    ALTER COLUMN content_settings SET NOT NULL;

