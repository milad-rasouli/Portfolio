SELECT p.id as id,p.title as title, p.body as body,p.caption as caption,p.image_path as image_path,c.id as category_id , c.name as category FROM post as p
LEFT JOIN post_category_relation as pc ON pc.post_id = p.id
INNER JOIN category as c ON pc.category_id = c.id;


SELECT * FROM user WHERE email='bar@gmail.com' LIMIT 1 ;

