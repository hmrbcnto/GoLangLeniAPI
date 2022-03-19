package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

const dbName = "leni-facts-api"
const mongoUri = "mongodb+srv://hmrbcnt:jonasbayot@fullstackopenmongodb.lqee8.mongodb.net/leniApi?retryWrites=true&w=majority"

func NewMongoClient() (*mongo.Client, error) {
	// Connecting to mongodb uri
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))

	// Defining a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err = client.Connect(ctx); err != nil {
		return nil, err
	}

	return client, nil
}
