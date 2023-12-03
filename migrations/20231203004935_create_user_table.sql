CREATE SCHEMA users;

CREATE TABLE users (
    id uuid PRIMARY KEY,
    login VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    fullname VARCHAR(255) NOT NULL
);

DROP TABLE users;
DROP SCHEMA users;

