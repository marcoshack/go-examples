package cli

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// Constants
const (
	DefaultModel = "anthropic/claude-sonnet-4"
)

// OpenRouter API structures
type OpenRouterRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenRouterResponse struct {
	ID      string   `json:"id"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}

type Usage struct {
	PromptTokens     int     `json:"prompt_tokens"`
	CompletionTokens int     `json:"completion_tokens"`
	TotalTokens      int     `json:"total_tokens"`
	TotalCost        float64 `json:"total_cost"`
}

type Choice struct {
	Message Message `json:"message"`
}

var QuoteCmd = &cobra.Command{
	Use:   "quote",
	Short: "Generate a quote of the day using AI",
	Long: `Generate an inspirational quote of the day using OpenRouter AI.
Requires OPENROUTER_API_KEY environment variable to be set.`,
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("OPENROUTER_API_KEY")
		if apiKey == "" {
			log.Error().Msg("OPENROUTER_API_KEY environment variable is required")
			os.Exit(1)
		}

		category, _ := cmd.Flags().GetString("category")
		model, _ := cmd.Flags().GetString("model")

		log.Info().
			Str("category", category).
			Str("model", model).
			Msg("Generating quote of the day")

		quote, cost, err := generateQuote(apiKey, category, model)
		if err != nil {
			log.Error().Err(err).Msg("Failed to generate quote")
			os.Exit(1)
		}

		log.Info().
			Str("quote", quote).
			Float64("cost_usd", cost).
			Msg("Quote generated successfully")
	},
}

func generateQuote(apiKey, category, model string) (string, float64, error) {
	prompt := fmt.Sprintf("Generate an inspirational quote of the day about %s. Return only the quote with attribution if known, nothing else.", category)

	requestBody := OpenRouterRequest{
		Model: model,
		Messages: []Message{
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		log.Error().Err(err).Msg("Failed to marshal request")
		return "", 0.0, errors.New("failed to marshal request")
	}

	req, err := http.NewRequest("POST", "https://openrouter.ai/api/v1/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Error().Err(err).Msg("Failed to create request")
		return "", 0.0, errors.New("failed to create request")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("HTTP-Referer", "https://github.com/go-examples/cli")
	req.Header.Set("X-Title", "Go CLI Example")

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	log.Debug().Msg("Sending request to OpenRouter API")

	resp, err := client.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("Failed to send request")
		return "", 0.0, errors.New("failed to send request")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Error().
			Int("status_code", resp.StatusCode).
			Str("response_body", string(body)).
			Msg("API request failed")
		return "", 0.0, errors.New("API request failed")
	}

	var response OpenRouterResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Error().Err(err).Msg("Failed to decode response")
		return "", 0.0, errors.New("failed to decode response")
	}

	if len(response.Choices) == 0 {
		log.Error().Msg("No choices returned from API")
		return "", 0.0, errors.New("no choices returned from API")
	}

	return response.Choices[0].Message.Content, response.Usage.TotalCost, nil
}

func init() {
	QuoteCmd.Flags().StringP("category", "c", "motivation", "Category for the quote (e.g., motivation, success, life)")
	QuoteCmd.Flags().StringP("model", "m", DefaultModel, "AI model to use for generation")
}