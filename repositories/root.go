package repositories

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoInit() mongo.Client {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading.env file")
	}

	mongoDBURI := os.Getenv("Mongo_URL")

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoDBURI))
	if err != nil {
		log.Fatal(err)
	}

	// 실제로 MongoDB에 연결하는 부분
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to mongoDB")
	return *client
}