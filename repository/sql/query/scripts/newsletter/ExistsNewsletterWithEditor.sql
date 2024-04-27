SELECT EXISTS (
    SELECT 1
    FROM newsletter n
    WHERE n.editor_id = @editorId
      AND n.id = @id
);