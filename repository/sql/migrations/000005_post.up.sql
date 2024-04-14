CREATE TABLE IF NOT EXISTS post
(
    id            BIGINT PRIMARY KEY,
    content       TEXT NOT NULL,
    newsletter_id BIGINT REFERENCES newsletter (id)
);