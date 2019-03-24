package db

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongoURI       = "mongodb://localhost:27017"
	databaseName   = "portsDB"
	collectionName = "ports"
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
