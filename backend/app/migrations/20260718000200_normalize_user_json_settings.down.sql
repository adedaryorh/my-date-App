ALTER TABLE users
    ALTER COLUMN push_notification_settings DROP NOT NULL,
    ALTER COLUMN push_notification_settings SET DEFAULT NULL,
    ALTER COLUMN content_settings DROP NOT NULL,
    ALTER COLUMN content_settings SET DEFAULT NULL;
