Alter Table "celebrations" ADD column caption text not Null;
Alter Table "celebrations" ADD column celebration_kind varchar(256) not Null;
Alter Table "celebrations" ADD column celebration_date timestamp not Null;

Alter Table "media" ADD column owner_id varchar(256) not Null;

