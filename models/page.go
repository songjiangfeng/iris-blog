package models

type Page struct {
	ID         int64  `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Slug       string `json:"slug"`
}
