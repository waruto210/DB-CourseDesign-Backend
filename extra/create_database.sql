create database student_score DEFAULT CHARACTER SET utf8 DEFAULT COLLATE utf8_general_ci;
create user 'db_class'@'localhost' identified by 'dbclassmm';
grant all privileges on student_score.* to 'db_class'@'localhost';
