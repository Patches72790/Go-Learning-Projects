package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDB() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	DB_URI := fmt.Sprintf("%s", os.Getenv("MONGO_ATLAS_URI"))

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(DB_URI))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Fatalf("Error pinging db client with err: %s", err)
	}

	fmt.Println("Successfully connected and pinged client")

	return client
}

var DB *mongo.Client = ConnectDB()

func GetDBCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("AlbumsDB").Collection(collectionName)

	return collection
}
