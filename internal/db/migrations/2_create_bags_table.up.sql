CREATE TABLE IF NOT EXISTS bags (
                      id SERIAL PRIMARY KEY,
                      title TEXT NOT NULL,
                      category TEXT,
                      color TEXT,
                      price NUMERIC(10, 2),
                      size TEXT,
                      user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                      created_at TIMESTAMP DEFAULT NOW(),
                      updated_at TIMESTAMP DEFAULT NOW(),
                      deleted_at TIMESTAMP DEFAULT NOW()
);
