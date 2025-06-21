-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

--  SELECT uuid_generate_v4();

CREATE TABLE IF NOT EXISTS users (
  session_id uuid UNIQUE,
  name VARCHAR(150)  NOT NULL,
  avatar_url TEXT --will be null if it is in ricky and morty
);

CREATE TABLE IF NOT EXISTS posts (
  post_id SERIAL PRIMARY KEY,
  user_id uuid REFERENCES users(session_id),
  title VARCHAR(150)  NOT NULL,
  post_content TEXT,
  -- has_image BOOLEAN,
  post_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE IF NOT EXISTS comments (
  comment_id SERIAL PRIMARY KEY,
  post_id INT REFERENCES posts(post_id) ON DELETE CASCADE,
  user_id uuid REFERENCES users(session_id) ON DELETE CASCADE,
  comment_content TEXT,
  parent_comment_id INT REFERENCES comments(comment_id) ON DELETE CASCADE,--
  -- has_image BOOLEAN,
  comment_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- CREATE TABLE IF NOT EXISTS subcomments (
--   comment_parent INT REFERENCES comments(comment_id) ON DELETE CASCADE,
--   comment_child INT REFERENCES comments(comment_id) ON DELETE CASCADE
-- );

/*
SELECT 
    c.comment_id,
    c.comment_content,
    c.parent_comment_id,
    (
        SELECT COUNT(*) 
        FROM comments sub 
        WHERE sub.parent_comment_id = c.comment_id
    ) AS reply_count
FROM comments c
WHERE c.post_id = $1
ORDER BY c.comment_time;
*/
