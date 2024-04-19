INSERT INTO editor (email, password)
VALUES (@email, @password) RETURNING id, email, password;
