INSERT INTO post
(content, newsletter_id)
VALUES
    (@content, newsletter_id)
RETURNING
    id, content, newsletter_id