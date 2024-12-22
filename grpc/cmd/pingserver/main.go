package main

import (
	"context"
	"fmt"
	"net"
	"os"

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
	// initialize zerolog logger
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "15:04:05"}
	logger := zerolog.New(output).With().Timestamp().Logger()

	// initialize grpc server and register service
	grpcOptions := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(grpcOptions...)
	api.RegisterPingServiceServer(grpcServer, server.NewPingService())

	// start server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		logger.Err(err).Msg("failed to create server listener")
	}
	logger.Info().Str("bindAddr", lis.Addr().String()).Msg("server listening")
	if err := grpcServer.Serve(lis); err != nil {
		log.Err(err).Msg("failed to create server listener")
	}
}

func CreateServerOptions(logger zerolog.Logger) []grpc.ServerOption {
	opts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
		// Add any other option (check functions starting with logging.With).
	}
	return []grpc.ServerOption{
		grpc.UnaryInterceptor(logging.UnaryServerInterceptor(interceptorLogger(logger), opts...)),
		grpc.StreamInterceptor(logging.StreamServerInterceptor(interceptorLogger(logger), opts...)),
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
