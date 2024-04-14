CREATE TABLE IF NOT EXISTS newsletter
(
    id          BIGINT PRIMARY KEY,
    name        VARCHAR(250) NOT NULL,
    description VARCHAR(250),
    editor_id   BIGINT REFERENCES editor (id)
);
