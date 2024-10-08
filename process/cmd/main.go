package main

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/marcoshack/go-examples/process"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	process.StartShutdownHandler("process shutdown", cancel)

	for {
		log.Info().Msg("processing stuff...")
		time.Sleep(2 * time.Second)
		if ctx.Err() != nil {
			break
		}
	}
}
