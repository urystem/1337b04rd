-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

--  SELECT uuid_generate_v4();

CREATE TABLE IF NOT EXISTS users (
  session_id UUID PRIMARY KEY /*UNIQUE*/,
  name VARCHAR(150),
  avatar_url TEXT--will be null if it is in ricky and morty
);
-------------------------------------------

CREATE TABLE IF NOT EXISTS posts (
  post_id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  user_id UUID REFERENCES users(session_id) ON DELETE CASCADE NOT NULL,
  author VARCHAR(150)  NOT NULL,
  title VARCHAR(150)  NOT NULL,
  post_content TEXT,
  has_image BOOLEAN NOT NULL,
  archived BOOLEAN DEFAULT FALSE NOT NULL,
  post_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- Быстрый доступ к только активным или архивным
CREATE INDEX idx_posts_archived ON posts(archived);

-- -- Комбинированный для главной (если часто WHERE archived=false ORDER BY post_time DESC)
-- CREATE INDEX idx_posts_main_page ON posts(archived, post_time DESC);

------------------------------------------
-- ) PARTITION BY LIST (archived);

-- -- Активные посты (на главной)
-- CREATE TABLE IF NOT EXISTS posts_active
--   PARTITION OF posts FOR VALUES IN (false);

-- -- Архивированные посты
-- CREATE TABLE IF NOT EXISTS posts_archived
--   PARTITION OF posts FOR VALUES IN (true);


-------------------------------------------------------

CREATE TABLE IF NOT EXISTS comments (
  comment_id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  post_id INT REFERENCES posts(post_id) ON DELETE CASCADE NOT NULL,
  user_id UUID REFERENCES users(session_id) ON DELETE CASCADE NOT NULL,
  comment_content TEXT,
  parent_comment_id BIGINT REFERENCES comments(comment_id) ON DELETE CASCADE,--DEFAULT NULL
  has_image BOOLEAN,
  comment_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

INSERT INTO users (session_id, name, avatar_url) VALUES
  ('11111111-1111-1111-1111-111111111111', 'Rick Sanchez', 'https://rickandmortyapi.com/api/character/avatar/183.jpeg'),
  ('22222222-2222-2222-2222-222222222222', 'Morty Smith', 'https://rickandmortyapi.com/api/character/avatar/2.jpeg'),
  ('33333333-3333-3333-3333-333333333333', 'Summer Smith', 'https://rickandmortyapi.com/api/character/avatar/1.jpeg');

INSERT INTO posts (user_id, author,title, post_content, has_image) VALUES
  ('11111111-1111-1111-1111-111111111111','doldan', 'Wubba Lubba Dub Dub', 'Rick''s famous quote', FALSE),
  ('22222222-2222-2222-2222-222222222222', 'shakrt','First Adventure', 'That was crazy...', TRUE);

-- Комментарии без родителя
INSERT INTO comments (post_id, user_id, comment_content, has_image) VALUES
  (1, '33333333-3333-3333-3333-333333333333', 'Awesome!', FALSE),
  (2, '11111111-1111-1111-1111-111111111111', 'Morty, be careful...', FALSE);

-- Ответ (reply) на первый комментарий
INSERT INTO comments (post_id, user_id, comment_content, parent_comment_id, has_image) VALUES
  (1, '22222222-2222-2222-2222-222222222222', 'I know, right?', 1, FALSE);


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
