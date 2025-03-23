package random

import (
	"crypto/rand"
	"fmt"
)

const CodeCharset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateRandomCode(length int) (string, error) {
	// Create a byte slice to store our random bytes
	randomBytes := make([]byte, length)

	// Read random bytes in a single call
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %w", err)
	}

	// Create result array
	result := make([]byte, length)

	// Use each byte to select a character from charset
	for i, b := range randomBytes {
		// Use modulo to map the byte to an index in our charset
		// This ensures even distribution across the charset
		result[i] = CodeCharset[b%byte(len(CodeCharset))]
	}

	return string(result), nil
}
