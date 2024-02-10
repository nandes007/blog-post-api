CREATE TABLE posts
(
    id SERIAL PRIMARY KEY,
    author_id INTEGER,
    title TEXT,
    content TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);