package server

import (
	"context"
	"testing"

	pb "ciklum_test/protobuf"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestServer_CreateOrUpdatePorts(t *testing.T) {
	ctx := context.TODO()
	s := &Server{Collection: &collectionMock{}}
	_, err := s.CreateOrUpdatePorts(ctx, &pb.Port{})
	if err != nil {
		t.Fatalf("err should be nil, got: %v", err)
	}
}

func TestServer_GetPorts(t *testing.T) {
	ctx := context.TODO()
	s := &Server{Collection: &collectionMock{}}
	_, err := s.GetPorts(ctx, &pb.PortsPage{Limit: 10, Offset: 10})
	if err == nil {
		t.Fatalf("err should be %v, got: %v", mongo.ErrNilCursor, err)
	}
}

type collectionMock struct{}

func (*collectionMock) UpdateOne(context.Context, interface{}, interface{}, ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return &mongo.UpdateResult{}, nil
}

func (*collectionMock) Find(context.Context, interface{}, ...*options.FindOptions) (*mongo.Cursor, error) {
	return nil, mongo.ErrNilCursor
}
