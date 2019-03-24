package client

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc"

	pb "ciklum_test/protobuf"
)

const (
	address = "port_domain_service:5001"
)

// NewGRPCClient will initialize and return a new Transporter client
func NewGRPCClient(ctx context.Context) (pb.TransporterClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, errors.Wrap(err, "could not establish connection wit grpc server")
	}
	return pb.NewTransporterClient(conn), nil
}
