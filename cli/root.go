package cli

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "cli",
	Short: "A sample CLI application built with Cobra",
	Long: `A sample CLI application that demonstrates the use of Cobra package
with zerolog for structured logging and colorized output.`,
}

func init() {
	// Configure zerolog for colorized console output with short timestamp
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.Kitchen, // Short time format like "3:04PM"
	}
	log.Logger = zerolog.New(output).With().Timestamp().Logger()

	// Add subcommands
	RootCmd.AddCommand(HelloCmd)
	RootCmd.AddCommand(QuoteCmd)
}

func Execute() error {
	return RootCmd.Execute()
}