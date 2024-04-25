INSERT INTO editor
    (id, email)
VALUES
    (@id, @email)
RETURNING
    id, email