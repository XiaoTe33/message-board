drop database message_board_database;
create database message_board_database;
use message_board_database;
create table userdata(
    uid int primary key auto_increment,
    username varchar(20),
    password varchar(20)
);
create table messages(
    mid int primary key auto_increment,
    sender varchar(20) not null ,
    receiver varchar(20) not null ,
    time time not null ,
    text text not null ,
    deleted varchar(10) default '0'
);

insert into userdata(uid, username, password) values (1,'xiaote','1');
insert into userdata(uid, username, password) values (null,'xiaote33','1');

insert into messages(mid, sender, receiver, time, text)
values (1, 'xiaote','xiaote33','09:32:33','I love U');
insert into messages(mid, sender, receiver, time, text)
values (2, 'xiaote33','xiaote','09:32:34','I love U, too');

create table comments(
    mid int not null ,
    cid int primary key auto_increment,
    sender varchar(20) not null ,
    time time not null ,
    text text not null,
    deleted varchar(10) default '0',
    rid int default '0'
);
insert into comments(mid, cid, sender,  time, text)
values (1,1,'xiaote33','10:14:22','I love U, too');
