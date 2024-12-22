package main

import (
	"net/http"
	"os"

	"github.com/marcoshack/go-examples/twirp/api/api"
	"github.com/marcoshack/go-examples/twirp/api/server"
	"github.com/rs/zerolog"
)

func main() {
	// initialize zerolog logger
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "15:04:05"}
	logger := zerolog.New(output).With().Timestamp().Logger()

	// initialize twirp server and register service
	service := server.NewPingService()
	twirpHandler := api.NewPingServiceServer(service)

	// TODO read from CLI argument
	bindAddr := ":8080"
	logger.Info().Str("bindAddr", bindAddr).Msg("server listening")
	http.ListenAndServe(bindAddr, twirpHandler)
}
