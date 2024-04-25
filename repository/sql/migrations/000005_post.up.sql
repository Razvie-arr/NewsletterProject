CREATE TABLE IF NOT EXISTS post
(
    id            uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    content       TEXT NOT NULL,
    newsletter_id uuid REFERENCES newsletter (id) ON DELETE CASCADE
);