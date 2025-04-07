package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Example of a long running process that spawns a goroutine every second and properly handle panics in the underlying function.
// The function used here (doSometihng) "randomly" panic to showcase the recovery process.
func main() {
	logWritter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339, NoColor: false}
	ctx := zerolog.New(logWritter).With().Timestamp().Logger().WithContext(context.Background())
	log.Ctx(ctx).Info().Msg("starting...")

	panicHandler := NewPanicHandler()
	ticker := time.Tick(1 * time.Second)

	for range ticker {
		log.Ctx(ctx).Info().Msg("TICK")
		go func() {
			defer panicHandler.Recover(ctx)
			doSomething(ctx)
		}()
	}

	log.Ctx(ctx).Info().Msg("stopped")
}

type PanicHandler struct {
	// ...
}

func NewPanicHandler() *PanicHandler {
	// initialize dependencies, if needed
	return &PanicHandler{}
}

func (h *PanicHandler) Recover(ctx context.Context) {
	// Just in case the Recover iself panic. Here we don't use context, logger or anything from the handler.
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("ERROR: panic while recovering from panic: %v\n", r)
		}
	}()

	if r := recover(); r != nil {
		log.Ctx(ctx).Error().Msgf("recovered from panic: %v", r)
		// further recover actions, e.g. emit metrics

		if time.Now().UTC().Nanosecond()%2 == 0 {
			panic("double panic!")
		}
	}
}

func doSomething(ctx context.Context) {
	log.Ctx(ctx).Info().Msg("running goroutine")

	// simulate some random workload from 1-1000ms
	randInt, err := rand.Int(rand.Reader, big.NewInt(1000))
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("failed to generate random number")
		return
	}
	workInMilliseconds := randInt.Int64()

	log.Ctx(ctx).Info().Int64("workInMilliseconds", workInMilliseconds).Msg("doing some work in goroutine")
	time.Sleep(time.Duration(workInMilliseconds) * time.Millisecond)

	if workInMilliseconds%2 == 0 {
		panic("panic from goroutine!") // This goroutine will die
	}

	log.Ctx(ctx).Info().Msg("goroutine done")
}
