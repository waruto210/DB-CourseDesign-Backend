create database student_score
DEFAULT CHARACTER
SET utf8 DEFAULT
COLLATE utf8_general_ci;
create user 'db_class' @'localhost' identified by 'dbclassmm';
grant all privileges on student_score.* to 'db_class' @'localhost';

create table
if not exists room
(
    room_id varchar
(31) not null
        primary key
);

create table
if not exists operation_log
(
    op_id   int auto_increment
        primary key,
    room_id varchar
(31) null,
    time    datetime    null,
    action  int         null,
    constraint operation_log_room_roomid_fk
        foreign key
(room_id) references room
(room_id)
            on
update cascade on
delete cascade
);

create table
if not exists `usage`
(
    usage_id   int auto_increment
        primary key,
    room_id    varchar
(31) null,
    start_time datetime    not null,
    end_time   datetime    not null,
    speed      tinyint     not null,
    fee_rate   float       not null,
    constraint usage_room_room_id_fk
        foreign key
(room_id) references room
(room_id)
            on
update cascade on
delete cascade
);

