package main

import (
	"context"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Example of zerolog context fields, including overrides of the same field name.
// The expected output is something like:
//
//	INF message from main field1=value0
//	INF message from method1 field1=value1 field2=value2
//	INF message from method2 field1=value1 field2=value2 field3=value3
func main() {
	// setup a logger with console output and a 'field1' in the context
	ctx := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).
		With().Timestamp().Str("field1", "value0").Logger().
		WithContext(context.Background())

	log.Ctx(ctx).Info().Msg("message from main")
	method1(ctx)
}

func method1(ctx context.Context) {
	// overrides 'field1'
	ctx = log.Ctx(ctx).With().Str("field1", "value1").Logger().WithContext(ctx)

	// adds a new 'field2'
	ctx = log.Ctx(ctx).With().Str("field2", "value2").Logger().WithContext(ctx)

	log.Ctx(ctx).Info().Msg("message from method1")

	method2(ctx)
}

func method2(ctx context.Context) {
	// adds a new 'field3'
	ctx = log.Ctx(ctx).With().Str("field3", "value3").Logger().WithContext(ctx)

	log.Ctx(ctx).Info().Msg("message from method2")
}
