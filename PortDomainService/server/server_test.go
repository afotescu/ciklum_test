package server

import (
	"context"
	"testing"

	"github.com/pkg/errors"

	pb "github.com/afotescu/ciklum_test/protobuf"
)

func TestServer_CreateOrUpdatePorts(t *testing.T) {
	ctx := context.TODO()
	t.Run("should return Response Status: \"ok\"", func(t *testing.T) {
		s := &Server{PortsService: &portServiceMock{}}
		response, err := s.CreateOrUpdatePorts(ctx, &pb.Port{})
		if err != nil {
			t.Fatalf("err should be nil, got: %v", err)
		}
		if response.Status != "ok" {
			t.Fatalf("response.status shoudl be ok, got: %s", response.Status)
		}
	})
	t.Run("resposne should be nil when an error occurred", func(t *testing.T) {
		s := &Server{PortsService: &portServiceMock{errorToReturn: errors.New("error")}}
		response, err := s.CreateOrUpdatePorts(ctx, &pb.Port{})
		if err == nil {
			t.Fatal("should not be nil")
		}
		if response != nil {
			t.Fatal("response should be nil")
		}
	})
}

func TestServer_GetPorts(t *testing.T) {
	ctx := context.TODO()
	t.Run("should return a list of ports", func(t *testing.T) {
		s := &Server{PortsService: &portServiceMock{portsToReturn: []*pb.Port{}}}
		ports, err := s.GetPorts(ctx, &pb.PortsPage{Limit: 10, Offset: 10})
		if err != nil {
			t.Fatalf("err should be nil, got: %v", err)
		}
		if ports == nil {
			t.Fatal("ports should be an empty slice, got nil")
		}

	})

	t.Run("should return nil ports if an error occurred", func(t *testing.T) {
		s := &Server{PortsService: &portServiceMock{portsToReturn: []*pb.Port{}, errorToReturn: errors.New("error")}}
		ports, err := s.GetPorts(ctx, &pb.PortsPage{Limit: 10, Offset: 10})
		if err == nil {
			t.Fatal("err should not be nil")
		}
		if ports != nil {
			t.Fatalf("ports should be nil, got: %v", ports)
		}
	})
}

type portServiceMock struct {
	portsToReturn []*pb.Port
	errorToReturn error
}

func (p *portServiceMock) RetrievePorts(ctx context.Context, portsPage *pb.PortsPage) ([]*pb.Port, error) {
	if p.errorToReturn != nil {
		return nil, p.errorToReturn
	}
	return p.portsToReturn, p.errorToReturn
}

func (p *portServiceMock) CreateOrUpdate(ctx context.Context, port *pb.Port) error {
	return p.errorToReturn
}
