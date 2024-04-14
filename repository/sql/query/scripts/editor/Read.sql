SELECT
	e.id,
	e.email,
	e.password
FROM
	editor as e
WHERE
	id = @id
