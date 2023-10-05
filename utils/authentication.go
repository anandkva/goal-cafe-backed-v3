package utils

import (
	"context"
	"goal-cafe-backend/models"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GenerateJWTToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"email":  user.Email,
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

func IsPasswordValid(inputPassword string, storedPasswordHash string) bool {
	// Implement password hashing and verification logic here.
	// Use a secure password hashing library like bcrypt.
	// For example:
	// return bcrypt.CompareHashAndPassword([]byte(storedPasswordHash), []byte(inputPassword)) == nil
	return inputPassword == storedPasswordHash // For demonstration purposes only, not secure
}
