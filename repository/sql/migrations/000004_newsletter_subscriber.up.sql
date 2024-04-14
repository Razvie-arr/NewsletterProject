CREATE TABLE IF NOT EXISTS newsletter_subscriber
(
    subscriber_id       BIGINT REFERENCES subscriber (id),
    newsletter_id       BIGINT REFERENCES newsletter (id),
    verification_string VARCHAR(100) NOT NULL,
    PRIMARY KEY (subscriber_id, newsletter_id)
);