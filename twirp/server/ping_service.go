package server

import (
	"context"

	"github.com/marcoshack/go-examples/twirp/api"
	"github.com/rs/zerolog/log"
)

type PingService struct {
}

func NewPingService() *PingService {
	return &PingService{}
}

func (s *PingService) Ping(ctx context.Context, req *api.PingRequest) (*api.PingResponse, error) {
	log.Ctx(ctx).Debug().Interface("request", req).Msg("Processing Ping request")
	return &api.PingResponse{Message: "Pong"}, nil
}

func (s *PingService) UnsafePing(ctx context.Context, req *api.PingRequest) (*api.PingResponse, error) {
	log.Ctx(ctx).Debug().Interface("request", req).Msg("Processing UnsafePing request")
	return &api.PingResponse{Message: "UnsafePong"}, nil
}
