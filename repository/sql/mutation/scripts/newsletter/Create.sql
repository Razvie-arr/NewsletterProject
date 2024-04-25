INSERT INTO newsletter (name, description, editor_id)
VALUES (@name, @description, @editor_id)
RETURNING id, name, description