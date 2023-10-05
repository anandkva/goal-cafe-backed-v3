package utils

import (
	"context"
	"goal-cafe-backend/models"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func GenerateJWTToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"email":  user.Email,
		"role":   user.Role,
	})
	tokenString, err := token.SignedString([]byte("Goal-Cafe"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func UserExists(collection *mongo.Collection, email string) bool {
	existingUser := models.User{}
	err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&existingUser)
	return err == nil
}

func InsertUser(collection *mongo.Collection, user models.User) error {
	_, err := collection.InsertOne(context.TODO(), user)
	return err
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ComparePasswords(providedPassword string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(providedPassword))
	return err == nil
}
