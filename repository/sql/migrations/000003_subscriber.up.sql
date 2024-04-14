CREATE TABLE IF NOT EXISTS subscriber
(
    id    uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(250) NOT NULL
);