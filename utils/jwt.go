package utils

import (
	"api_auth/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gobuffalo/envy"
)

func GenerateJWT(user models.User) (string, error) {
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
