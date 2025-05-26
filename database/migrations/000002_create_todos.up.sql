CREATE TABLE todos (
   id TEXT PRIMARY KEY,
   title TEXT NOT NULL,
   completed BOOLEAN NOT NULL DEFAULT false,
   category_id TEXT,
   created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
   FOREIGN KEY (category_id) REFERENCES categories(id)
);