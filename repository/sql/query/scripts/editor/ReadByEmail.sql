SELECT
	e.id,
	e.email,
	e.password
FROM
	editor as e
WHERE
	email = @email
