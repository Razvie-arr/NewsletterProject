SELECT verification_string FROM newsletter_subscriber
WHERE subscriber_id = @subscriber_id
AND newsletter_id = @newsletter_id