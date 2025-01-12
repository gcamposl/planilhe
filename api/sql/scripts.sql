CREATE DATABASE IF NOT EXISTS planilhe;
USE planilhe;


DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    nick VARCHAR(255) NOT NULL unique,
    email VARCHAR(255) NOT NULL unique,
    password VARCHAR(255) NOT NULL,
    created_at_dt TIMESTAMP DEFAULT current_timestamp(),
) ENGINE=InnoDB;