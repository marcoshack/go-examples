package process

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
)

func StartShutdownHandler(message string, cancel context.CancelFunc) {
	go func() {
		shutdownSignalChannel := make(chan os.Signal, 1)
		signal.Notify(shutdownSignalChannel, syscall.SIGINT, syscall.SIGTERM)
		sig := <-shutdownSignalChannel
		log.Info().Str("signal", sig.String()).Msg(message)
		cancel()
	}()
}
