package models

// ----------> User Model
type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"pass"`
}
