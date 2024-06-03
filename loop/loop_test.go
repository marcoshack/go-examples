package loop

import (
	"fmt"
	"testing"
)

func TestForLoopWithRerence(t *testing.T) {
	origNumbers := []int{1, 2, 3}
	newNumbers := make([]*int, 0, len(origNumbers))

	for _, n := range origNumbers {
		newNumbers = append(newNumbers, &n)
	}

	fmt.Printf("newNumbers: %v", newNumbers)
}
