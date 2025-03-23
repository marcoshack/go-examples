package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/marcoshack/go-examples/timeutil"
	"github.com/marcoshack/go-examples/twirp/api"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	var clientType, baseURL string
	var numRequests, numWorkers int
	flag.StringVar(&baseURL, "s", "http://localhost:8080", "Server endpoint")
	flag.StringVar(&clientType, "t", "protobuf", "Type of client to use (json or protobuf)")
	flag.IntVar(&numRequests, "n", 5, "Number of requests to make")
	flag.IntVar(&numWorkers, "w", 1, "Number of works to use to send the requests in parallel")
	flag.Parse()

	// initialize zerolog logger
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "15:04:05"}
	logger := zerolog.New(output).With().Timestamp().Logger()
	ctx := logger.WithContext(context.Background())

	// initialize twirp client
	httpClient := &http.Client{}
	var twirpClient api.PingService
	switch clientType {
	case "protobuf":
		twirpClient = api.NewPingServiceProtobufClient(baseURL, httpClient)
	case "json":
		twirpClient = api.NewPingServiceJSONClient(baseURL, httpClient)
	}
	log.Ctx(ctx).Info().Str("client", clientType).Str("baseURL", baseURL).Msg("twirp client initialized")

	var wg sync.WaitGroup
	wg.Add(numWorkers)
	numRequestsPerWorker := numRequests / numWorkers

	start := time.Now()
	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			ping(ctx, twirpClient, numRequestsPerWorker)
		}()
	}

	wg.Wait()
	log.Ctx(ctx).Info().Str("elapsed", time.Since(start).String()).Msg("Done")
}

func ping(ctx context.Context, twirpClient api.PingService, numRequests int) {
	latencies := make([]time.Duration, numRequests)
	for i := 0; i < numRequests; i++ {
		ctx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()
		start := time.Now()
		_, err := twirpClient.Ping(ctx, &api.PingRequest{Message: "Hello, Twirp!"})
		if err != nil {
			log.Ctx(ctx).Error().Err(err).Msg("could not ping")
			continue
		}
		latency := time.Since(start)
		latencies[i] = latency
	}

	stats := timeutil.NewLatencyStats(latencies)
	log.Ctx(ctx).Info().EmbedObject(stats).Msg("Latency statistics")
}
