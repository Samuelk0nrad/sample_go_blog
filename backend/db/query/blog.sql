-- name: GetBlog :one
SELECT * FROM blog_post
WHERE id = $1
LIMIT 1;

-- name: ListBlogs :many
SELECT * FROM blog_post
order BY title;
