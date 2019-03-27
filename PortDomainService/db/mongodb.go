package db

import (
	"context"
	"os"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoURI       = os.Getenv("MONGO_URI")
	databaseName   = os.Getenv("DB")
	collectionName = os.Getenv("COLLECTION")
)

// NewMongoDB creates and returns a mongodb database and mongo client
func NewMongoDB(ctx context.Context) (*mongo.Collection, *mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to connect to mongo")
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to ping mongo")
	}
	return client.Database(databaseName).Collection(collectionName), client, nil
}
