USE questionner_db;
CREATE TABLE questions(
                      id int PRIMARY KEY AUTO_INCREMENT,
                      title VARCHAR(50) NOT NULL,
                      description VARCHAR(100) NOT NULL
);