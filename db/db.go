package db

import (
	"context"
	"log"

	middlewares "github.com/chattertechno/challenge-platform-api/handlers"
	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// Dbconnect -> connects mongo
func Dbconnect() *mongo.Client {

	clientOptions := options.Client().ApplyURI(middlewares.DotEnvVariable("MONGO_URL"))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("⛒ Connection Failed to Database 1")
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("⛒ Connection Failed to Database 2")
		log.Fatal(err)
	}
	color.Green("⛁ Connected to Database")
	return client
}
