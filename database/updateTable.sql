UPDATE post
SET title = 'Updated Title', body = 'Updated Body'
WHERE id = 1;

UPDATE category
SET name = 'Updated Category'
WHERE id = 3;

UPDATE post_category_relation
SET category_id = 4
WHERE post_id = 1 AND category_id = 2;

