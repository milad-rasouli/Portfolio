-- Insert sample data
INSERT INTO post (title, body, image_path)
VALUES ('Redis In ACTION', 'dummy dummy text', 'dummy image path'),
       ('Golang In ACTION', 'dummy dummy text', 'dummy image path');

INSERT INTO category (name)
VALUES ('Database'), ('Message Broker'), ('Golang'), ('Programming'), ('Microservices');

INSERT INTO post_category_relation (post_id, category_id)
VALUES (1, 1), (1, 2), (1, 5),(2, 3),(2, 4);