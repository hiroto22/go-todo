USE go_todo;

CREATE TABLE IF NOT EXISTS `users`(
     ID int NOT NULL AUTO_INCREMENT ,
     Name varchar(255),
     Email varchar(255),
     PassWord varchar(255),
     CreatedAt datetime,
     UpdatedAt datetime,
     PRIMARY KEY(id)
);

SELECT * FROM todos;


-- DESC users