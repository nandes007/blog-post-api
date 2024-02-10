CREATE TABLE post_comments
(
    id SERIAL PRIMARY KEY,
    post_id INT REFERENCES posts (id) NOT NULL,
    user_id INT REFERENCES users (id) NOT NULL,
    parent_id INT,
    content TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);