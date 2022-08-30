package models

// ----------> Ticket Model
type Ticket struct {
	ID          uint64 `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Description string `json:"desc"`
	Allocation  uint   `json:"allocation"`
}
