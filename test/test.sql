create table configuration(
    id int primary key auto_increment,
    name varchar(64),
    type tinyint,
    value text
)engine=innodb charset=utf8;

insert into configuration values (null, "test",1,"hello,test");