package server

import (
	"context"

	"github.com/marcoshack/go-examples/grpc/api/api"
)

type PingService struct {
	api.UnimplementedPingServiceServer
}

func (s *PingService) Ping(ctx context.Context, in *api.PingRequest) (*api.PingResponse, error) {
	return &api.PingResponse{Message: "Pong"}, nil
}
