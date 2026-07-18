Alter Table "users" DROP column push_notification_settings;
Alter Table "users" DROP column content_settings;
Alter Table "users" DROP column banned_words;

drop table followers;
drop table blocked;
