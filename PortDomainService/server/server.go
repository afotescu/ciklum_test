package server

import (
	"context"
	"log"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	pb "ciklum_test/protobuf"
)

type MongoCollection interface {
	UpdateOne(ctx context.Context, filter interface{}, update interface{},
		opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	Find(ctx context.Context, filter interface{},
		opts ...*options.FindOptions) (*mongo.Cursor, error)
}

type Server struct {
	Collection MongoCollection
}

func (s *Server) CreateOrUpdatePorts(ctx context.Context, port *pb.Port) (*pb.Response, error) {
	upsert := true
	filter := bson.D{{"name", port.Name}}
	update := bson.D{{"$set", port}}
	ports := s.Collection
	_, err := ports.UpdateOne(ctx, filter, update, &options.UpdateOptions{Upsert: &upsert})
	if err != nil {
		err = errors.Wrap(err, "could not update or insert")
		log.Printf("An error ocured during request: %v", err)
		return nil, err
	}
	return &pb.Response{Status: "ok"}, nil
}

func (s *Server) GetPorts(ctx context.Context, portsPage *pb.PortsPage) (*pb.Ports, error) {
	findOptions := options.Find()
	findOptions.SetSkip(int64(portsPage.Offset))
	if portsPage.Limit > 100 {
		portsPage.Limit = 100
	}
	findOptions.SetLimit(int64(portsPage.Limit))

	var result []*pb.Port
	ports := s.Collection

	cursor, err := ports.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		err = errors.Wrap(err, "got an error on cursor")
		log.Printf("An error ocured during request: %v", err)
		return nil, err
	}
	for cursor.Next(ctx) {
		var port pb.Port
		err := cursor.Decode(&port)
		if err != nil {
			err = errors.Wrap(err, "could not decode")
			log.Printf("An error ocured during request: %v", err)
			return nil, err
		}
		result = append(result, &port)
	}

	if err := cursor.Err(); err != nil {
		err = errors.Wrap(err, "cursor received an error")
		log.Printf("An error ocured during request: %v", err)
		return nil, err
	}

	if err := cursor.Close(ctx); err != nil {
		err = errors.Wrap(err, "could not close the cursor")
		log.Printf("An error ocured during request: %v", err)
		return nil, err
	}
	return &pb.Ports{Data: result}, nil
}
