drop database message_board_database;
create database message_board_database;
use message_board_database;
create table userdata
(
    uid      int primary key auto_increment,
    username varchar(20),
    password varchar(20)
);
create table messages
(
    mid      int primary key auto_increment,
    sender   varchar(20) not null,
    receiver varchar(20) not null,
    time     time        not null,
    text     text        not null,
    deleted  varchar(10) default '0'
);

insert into userdata(uid, username, password)
values (1, 'xiaote11', '123456');
insert into userdata(uid, username, password)
values (null, 'xiaote22', '123456');
insert into userdata(uid, username, password)
values (null, 'xiaote33', '123456');
insert into userdata(uid, username, password)
values (null, 'xiaote44', '123456');
insert into userdata(uid, username, password)
values (null, 'xiaote55', '123456');
insert into userdata(uid, username, password)
values (null, 'xiaote66', '123456');
insert into userdata(uid, username, password)
values (null, 'xiaote77', '123456');

insert into messages(mid, sender, receiver, time, text)
values (1, 'xiaote11', 'xiaote33', '09:32:33', 'I love U');
insert into messages(mid, sender, receiver, time, text)
values (null, 'xiaote33', 'xiaote11', '09:32:34', 'I love U, too');
insert into messages(mid, sender, receiver, time, text)
values (null, 'xiaote11', 'xiaote33', '09:32:35', 'really');
insert into messages(mid, sender, receiver, time, text)
values (null, 'xiaote33', 'xiaote11', '09:32:36', 'yes');
insert into messages(mid, sender, receiver, time, text)
values (null, 'xiaote33', 'xiaote11', '09:32:37', 'I truly love you');


create table comments
(
    mid     int         not null,
    cid     int primary key auto_increment,
    sender  varchar(20) not null,
    time    time        not null,
    text    text        not null,
    deleted varchar(10) default '0',
    rid     int         default '0'
);
insert into comments(mid, cid, sender, time, text)
values (1, 1, 'xiaote22', '10:14:22', '99');
insert into comments(mid, cid, sender, time, text)
values (1, null, 'xiaote44', '10:14:23', '9999');
insert into comments(mid, cid, sender, time, text)
values (1, null, 'xiaote55', '10:14:24', '99999');
insert into comments(mid, cid, sender, time, text)
values (1, null, 'xiaote66', '10:14:25', '999999');
insert into comments(mid, cid, sender, time, text)
values (1, null, 'xiaote77', '10:14:26', '9999999');



insert into comments(mid, sender, time, text, rid)
values (1,'xiaote','10:14:27','6',1);
insert into comments(mid, sender, time, text, rid)
values (1,'xiaote33','10:14:28','666',6);
insert into comments(mid, sender, time, text, rid)
values (1,'xiaote33','10:14:29','666',6);
insert into comments(mid, sender, time, text, rid)
values (1,'xiaote','10:14:30','6666',8);
insert into comments(mid, sender, time, text, rid)
values (1,'xiaote','10:14:31','6666',8);
insert into comments(mid, sender, time, text, rid)
values (1,'xiaote33','10:14:32','666666',10);
insert into comments(mid, sender, time, text, rid)
values (1,'xiaote33','10:14:33','666666',10);
