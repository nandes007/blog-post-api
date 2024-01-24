CREATE TABLE posts
(
    id SERIAL PRIMARY KEY,
    author_id INTEGER,
    title TEXT,
    content TEXT,
    created_at date,
    updated_at date,
    deleted_at date
);