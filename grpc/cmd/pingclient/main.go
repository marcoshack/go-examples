package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/marcoshack/go-examples/grpc/api/api"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var baseURL string
	var numRequests int
	flag.StringVar(&baseURL, "s", "localhost:50051", "Server endpoint")
	flag.IntVar(&numRequests, "n", 5, "Number of requests to make")
	flag.Parse()

	// initialize zerolog logger
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "15:04:05"}
	logger := zerolog.New(output).With().Timestamp().Logger()

	// Set up a connection to the server.
	conn, err := grpc.NewClient(baseURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal().Err(err).Msg("did not connect")
	}
	defer conn.Close()
	c := api.NewPingServiceClient(conn)

	// Contact the server and print out its response.
	avgLatency := time.Duration(0)
	for i := 0; i < numRequests; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		start := time.Now()
		r, err := c.Ping(ctx, &api.PingRequest{Message: "ping"})
		latency := time.Since(start)
		avgLatency += latency
		if err != nil {
			logger.Fatal().Err(err).Msg("could not ping")
		}
		logger.Info().Interface("response", r.GetMessage()).Str("latency", latency.String()).Msg("PING response received")
	}
	logger.Info().Int("numRequests", numRequests).Str("avgLatency", fmt.Sprintf("%dµs", avgLatency.Microseconds()/int64(numRequests))).Msg("done")
}
