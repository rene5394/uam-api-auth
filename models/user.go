package models

type User struct {
	ID       int    `json:"id" db:"id" rw:"r"`
	Email    string `json:"email" db:"email" rw:"r"`
	Password string `json:"password" db:"password" rw:"r"`
}
