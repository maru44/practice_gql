DROP TABLE IF EXISTS blogs;
create table blogs (
    id SERIAL NOT NULL PRIMARY KEY,
    slug VARCHAR(16) NOT NULL,
    title VARCHAR(64) NOT NULL,
    abstract VARCHAR(160) NULL,
    content TEXT NOT NULL,
    user_id VARCHAR(128) NOT NULL,
    is_public BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- index
CREATE INDEX blog_user ON blogs (user_id);
CREATE INDEX blog_user_public ON blogs (user_id, is_public);
CREATE INDEX blog_slug ON blogs (slug);