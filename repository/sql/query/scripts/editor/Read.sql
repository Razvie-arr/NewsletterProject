SELECT
	e.id,
	e.email
FROM
	editor as e
WHERE
	id = @id
