package utils

import (
	"goal-cafe-backend/models"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWTToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"email":  user.Email,
	})

	// Sign the token with a secret key
	tokenString, err := token.SignedString([]byte("Goal-Cafe"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
