package main

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	var writer io.Writer
	writer = os.Stdout

	if os.Getenv("CONSOLE_LOG") == "1" {
		writer = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: "2006-01-02T15:04:05.000000000",
		}
	}

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixNano
	zerolog.TimestampFunc = func() time.Time {
		return time.Now().UTC()
	}
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	logger := zerolog.New(writer).With().Timestamp().Stack().Logger()

	ctx := logger.WithContext(context.Background())

	for i := range 10 {
		log.Ctx(ctx).Info().Int("i", i).Msg("Hello World")
	}
}
