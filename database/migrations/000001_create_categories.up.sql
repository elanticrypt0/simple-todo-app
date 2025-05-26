CREATE TABLE categories (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO categories (id, name) VALUES
  ('cat1', 'Trabajo'),
  ('cat2', 'Personal');