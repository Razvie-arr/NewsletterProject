UPDATE newsletter
SET "name" = @name,
description = @description
WHERE id = @id
RETURNING id, "name", description;