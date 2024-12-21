package main

import (
	"context"
	"fmt"
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/marcoshack/go-examples/grpc/api/api"
	"github.com/marcoshack/go-examples/grpc/api/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	logger := zerolog.Nop()

	// initialize grpc server and register service
	grpcOptions := createServerOptions(logger)
	grpcServer := grpc.NewServer(grpcOptions...)
	api.RegisterPingServiceServer(grpcServer, &server.PingService{})

	// start server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Err(err).Msg("failed to create server listener")
	}
	log.Info().Str("bindAddr", lis.Addr().String()).Msg("server listening")
	if err := grpcServer.Serve(lis); err != nil {
		log.Err(err).Msg("failed to create server listener")
	}
}

func createServerOptions(logger zerolog.Logger) []grpc.ServerOption {
	return []grpc.ServerOption{
		grpc.UnaryInterceptor(logging.UnaryServerInterceptor(interceptorLogger(logger))),
		grpc.StreamInterceptor(logging.StreamServerInterceptor(interceptorLogger(logger))),
	}
}

// interceptorLogger adapts zerolog logger to interceptor logger.
func interceptorLogger(l zerolog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		l := l.With().Fields(fields).Logger()

		switch lvl {
		case logging.LevelDebug:
			l.Debug().Msg(msg)
		case logging.LevelInfo:
			l.Info().Msg(msg)
		case logging.LevelWarn:
			l.Warn().Msg(msg)
		case logging.LevelError:
			l.Error().Msg(msg)
		default:
			panic(fmt.Sprintf("unknown level %v", lvl))
		}
	})
}
