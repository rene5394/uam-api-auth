package dtos

type Auth struct {
	Email    string `json:"email" db:"email" rw:"r"`
	Password string `json:"password" db:"password" rw:"r"`
}
