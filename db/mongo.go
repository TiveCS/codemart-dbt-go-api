package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	once   sync.Once
)

func getMongoURI() string {
	env := godotenv.Load()
	if env != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("MONGODB_URI")
}

func ConnectMongo() {
	uri := getMongoURI()
	options := options.Client().ApplyURI(uri)

	newClient, err := mongo.NewClient(options)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = newClient.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = newClient.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	client = newClient
}

func GetMongoClient() *mongo.Client {
	once.Do(func() {
		ConnectMongo()
	})

	return client
}
