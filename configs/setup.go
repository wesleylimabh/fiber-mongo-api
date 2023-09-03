package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(EnvMongoURI()))

	if err != nil {
		log.Fatal(err)
	}

	//ping database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")
	return client

}

// client instance
var db *mongo.Client = ConnectDB()

func GetCollection(collectionName string) *mongo.Collection {
	collection := db.Database("golangAPI").Collection(collectionName)
	return collection
}
