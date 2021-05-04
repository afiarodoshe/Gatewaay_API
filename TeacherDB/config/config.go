package config

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var client *mongo.Client
var e error

func LoadDB() (*mongo.Client, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI(os.Getenv("DB_URL"))

	// Connect to MongoDB
	client, e = mongo.Connect(context.TODO(), clientOptions)
	CheckError(e)

	// Check the connection
	e = client.Ping(context.TODO(), nil)
	CheckError(e)

	return client, e
}

func CheckError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func LoadEnvironments() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading ..env file")
	}
}

func GetTeachersCollection() *mongo.Collection {
	client, _ := LoadDB()
	// get collection as ref
	collection := client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("DB_COLLECTION"))
	return collection
}
