CREATE TABLE IF NOT EXISTS newsletter
(
    id          uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name        VARCHAR(250) NOT NULL,
    description VARCHAR(250),
    editor_id   uuid REFERENCES editor (id)
);
