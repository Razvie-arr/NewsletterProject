SELECT
    n.id,
    n.name,
    n.description,
    n.editor_id
FROM newsletter as n
WHERE n.id = @id