package models

// ----------> Purchase Model
type Purchase struct {
	Quantity uint   `json:"quantity"`
	UserID   string `json:"user_id"`
}
