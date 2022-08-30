package models

// ----------> User Model
type User struct {
	ID       string `json:"id"`
	Username string `json:"user_name"`
	Password string `json:"pass"`
}
