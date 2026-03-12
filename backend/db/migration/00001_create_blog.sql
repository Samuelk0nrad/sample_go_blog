-- +goose Up
CREATE TABLE blog_post (
  id          SERIAL PRIMARY KEY,
  title       TEXT not null,
  slug        TEXT unique not null,
  description TEXT,
  author      TEXT not null,

  created_by TEXT not null,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP not null,

  updated_by TEXT,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tag (
  id          SERIAL PRIMARY KEY,
  title       TEXT not null,
  slug        TEXT unique not null,
  description TEXT,

  created_by TEXT not null,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP not null,

  updated_by TEXT,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE blog_tag (
  tag_id  INTEGER not null,
  blog_id INTEGER not null,

  CONSTRAINT fk_tag FOREIGN KEY(tag_id) REFERENCES tag(id) ON DELETE CASCADE,
  CONSTRAINT fk_blog FOREIGN KEY(blog_id) REFERENCES blog_post(id) ON DELETE CASCADE,

  PRIMARY KEY (tag_id, blog_id)
);

-- +goose Down
DROP TABLE blog_tag;
DROP TABLE tag;
DROP TABLE blog_post;
