package server

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"

	"github.com/afotescu/ciklum_test/PortDomainService/services"
	pb "github.com/afotescu/ciklum_test/protobuf"
)

// NewGRPCServer creates, register and return a new GRPC server
func NewGRPCServer(ctx context.Context, collection *mongo.Collection) (*grpc.Server, error) {

	grpcServer := grpc.NewServer()

	localServer := &Server{PortsService: &services.PortsService{PortsCollection: collection}}

	pb.RegisterTransporterServer(grpcServer, localServer)

	return grpcServer, nil
}
