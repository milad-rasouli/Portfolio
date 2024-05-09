SELECT p.id,p.title, p.body,p.caption,p.image_path, c.name FROM post as p
LEFT JOIN post_category_relation as pc ON pc.post_id = p.id
INNER JOIN category as c ON pc.category_id = c.id;


SELECT * FROM user WHERE email='bar@gmail.com' LIMIT 1 ;

