package models

type Tag struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	PostID     int    `json:"post_id"`
	PostTitle  string `json:"post_title"`
	Content    string `json:"content"`
	Created_at string `json:"created_at"`
}
