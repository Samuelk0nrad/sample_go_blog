-- +goose Up
CREATE TABLE blog_metadata (
  id SERIAL PRIMARY KEY,
  title TEXT,
  description TEXT,
  author TEXT
);

-- +goose Down
DROP TABLE blog_metadata;
