package server

import (
	"context"

	"github.com/marcoshack/go-examples/twirp/api/api"
)

type PingService struct {
}

func NewPingService() *PingService {
	return &PingService{}
}

func (s *PingService) Ping(ctx context.Context, req *api.PingRequest) (*api.PingResponse, error) {
	return &api.PingResponse{Message: "Pong"}, nil
}
