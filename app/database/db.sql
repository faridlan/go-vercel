-- Active: 1662825965016@@128.199.67.28@3306@go_vercel

create table
    users (
        id int AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        age int NOT NULL
    ) engine = InnoDB;

SHOW tables;

select * from users;

INSERT INTO users (name,age) VALUES ("udin", 30);

select id,name,age from users;