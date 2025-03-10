package maps_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaps_IncrementValuesToNonExistingKeys(t *testing.T) {
	myMap := make(map[string]int)
	myMap["foo"]++
	myMap["bar"]++
	myMap["bar"]++

	require.Equal(t, 1, myMap["foo"])
	require.Equal(t, 2, myMap["bar"])
	require.Equal(t, 0, myMap["baz"])
}
