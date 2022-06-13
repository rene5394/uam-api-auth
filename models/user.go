package models

type User struct {
	ID       int    `json:"id" db:"id" rw:"r"`
	Email    string `json:"email" db:"email" rw:"r"`
	Password string `json:"password" db:"password" rw:"r"`
	RoleID   string `json:"role_id" db:"role_id" rw:"r"`
}
