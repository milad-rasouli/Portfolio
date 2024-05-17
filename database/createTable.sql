CREATE TABLE user (
    id INTEGER,
    full_name TEXT NOT NULL,
    email INTEGER NOT NULL UNIQUE,
    password TEXT NOT NULL,
    is_github INTEGER DEFAULT 0,
    online_at DATETIME,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    modified_at DATETIME,
    PRIMARY KEY(id)
);
CREATE TABLE IF NOT EXISTS post (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    body TEXT NOT NULL,
    caption TEXT NOT NULL,
    image_path TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    modified_at DATETIME
);
CREATE TABLE IF NOT EXISTS category (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS post_category_relation (
    category_id INTEGER,
    post_id INTEGER,
    PRIMARY KEY(category_id, post_id),
    FOREIGN KEY(category_id) REFERENCES category (id),
    FOREIGN KEY(post_id) REFERENCES post (id)
);
CREATE TABLE contact (
    id INTEGER,
    subject TEXT NOT NULL,
    email INTEGER NOT NULL,
    message TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);
CREATE TABLE about_me (
    id INTEGER,
    content TEXT NOT NULL,
    PRIMARY KEY(id)
);