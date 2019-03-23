package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "ciklum_test/protobuf"
	"google.golang.org/grpc"
)

type Server struct{}

func (s *Server) CreateOrUpdatePorts(ctx context.Context, port *pb.Port) (*pb.Response, error) {
	fmt.Println(port)
	return &pb.Response{}, nil
}

func (s *Server) GetPorts(ctx context.Context, ports *pb.PortsPage) (*pb.Ports, error) {
	return &pb.Ports{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	server := &Server{}
	pb.RegisterTransporterServer(s, server)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
