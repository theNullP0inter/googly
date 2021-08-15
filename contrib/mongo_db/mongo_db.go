package mongo_db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewMongoDatabase creates a new instance of mongo db
func NewMongoDatabase(mongoUri string, databaseName string) *mongo.Database {

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))

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
	return client.Database(databaseName)

}
