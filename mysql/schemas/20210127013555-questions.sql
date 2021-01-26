
-- +migrate Up
CREATE TABLE questions(
                          id int PRIMARY KEY AUTO_INCREMENT,
                          contents VARCHAR(100) NOT NULL,
                          is_answered BOOLEAN DEFAULT FALSE,
                          created_at DATETIME NOT NULL ,
                          updated_at DATETIME
);
-- +migrate Down
DROP TABLE questions;
