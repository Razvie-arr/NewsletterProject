CREATE TABLE IF NOT EXISTS newsletter_subscriber
(
    subscriber_id       uuid REFERENCES subscriber (id) ON DELETE CASCADE,
    newsletter_id       uuid REFERENCES newsletter (id) ON DELETE CASCADE,
    verification_string VARCHAR(100) NOT NULL,
    PRIMARY KEY (subscriber_id, newsletter_id)
);