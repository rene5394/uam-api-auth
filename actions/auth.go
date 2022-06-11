package actions

import (
	"api_auth/dtos"
	"api_auth/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	"golang.org/x/crypto/bcrypt"
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

	matchPassword := checkPasswordHash(auth.Password, user.Password)

	if !matchPassword {
		return c.Render(http.StatusUnauthorized, r.JSON(map[string]string{"message": "Wrong Password!"}))
	}

	jwt, err := generateJWT(user)

	if err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(map[string]string{"error": err.Error()}))
	}

	return c.Render(http.StatusOK, r.JSON(jwt))
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

func generateJWT(user models.User) (string, error) {
	jwtKey, err := envy.MustGet("JWT_KEY")

	if jwtKey == "" {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["email"] = user.Email
	claims["expiresAt"] = time.Now().Add(time.Hour * 10).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(jwtKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
