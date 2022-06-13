package models

type Employee struct {
	ID     int `json:"id" db:"id" rw:"r"`
	UserID int `json:"user_id" db:"user_id" rw:"r"`
}
