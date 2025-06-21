-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

--  SELECT uuid_generate_v4();

CREATE TABLE IF NOT EXISTS users (
  session_id UUID PRIMARY KEY /*UNIQUE*/,
  name VARCHAR(150)  NOT NULL,
  avatar_url TEXT --will be null if it is in ricky and morty
);

CREATE TABLE IF NOT EXISTS posts (
  post_id GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  user_id UUID REFERENCES users(session_id) NOT NULL,
  title VARCHAR(150)  NOT NULL,
  post_content TEXT,
  has_image BOOLEAN,
  post_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);


CREATE TABLE IF NOT EXISTS comments (
  comment_id GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  post_id INT REFERENCES posts(post_id) ON DELETE CASCADE NOT NULL,
  user_id uuid REFERENCES users(session_id) ON DELETE CASCADE NOT NULL,
  comment_content TEXT,
  parent_comment_id INT REFERENCES comments(comment_id) ON DELETE CASCADE,--DEFAULT NULL
  has_image BOOLEAN,
  comment_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
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
