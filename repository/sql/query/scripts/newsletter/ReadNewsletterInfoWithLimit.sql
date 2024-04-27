SELECT n.id,
       n."name",
       n.description,
       e.id    AS "editor_id",
       e.email AS "editor_email"
FROM newsletter n
         JOIN editor e ON e.id = n.editor_id
limit $1;