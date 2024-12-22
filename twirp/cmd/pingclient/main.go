package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/marcoshack/go-examples/twirp/api/api"
	"github.com/rs/zerolog"
)

func main() {

	var clientType, baseURL string
	var numRequests int
	flag.StringVar(&baseURL, "s", "http://localhost:8080", "Server endpoint")
	flag.StringVar(&clientType, "t", "protobuf", "Type of client to use (json or protobuf)")
	flag.IntVar(&numRequests, "n", 1, "Number of requests to make")
	flag.Parse()

	// initialize zerolog logger
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "15:04:05"}
	logger := zerolog.New(output).With().Timestamp().Logger()

	// initialize twirp client
	httpClient := &http.Client{}
	var twirpClient api.PingService
	switch clientType {
	case "protobuf":
		twirpClient = api.NewPingServiceProtobufClient(baseURL, httpClient)
	case "json":
		twirpClient = api.NewPingServiceJSONClient(baseURL, httpClient)
	}
	logger.Info().Str("client", clientType).Str("baseURL", baseURL).Msg("twirp client initialized")

	// contact the server and print out its response.
	avgLatency := time.Duration(0)
	for i := 0; i < numRequests; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		start := time.Now()
		response, err := twirpClient.Ping(ctx, &api.PingRequest{Message: "Hello, Twirp!"})
		latency := time.Since(start)
		avgLatency += latency
		if err != nil {
			logger.Error().Err(err).Msg("could not ping")
		}
		logger.Info().Str("response", response.GetMessage()).Str("latency", latency.String()).Msg("PING response received")
	}
	logger.Info().Int("numRequests", numRequests).Str("avgLatency", fmt.Sprintf("%dµs", avgLatency.Microseconds()/int64(numRequests))).Msg("done")
}
