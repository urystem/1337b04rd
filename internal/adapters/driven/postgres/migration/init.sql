-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

--  SELECT uuid_generate_v4();

CREATE TABLE IF NOT EXISTS users (
  session_id uuid UNIQUE,
  name VARCHAR(150)  NOT NULL,
  avatar_url TEXT
);

CREATE TABLE IF NOT EXISTS posts (
  post_id SERIAL PRIMARY KEY,
  user_id uuid REFERENCES users(session_id),
  title VARCHAR(150)  NOT NULL,
  post_content TEXT,
  post_image TEXT,
  post_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE IF NOT EXISTS comments (
  comment_id SERIAL PRIMARY KEY,
  post_id INT REFERENCES posts(post_id) ON DELETE CASCADE,
  user_id uuid REFERENCES users(session_id),
  comment_content TEXT,
  comment_image TEXT,
  comment_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS subcomments (
  comment_parent INT REFERENCES comments(comment_id) ON DELETE CASCADE,
  comment_child INT REFERENCES comments(comment_id) ON DELETE CASCADE
);

