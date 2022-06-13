package repositories

import "api_auth/models"

func FindEmployeeByUserID(employee *models.Employee, userID int) error {
	query := models.DB.Where("user_id = ?", userID)
	err := query.First(employee)

	return err
}
