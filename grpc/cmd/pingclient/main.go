package main

import (
	"context"
	"flag"
	"os"
	"time"

	"github.com/marcoshack/go-examples/grpc/api/api"
	"github.com/marcoshack/go-examples/timeutil"
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
	latencies := make([]time.Duration, numRequests)
	for i := 0; i < numRequests; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		start := time.Now()
		r, err := c.Ping(ctx, &api.PingRequest{Message: "ping"})
		if err != nil {
			logger.Error().Err(err).Msg("could not ping")
			continue
		}
		latency := time.Since(start)
		latencies[i] = latency
		logger.Info().Interface("response", r.GetMessage()).Str("latency", latency.String()).Msg("PING response received")
	}

	stats := timeutil.NewLatencyStats(latencies)
	logger.Info().EmbedObject(stats).Msg("Latency statistics")
}
