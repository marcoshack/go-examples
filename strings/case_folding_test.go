package strings_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkCaseFolding_WithStringsToLower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		assert.True(b, strings.ToLower("Foo") == strings.ToLower("foO"))
	}
}

func BenchmarkCaseFolding_WithStringsEqualFold(b *testing.B) {
	for i := 0; i < b.N; i++ {
		assert.True(b, strings.EqualFold("Foo", "foO"))
	}
}
