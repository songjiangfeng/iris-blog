package models

type Post struct {
	ID         int64  `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Views      string `json:"views"`
}
