package select_test

import (
	"testing"
)

func TestSelect_WithResponseCodes(t *testing.T) {
	responseCode := 400

	switch {
	case responseCode > 400:
	}
}
