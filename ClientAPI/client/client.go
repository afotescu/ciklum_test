package client

import (
	"context"
	"os"

	"github.com/pkg/errors"
	"google.golang.org/grpc"

	pb "ciklum_test/protobuf"
)

var portDomainServiceURI = os.Getenv("PORT_DOMAIN_SERVICE_URI")

// NewGRPCClient will initialize and return a new Transporter client
func NewGRPCClient(ctx context.Context) (pb.TransporterClient, error) {
	conn, err := grpc.Dial(portDomainServiceURI, grpc.WithInsecure())
	if err != nil {
		return nil, errors.Wrap(err, "could not establish connection wit grpc server")
	}
	return pb.NewTransporterClient(conn), nil
}
