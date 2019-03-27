package server

import (
	"context"
	"log"

	"github.com/afotescu/ciklum_test/PortDomainService/services"

	pb "github.com/afotescu/ciklum_test/protobuf"
)

// Server is an implementation of grpc TransporterServer
type Server struct {
	PortsService services.PortsRetrieverUpdater
}

// CreateOrUpdatePorts will try to add or update a port inside database
func (s *Server) CreateOrUpdatePorts(ctx context.Context, port *pb.Port) (*pb.Response, error) {
	if err := s.PortsService.CreateOrUpdate(ctx, port); err != nil {
		log.Printf("An error ocured during request: %v", err)
		return nil, err
	}
	return &pb.Response{Status: "ok"}, nil
}

// GetPorts will retrieve a list of ports from the database
func (s *Server) GetPorts(ctx context.Context, portsPage *pb.PortsPage) (*pb.Ports, error) {
	result, err := s.PortsService.RetrievePorts(ctx, portsPage)
	if err != nil {
		log.Printf("An error ocured during request: %v", err)
		return nil, err
	}
	return &pb.Ports{Data: result}, nil
}
