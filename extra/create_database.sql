create database student_score;
create user 'db_class'@'localhost' identified by 'dbclassmm';
grant all privileges on student_score.* to 'db_class'@'localhost';
