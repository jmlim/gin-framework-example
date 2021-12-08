package db

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

// Client Database instance
var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("tutorial2").Collection(collectionName)
	return collection
}

func DBinstance() *mongo.Client {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	MongoDb := os.Getenv("MONGODB_URL")

	//MongoDB 드라이버는 커넥션 풀을 명시적으로 설정하지 않아도 커넥션 풀을 사용하고 기본 값이 MaxPoolSize 100, MinPoolSize 0 이다
	clientOptions := options.Client().ApplyURI(MongoDb)
	/**
	clientOptions.SetMaxPoolSize(100)
	clientOptions.SetMinPoolSize(0)
	*/
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	return client
}
