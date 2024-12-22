package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/marcoshack/go-examples/twirp/api/api"
	"github.com/rs/zerolog"
)

func main() {
	// initialize zerolog logger
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "15:04:05"}
	logger := zerolog.New(output).With().Timestamp().Logger()

	// initialize twirp client
	httpClient := &http.Client{}
	baseURL := "http://localhost:8080"
	twirpClient := api.NewPingServiceProtobufClient(baseURL, httpClient)

	// contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	start := time.Now()
	response, err := twirpClient.Ping(ctx, &api.PingRequest{Message: "Hello, Twirp!"})
	elapsed := time.Since(start)
	if err != nil {
		logger.Error().Err(err).Msg("could not ping")
	}
	logger.Info().Str("response", response.GetMessage()).Str("elapsed", elapsed.String()).Msg("PING response received")
}
