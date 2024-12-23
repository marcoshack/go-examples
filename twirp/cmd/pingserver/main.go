package main

import (
	"context"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/marcoshack/go-examples/twirp/api/api"
	"github.com/marcoshack/go-examples/twirp/api/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/twitchtv/twirp"
)

type contextKey string

const (
	startTimeCtxKey contextKey = "startTime"
)

func main() {
	// initialize zerolog logger
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "15:04:05"}
	logger := zerolog.New(output).With().Timestamp().Logger()

	// initialize twirp server and register service
	service := server.NewPingService()
	twirpHandler := api.NewPingServiceServer(service,
		twirp.WithServerHooks(NewServerHooks(logger)),
	)

	// TODO read from CLI argument
	bindAddr := ":8080"
	logger.Info().Str("bindAddr", bindAddr).Msg("server listening")
	http.ListenAndServe(bindAddr, twirpHandler)
}

func NewServerHooks(logger zerolog.Logger) *twirp.ServerHooks {
	return &twirp.ServerHooks{
		RequestRouted: func(ctx context.Context) (context.Context, error) {
			ctx = context.WithValue(ctx, startTimeCtxKey, time.Now())
			method, _ := twirp.MethodName(ctx)
			requestId := getOrCreateRequestId(ctx)
			requestCtx := logger.With().Fields(map[string]interface{}{
				"method":    method,
				"requestId": requestId,
			}).Logger().WithContext(ctx)
			log.Ctx(requestCtx).Info().Msg("REQUEST")
			return requestCtx, nil
		},
		ResponseSent: func(ctx context.Context) {
			startTime := ctx.Value(startTimeCtxKey).(time.Time)
			log.Ctx(ctx).Info().Str("elapsed", time.Since(startTime).String()).Msg("RESPONSE")
		},
	}
}

func getOrCreateRequestId(ctx context.Context) string {
	if header, ok := twirp.HTTPRequestHeaders(ctx); ok {
		if headerRequestId, ok := header["X-Request-Id"]; ok {
			return strings.Join(headerRequestId, ":")
		}
	}
	return uuid.New().String()
}
