CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username TEXT UNIQUE NOT NULL,
                       name TEXT NOT NULL DEFAULT '',
                       password TEXT NOT NULL,
                       role TEXT NOT NULL DEFAULT 'user',
                       created_at TIMESTAMP DEFAULT NOW()
);