-- Create the 'post' table
CREATE TABLE IF NOT EXISTS post (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    body TEXT,
    image_path TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    modified_at DATETIME 
);

-- Create the 'category' table
CREATE TABLE IF NOT EXISTS category (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL
);

-- Create the 'post_category_relation' table

CREATE TABLE IF NOT EXISTS "post_category_relation" (
	"category_id"	INTEGER,
	"post_id"	INTEGER,
	PRIMARY KEY("category_id","post_id"),
	FOREIGN KEY("category_id") REFERENCES "category"("id"),
	FOREIGN KEY("post_id") REFERENCES "post"("id")
);