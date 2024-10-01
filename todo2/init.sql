CREATE DATABASE IF NOT EXISTS tasksdb;
USE tasksdb;

CREATE TABLE `tasks` (
  `id` int NOT NULL AUTO_INCREMENT,
  `description` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO `tasks` VALUES (1,'my first task');