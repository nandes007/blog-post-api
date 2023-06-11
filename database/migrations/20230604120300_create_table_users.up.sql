CREATE TABLE users
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(225) NOT NULL,
    created_at date,
    updated_at date,
    deleted_at date
)