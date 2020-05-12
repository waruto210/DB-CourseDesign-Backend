create table class_info
(
    class_no varchar(255) not null
        primary key
);

create table user
(
    user_id   varchar(255)     not null
        primary key,
    user_type tinyint unsigned null,
    passwd    varbinary(255)   null
);

create table admin
(
    admin_no varchar(255) not null,
    constraint admin_admin_no_uindex
        unique (admin_no),
    constraint admin_user_user_id_fk
        foreign key (admin_no) references user (user_id)
);

create table student_info
(
    stu_no   varchar(255) not null,
    stu_name varchar(255) null,
    class_no varchar(255) null,
    constraint student_info_stu_no_uindex
        unique (stu_no),
    constraint student_info_user_user_id_fk
        foreign key (stu_no) references user (user_id)
            on delete cascade
);

create table teacher_info
(
    tea_no   varchar(255) not null,
    tea_name varchar(255) null,
    constraint teacher_info_tea_no_uindex
        unique (tea_no),
    constraint teacher_info_user_user_id_fk
        foreign key (tea_no) references user (user_id)
            on delete cascade
);

create table course_info
(
    course_no   varchar(255) not null
        primary key,
    course_name varchar(255) null,
    tea_no      varchar(255) null,
    constraint course_info_teacher_info_tea_no_fk
        foreign key (tea_no) references teacher_info (tea_no)
            on delete cascade
);

create table student_course
(
    stu_no    varchar(255) not null,
    course_no varchar(255) not null,
    score     int          null,
    constraint course_no
        unique (course_no, stu_no),
    constraint student_course_course_info_course_no_fk
        foreign key (course_no) references course_info (course_no)
            on delete cascade,
    constraint student_course_student_info_stu_no_fk
        foreign key (stu_no) references student_info (stu_no)
            on delete cascade
);


