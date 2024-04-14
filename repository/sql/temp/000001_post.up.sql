CREATE TABLE IF NOT EXISTS Post
(
    id            BIGINT PRIMARY KEY,
    content       TEXT NOT NULL,
    newsletter_id BIGINT REFERENCES Newsletter (id)
);