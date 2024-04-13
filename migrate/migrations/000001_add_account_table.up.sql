CREATE TABLE IF NOT EXISTS accounts (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL,
    password_hash TEXT,
    created_time TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_time TIMESTAMPTZ DEFAULT NULL
);