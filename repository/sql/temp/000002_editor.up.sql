CREATE TABLE IF NOT EXISTS Editor
(
    id       BIGINT PRIMARY KEY,
    email    VARCHAR(250) NOT NULL,
    password VARCHAR(250) NOT NULL
);
