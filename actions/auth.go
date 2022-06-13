package actions

import (
	"api_auth/dtos"
	"api_auth/models"
	"api_auth/utils"
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// AuthenticateHandler is a handler to server the auth route
func AuthenticateHandler(c buffalo.Context) error {
	auth := &dtos.Auth{}
	if err := c.Bind(auth); err != nil {
		return err
	}

	user := models.User{}
	query := models.DB.Where("email = ?", auth.Email).Where("status_id = ?", 1)
	err := query.First(&user)

	if err != nil {
		return c.Render(http.StatusNotFound, r.JSON(map[string]string{"message": "User not found or deactivated!"}))
	}

	matchPassword := utils.CheckPasswordHash(auth.Password, user.Password)

	if !matchPassword {
		return c.Render(http.StatusUnauthorized, r.JSON(map[string]string{"message": "Wrong Password!"}))
	}

	jwt, err := utils.GenerateJWT(user)

	if err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(map[string]string{"error": err.Error()}))
	}

	return c.Render(http.StatusOK, r.JSON(jwt))
}
