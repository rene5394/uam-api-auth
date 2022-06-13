package repositories

import (
	"api_auth/models"
)

func FindUserByEmail(user *models.User, email string) error {
	query := models.DB.Where("email = ?", email).Where("status_id = ?", 1)
	err := query.First(user)

	return err
}
