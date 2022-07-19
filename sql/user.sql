USE go_todo;

CREATE TABLE IF NOT EXISTS `users`(
     ID int NOT NULL AUTO_INCREMENT ,
     Name varchar(255) NOT NULL,
     Email varchar(255) NOT NULL,
     PassWord varchar(255) NOT NULL,
     CreatedAt datetime,
     UpdatedAt datetime,
     PRIMARY KEY(id)
);

SELECT * FROM users;


DESC users