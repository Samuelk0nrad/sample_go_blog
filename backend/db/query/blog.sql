-- name: GetBlog :one
SELECT * FROM blog_metadata
WHERE id = $1
LIMIT 1;

-- name: ListBlogs :many
SELECT * FROM blog_metadata
order BY title;
