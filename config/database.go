package config

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
	dbHost := os.Getenv("DB_URL")
	clientOptions := options.Client().ApplyURI(dbHost)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	DB = client.Database("goal-cafe")
}
