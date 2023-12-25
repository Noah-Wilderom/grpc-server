package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var mongoUrl = "mongodb://database:27017"
var mongoTimeout = 15 * time.Second

func ConnectToDatabase() (context.Context, *mongo.Client) {
	clientOptions := options.Client().ApplyURI(mongoUrl)
	clientOptions.SetAuth(options.Credential{
		Username: os.Getenv("MONGO_USERNAME"),
		Password: os.Getenv("MONGO_PASSWORD"),
	})
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	log.Println("Connected to Mongo database")

	ctx, cancel := context.WithTimeout(context.Background(), mongoTimeout)
	defer cancel()

	return ctx, client
}
