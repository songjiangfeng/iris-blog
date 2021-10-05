// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package models

import (
	"context"
	"database/sql"
)

const createPost = `-- name: CreatePost :execresult
INSERT INTO iris_posts (
  title, content, views
) VALUES (
  ?, ?, ?
)
`

type CreatePostParams struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Views   int64  `json:"views"`
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createPost, arg.Title, arg.Content, arg.Views)
}

const deletePost = `-- name: DeletePost :exec
DELETE FROM iris_posts
WHERE id = ?
`

func (q *Queries) DeletePost(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deletePost, id)
	return err
}

const getHotPost = `-- name: GetHotPost :many
SELECT id, title, content, created_at, updated_at, views From iris_posts ORDER BY views DESC LIMIT ?
`

func (q *Queries) GetHotPost(ctx context.Context, limit int32) ([]IrisPost, error) {
	rows, err := q.db.QueryContext(ctx, getHotPost, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []IrisPost
	for rows.Next() {
		var i IrisPost
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Views,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getLatestPost = `-- name: GetLatestPost :many
SELECT id, title, content, created_at, updated_at, views From iris_posts ORDER BY created_at DESC LIMIT ?
`

func (q *Queries) GetLatestPost(ctx context.Context, limit int32) ([]IrisPost, error) {
	rows, err := q.db.QueryContext(ctx, getLatestPost, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []IrisPost
	for rows.Next() {
		var i IrisPost
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Views,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMenu = `-- name: GetMenu :many
SELECT id, name, path, weight FROM iris_menus ORDER BY weight, id ASC
`

func (q *Queries) GetMenu(ctx context.Context) ([]IrisMenu, error) {
	rows, err := q.db.QueryContext(ctx, getMenu)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []IrisMenu
	for rows.Next() {
		var i IrisMenu
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Path,
			&i.Weight,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getNextPost = `-- name: GetNextPost :one
SELECT id, title, content, created_at, updated_at, views FROM iris_posts WHERE id > ? ORDER BY id ASC LIMIT 1
`

func (q *Queries) GetNextPost(ctx context.Context, id int64) (IrisPost, error) {
	row := q.db.QueryRowContext(ctx, getNextPost, id)
	var i IrisPost
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Views,
	)
	return i, err
}

const getPage = `-- name: GetPage :one
SELECT id, title, content, created_at, updated_at, slug FROM iris_pages WHERE slug = ? LIMIT  1
`

func (q *Queries) GetPage(ctx context.Context, slug string) (IrisPage, error) {
	row := q.db.QueryRowContext(ctx, getPage, slug)
	var i IrisPage
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Slug,
	)
	return i, err
}

const getPost = `-- name: GetPost :one
SELECT id, title, content, created_at, updated_at, views FROM iris_posts WHERE id = ? LIMIT 1
`

func (q *Queries) GetPost(ctx context.Context, id int64) (IrisPost, error) {
	row := q.db.QueryRowContext(ctx, getPost, id)
	var i IrisPost
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Views,
	)
	return i, err
}

const getPostByPage = `-- name: GetPostByPage :many
SELECT id, title, content, created_at, updated_at, views From iris_posts ORDER BY created_at DESC LIMIT ?,?
`

type GetPostByPageParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetPostByPage(ctx context.Context, arg GetPostByPageParams) ([]IrisPost, error) {
	rows, err := q.db.QueryContext(ctx, getPostByPage, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []IrisPost
	for rows.Next() {
		var i IrisPost
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Views,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPrevPost = `-- name: GetPrevPost :one
SELECT id, title, content, created_at, updated_at, views FROM iris_posts WHERE id < ?  ORDER BY id  DESC LIMIT 1
`

func (q *Queries) GetPrevPost(ctx context.Context, id int64) (IrisPost, error) {
	row := q.db.QueryRowContext(ctx, getPrevPost, id)
	var i IrisPost
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Views,
	)
	return i, err
}

const getSite = `-- name: GetSite :one
SELECT id, site_name, site_email, slogan, notice FROM iris_site WHERE id = ? LIMIT  1
`

func (q *Queries) GetSite(ctx context.Context, id int64) (IrisSite, error) {
	row := q.db.QueryRowContext(ctx, getSite, id)
	var i IrisSite
	err := row.Scan(
		&i.ID,
		&i.SiteName,
		&i.SiteEmail,
		&i.Slogan,
		&i.Notice,
	)
	return i, err
}

const getTag = `-- name: GetTag :one
SELECT id, name FROM iris_tags WHERE id = ? LIMIT 1
`

func (q *Queries) GetTag(ctx context.Context, id int64) (IrisTag, error) {
	row := q.db.QueryRowContext(ctx, getTag, id)
	var i IrisTag
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getTagPostId = `-- name: GetTagPostId :many
SELECT post_id FROM iris_tags_posts WHERE tag_id = ?
`

func (q *Queries) GetTagPostId(ctx context.Context, tagID int64) ([]int64, error) {
	rows, err := q.db.QueryContext(ctx, getTagPostId, tagID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int64
	for rows.Next() {
		var post_id int64
		if err := rows.Scan(&post_id); err != nil {
			return nil, err
		}
		items = append(items, post_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTags = `-- name: GetTags :many
SELECT id, name FROM iris_tags
`

func (q *Queries) GetTags(ctx context.Context) ([]IrisTag, error) {
	rows, err := q.db.QueryContext(ctx, getTags)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []IrisTag
	for rows.Next() {
		var i IrisTag
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUser = `-- name: GetUser :one
SELECT id, username, password, name, avatar, remember_token, created_at, updated_at FROM goadmin_users WHERE id = ? LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int32) (GoadminUser, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i GoadminUser
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Name,
		&i.Avatar,
		&i.RememberToken,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT id, username, password, name, avatar, remember_token, created_at, updated_at FROM goadmin_users
`

func (q *Queries) GetUsers(ctx context.Context) ([]GoadminUser, error) {
	rows, err := q.db.QueryContext(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GoadminUser
	for rows.Next() {
		var i GoadminUser
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Password,
			&i.Name,
			&i.Avatar,
			&i.RememberToken,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listPosts = `-- name: ListPosts :many
SELECT id, title, content, created_at, updated_at, views FROM iris_posts ORDER BY created_at
`

func (q *Queries) ListPosts(ctx context.Context) ([]IrisPost, error) {
	rows, err := q.db.QueryContext(ctx, listPosts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []IrisPost
	for rows.Next() {
		var i IrisPost
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Views,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const viewPlus = `-- name: ViewPlus :exec
UPDATE iris_posts SET views=views+1 WHERE id = ?
`

func (q *Queries) ViewPlus(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, viewPlus, id)
	return err
}