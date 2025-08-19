package main

import (
	"os"

	"github.com/marcoshack/go-examples/cli"
	"github.com/rs/zerolog/log"
)

func main() {
	if err := cli.Execute(); err != nil {
		log.Error().Err(err).Msg("Failed to execute command")
		os.Exit(1)
	}
}