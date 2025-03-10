package main_test

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestForLoopWithRerence(t *testing.T) {
	origNumbers := []int{1, 2, 3}
	newNumbers := make([]*int, 3, len(origNumbers))

	// see https://go.dev/blog/loopvar-preview
	for i, n := range origNumbers {
		newNumbers[i] = &n
	}

	for i := range origNumbers {
		assert.Equal(t, origNumbers[i], newNumbers[i])
	}
}
