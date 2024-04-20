INSERT INTO subscriber
    (email)
VALUES
    (@email)
RETURNING
    id, email