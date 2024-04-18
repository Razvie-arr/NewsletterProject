SELECT
    s.id,
    s.email
FROM
    subscriber as s
WHERE
    email = @email
