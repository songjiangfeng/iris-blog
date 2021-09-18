package models

type User struct {
	ID             int64  `json:"id"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	Name           string `json:"name"`
	Avatar         string `json:"avatar"`
	Remember_token string `json:"remember_token"`
	Created_at     string `json:"created_at"`
	Updated_at     string `json:"updated_at"`
}
