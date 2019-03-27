package services

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	pb "ciklum_test/protobuf"
)

// PortsRetrieverUpdater wraps RetrievePorts and CreateOrUpdate methods
//
// RetrievePorts returns an array of *pb.Port with portsPage as pagination parameters
// which contains limit and offset
//
// CreateOrUpdate accepts a *pb.Port, inserts it into database and returns a *pb.Response
// with Status: "ok" if everything went ok, otherwise it will return an error
type PortsRetrieverUpdater interface {
	RetrievePorts(ctx context.Context, portsPage *pb.PortsPage) ([]*pb.Port, error)
	CreateOrUpdate(ctx context.Context, port *pb.Port) error
}

// PortsService implements PortsRetrieverUpdater interface
type PortsService struct {
	PortsCollection *mongo.Collection
}

// RetrievePorts returns an array of *pb.Port with portsPage as pagination parameters
// which contains limit and offset
func (p *PortsService) RetrievePorts(ctx context.Context, portsPage *pb.PortsPage) ([]*pb.Port, error) {
	findOptions := options.Find()
	findOptions.SetSkip(int64(portsPage.Offset))
	if portsPage.Limit > 100 {
		portsPage.Limit = 100
	}
	findOptions.SetLimit(int64(portsPage.Limit))

	var result []*pb.Port

	cursor, err := p.PortsCollection.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		err = errors.Wrap(err, "got an error on cursor")
		return nil, err
	}
	for cursor.Next(ctx) {
		var port pb.Port
		err := cursor.Decode(&port)
		if err != nil {
			err = errors.Wrap(err, "could not decode")
			return nil, err
		}
		result = append(result, &port)
	}

	if err := cursor.Err(); err != nil {
		err = errors.Wrap(err, "cursor received an error")
		return nil, err
	}

	if err := cursor.Close(ctx); err != nil {
		err = errors.Wrap(err, "could not close the cursor")
		return nil, err
	}

	return result, nil
}

// CreateOrUpdate accepts a *pb.Port, inserts it into database and returns a *pb.Response
// with Status: "ok" if everything went ok, otherwise it will return an error
func (p *PortsService) CreateOrUpdate(ctx context.Context, port *pb.Port) error {
	upsert := true
	filter := bson.D{{"name", port.Name}}
	update := bson.D{{"$set", port}}
	_, err := p.PortsCollection.UpdateOne(ctx, filter, update, &options.UpdateOptions{Upsert: &upsert})
	if err != nil {
		err = errors.Wrap(err, "could not update or insert")
		return err
	}
	return nil
}
