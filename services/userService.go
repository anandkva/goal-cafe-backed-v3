package services

import (
	"context"
	"goal-cafe-backend/config"
	"goal-cafe-backend/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetUserByEmail(email string) (*models.User, error) {
	collection := config.DB.Collection("users")
	user := &models.User{}
	err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(user)
	return user, err
}
