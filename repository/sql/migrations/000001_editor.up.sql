CREATE TABLE IF NOT EXISTS editor
(
    id       uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    email    VARCHAR(250) NOT NULL,
    password VARCHAR(250) NOT NULL
);
