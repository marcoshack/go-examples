package random_test

import (
	"fmt"
	"testing"

	"github.com/marcoshack/go-examples/random"
	"github.com/stretchr/testify/require"
)

func BenchmarkGenerateRandomCode(b *testing.B) {
	// Test cases with different lengths
	lengths := []int{6, 8, 12, 16, 32}

	for _, length := range lengths {
		b.Run(fmt.Sprintf("Length_%d", length), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				code, err := random.GenerateRandomCode(length)
				require.NoError(b, err)
				require.NotEmpty(b, code)
			}
		})
	}
}

// Test the output format
func TestGenerateRandomCode(t *testing.T) {
	testCases := []struct {
		name   string
		length int
	}{
		{"empty", 0},
		{"small", 6},
		{"medium", 16},
		{"large", 32},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			code, err := random.GenerateRandomCode(tc.length)
			require.NoError(t, err)

			// Check length
			require.Equal(t, tc.length, len(code))

			// Check if all characters are valid
			for i := 0; i < len(code); i++ {
				require.Contains(t, random.CodeCharset, string(code[i]))
			}
		})
	}
}

func TestGenerateRandomCodeDuplicates(t *testing.T) {
	testCases := []struct {
		codeLength int
		numCodes   int
		maxDups    float64 // maximum acceptable duplicate percentage
	}{
		{
			codeLength: 4,
			numCodes:   10000,
			maxDups:    0.1, // 10% duplicates maximum
		},
		{
			codeLength: 6,
			numCodes:   100000,
			maxDups:    0.01, // 1% duplicates maximum
		},
		{
			codeLength: 8,
			numCodes:   500000,
			maxDups:    0.001, // 0.1% duplicates maximum
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Length_%d_%dk", tc.codeLength, tc.numCodes/1000), func(t *testing.T) {
			// Use a map to track unique codes
			uniqueCodes := make(map[string]int)

			// Generate the specified number of codes
			for i := 0; i < tc.numCodes; i++ {
				code, err := random.GenerateRandomCode(tc.codeLength)
				require.NoError(t, err)
				uniqueCodes[code]++
			}

			// Calculate statistics
			duplicates := 0
			for _, count := range uniqueCodes {
				if count > 1 {
					duplicates += (count - 1)
				}
			}

			dupPercentage := float64(duplicates) / float64(tc.numCodes)

			// Log the results
			t.Logf("Total codes generated: %d", tc.numCodes)
			t.Logf("Unique codes: %d", len(uniqueCodes))
			t.Logf("Duplicate codes: %d", duplicates)
			t.Logf("Duplicate percentage: %.4f%%", dupPercentage*100)

			// Assert that the duplicate percentage is below the maximum allowed
			require.LessOrEqual(t, dupPercentage, tc.maxDups,
				"Too many duplicates generated. Got %.4f%% duplicates, maximum allowed is %.4f%%",
				dupPercentage*100, tc.maxDups*100)
		})
	}
}
