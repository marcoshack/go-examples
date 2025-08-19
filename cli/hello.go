package cli

import (
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var HelloCmd = &cobra.Command{
	Use:   "hello [name]",
	Short: "Print a friendly greeting",
	Long: `Print a friendly greeting message. You can optionally provide a name
to personalize the greeting.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := "World"
		if len(args) > 0 {
			name = args[0]
		}

		uppercase, _ := cmd.Flags().GetBool("uppercase")

		log.Info().
			Str("greeting_target", name).
			Bool("uppercase", uppercase).
			Msg("Generating greeting")

		greeting := "Hello, " + name + "!"
		if uppercase {
			greeting = strings.ToUpper(greeting)
		}

		log.Info().
			Str("message", greeting).
			Msg("Greeting generated successfully")
	},
}

func init() {
	HelloCmd.Flags().BoolP("uppercase", "u", false, "Print greeting in uppercase")
}