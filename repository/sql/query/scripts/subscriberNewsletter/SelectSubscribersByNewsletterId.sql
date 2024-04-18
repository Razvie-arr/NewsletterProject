SELECT
    s.id,
    s.email
FROM
    newsletter_subscriber ns
LEFT JOIN subscriber as s
    ON ns.subscriber_id = s.id
WHERE
    ns.newsletter_id = @newsletter_id