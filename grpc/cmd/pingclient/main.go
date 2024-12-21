package main

import (
	"context"
	"time"

	"github.com/marcoshack/go-examples/grpc/api/api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal().Err(err).Msg("did not connect")
	}
	defer conn.Close()
	c := api.NewPingServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	start := time.Now()
	r, err := c.Ping(ctx, &api.PingRequest{Message: "ping"})
	elapsed := time.Since(start)
	if err != nil {
		log.Fatal().Err(err).Msg("could not ping")
	}
	log.Info().Interface("response", r.GetMessage()).Str("elapsed", elapsed.String()).Msg("response received")
}
