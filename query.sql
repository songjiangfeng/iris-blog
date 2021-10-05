-- name: GetPost :one
SELECT * FROM iris_posts WHERE id = ? LIMIT 1;

-- name: GetPrevPost :one
SELECT * FROM iris_posts WHERE id < ?  ORDER BY id  DESC LIMIT 1;

-- name: GetNextPost :one
SELECT * FROM iris_posts WHERE id > ? ORDER BY id ASC LIMIT 1 ;

-- name: ListPosts :many
SELECT * FROM iris_posts ORDER BY created_at;

-- name: CreatePost :execresult
INSERT INTO iris_posts (
  title, content, views
) VALUES (
  ?, ?, ?
);

-- name: DeletePost :exec
DELETE FROM iris_posts
WHERE id = ?;

-- name: ViewPlus :exec
UPDATE iris_posts SET views=views+1 WHERE id = ?;

-- name: GetPostByPage :many
SELECT * From iris_posts ORDER BY created_at DESC LIMIT ?,?;

-- name: GetLatestPost :many
SELECT * From iris_posts ORDER BY created_at DESC LIMIT ?;

-- name: GetHotPost :many
SELECT * From iris_posts ORDER BY views DESC LIMIT ?;

-- name: GetSite :one
SELECT * FROM iris_site WHERE id = ? LIMIT  1;

-- name: GetMenu :many
SELECT * FROM iris_menus ORDER BY weight, id ASC;

-- name: GetPage :one
SELECT * FROM iris_pages WHERE slug = ? LIMIT  1;


-- name: GetUser :one
SELECT * FROM goadmin_users WHERE id = ? LIMIT 1;

-- name: GetUsers :many
SELECT * FROM goadmin_users;

-- name: GetTag :one
SELECT * FROM iris_tags WHERE id = ? LIMIT 1;

-- name: GetTags :many
SELECT * FROM iris_tags;

-- name: GetTagPostId :many
SELECT post_id FROM iris_tags_posts WHERE tag_id = ?;

