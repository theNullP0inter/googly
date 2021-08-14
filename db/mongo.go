package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewMongoDatabase creates a new instance of mongo db
func NewMongoDatabase(mongo_uri string, database_name string) *mongo.Database {

	client, err := mongo.NewClient(options.Client().ApplyURI(mongo_uri))

	if err != nil {
		fmt.Printf("MongoDB Client Init Error: %v", err)
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		fmt.Printf("MongoDB Client Failed to Connect: %v", err)
		panic(err)
	}
	return client.Database(database_name)

}
