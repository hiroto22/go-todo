USE go_todo;

CREATE TABLE IF NOT EXISTS `todos`(
     ID int NOT NULL AUTO_INCREMENT ,
     UserID int NOT NULL,
     Todo text NOT NULL,
     CreatedAt datetime,
     UpdatedAt datetime,
     PRIMARY KEY(id)
);

DESC todos