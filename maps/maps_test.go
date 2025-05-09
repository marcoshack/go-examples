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

func TestMaps_LookupInUnitializedMap(t *testing.T) {
	var myMap map[string]bool
	require.False(t, myMap["foo"], "non existing map key should return false")
	require.NotContains(t, myMap, "foo")
}

func TestMaps_StructDefaultValues(t *testing.T) {
	type myType struct {
		key   string
		value string
	}
	myMap := make(map[string]myType)
	myMap["foo"] = myType{key: "k", value: "v"}

	value, found := myMap["baz"]
	require.False(t, found)
	require.NotNil(t, value)
}

func TestMaps_WithEmptyStringKeys(t *testing.T) {
	myMap := make(map[string][]string)
	myMap[""] = []string{"foo", "bar"}
	myMap["key1"] = []string{"baz", "qux", "quux"}

	require.Len(t, myMap[""], 2)
	require.Len(t, myMap["key1"], 3)
}
