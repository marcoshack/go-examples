package main

import (
	"context"
	"sync"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	logger := zerolog.New(zerolog.NewConsoleWriter()).With().Timestamp().Stack().Logger()
	zerolog.DefaultContextLogger = &logger
	log.Logger = logger

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Ctx(context.Background()).Info().Msg("Logging with log.Ctx(context.Background) within goroutine")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Info().Msg("Logging with log.Info() within goroutine")
	}()

	wg.Wait()
	logger.Info().Msg("Done")
}
