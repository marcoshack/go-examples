package main

import (
	"context"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/marcoshack/go-examples/twirp/api"
	"github.com/marcoshack/go-examples/twirp/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/twitchtv/twirp"
)

type contextKey int

const (
	startTimeCtxKey contextKey = 1 + iota
	requestIdCtxKey
)

func main() {
	// initialize zerolog logger
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "15:04:05"}
	logger := zerolog.New(output).With().Timestamp().Logger()

	// initialize twirp server and register service
	service := server.NewPingService()
	twirpHandler := api.NewPingServiceServer(service,
		twirp.WithServerHooks(NewServerHooks(logger, api.PingServicePathPrefix)),
	)

	// TODO read from CLI argument
	bindAddr := ":8081"
	logger.Info().Str("bindAddr", bindAddr).Msg("server listening")
	err := http.ListenAndServe(bindAddr, twirpHandler)
	if err != nil {
		logger.Fatal().Err(err).Msg("error listening")
	}
}

func NewServerHooks(logger zerolog.Logger, pathPrefix string) *twirp.ServerHooks {
	return &twirp.ServerHooks{
		RequestRouted: func(ctx context.Context) (context.Context, error) {
			ctx = context.WithValue(ctx, startTimeCtxKey, time.Now())
			method, _ := twirp.MethodName(ctx)
			blocked, err := isBlocked(pathPrefix, method)

			requestId := getOrCreateRequestId(ctx)
			requestCtx := logger.With().Fields(map[string]interface{}{
				"method":    method,
				"requestId": requestId,
			}).Logger().WithContext(ctx)

			log.Ctx(requestCtx).Info().Bool("blocked", blocked).Msg("REQUEST")

			return requestCtx, err
		},
		ResponseSent: func(ctx context.Context) {
			startTime := time.Now()
			startTimeCtxValue := ctx.Value(startTimeCtxKey)
			if startTimeCtxValue != nil {
				startTime = startTimeCtxValue.(time.Time)
			}
			status, _ := twirp.StatusCode(ctx)
			log.Ctx(ctx).Info().
				Str("elapsed", time.Since(startTime).String()).
				Str("status", status).
				Msg("RESPONSE")
		},
	}
}

func isBlocked(pathPrefix, method string) (bool, error) {
	if strings.ToLower(os.Getenv("STAGE")) != "prod" || !strings.HasPrefix(method, "Unsafe") {
		return false, nil
	}

	// reproduce the same message and meta as a real not found from the generated twirp code.
	path := pathPrefix + method
	err := twirp.NewError(twirp.BadRoute, "no handler for path \""+path+"\"")
	err = err.WithMeta("twirp_invalid_route", "POST "+path)

	return true, err
}

func getOrCreateRequestId(ctx context.Context) string {
	if header, ok := twirp.HTTPRequestHeaders(ctx); ok {
		if headerRequestId, ok := header["X-Request-Id"]; ok {
			return strings.Join(headerRequestId, ":")
		}
	}
	return uuid.New().String()
}
