package maps_test

import (
	"sort"
	"testing"

	"github.com/marcoshack/go-examples/maps"
	"github.com/stretchr/testify/require"
)

var (
	slice1 = []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	slice2 = []string{"i", "j", "k", "l", "m", "n", "o", "p"}
	slice3 = []string{"q", "r", "s", "t", "a", "b", "c"}
	slice4 = []string{"u", "v", "w", "y", "a", "b", "c"}

	map1 = map[string]int{"a": 1, "b": 1, "c": 1, "d": 1, "e": 1, "f": 1, "g": 1, "h": 1}
	map2 = map[string]int{"i": 1, "j": 1, "k": 1, "l": 1, "m": 1, "n": 1, "o": 1, "p": 1}
	map3 = map[string]int{"q": 1, "r": 1, "s": 1, "t": 1, "u": 1, "v": 1, "w": 1, "y": 1}
	map4 = map[string]int{"q": 1, "r": 1, "s": 1, "t": 1, "a": 1, "b": 1, "c": 1}

	expectedDuplicates = []string{"a", "b", "c"}
)

func TestMaps_StringIntersection_Empty(t *testing.T) {
	duplicates, found := maps.StringIntersection(slice1, slice2)
	require.False(t, found, "StringIntersection should return false for non intersecting slices")
	require.Empty(t, duplicates, "StringIntersection should return an empty result non intersecting slices")
}

func TestMaps_StringIntersection_WithDuplications(t *testing.T) {
	duplicates, found := maps.StringIntersection(slice1, slice2, slice3, slice4)
	require.True(t, found, "StringIntersection should return true for intersecting slices")
	sort.Strings(duplicates)
	require.Equal(t, expectedDuplicates, duplicates, "StringIntersection should return the expected duplicates")
}

func TestMaps_Intersection_Empty(t *testing.T) {
	duplicates, found := maps.Intersection(slice1, slice2)
	require.False(t, found, "Intersection should return false for non intersecting slices")
	require.Empty(t, duplicates, "Intersection should return an empty result non intersecting slices")
}

func TestMaps_Intersection_WithDuplications(t *testing.T) {
	duplicates, found := maps.Intersection(slice1, slice2, slice3, slice4)
	require.True(t, found, "Intersection should return true for intersecting slices")
	sort.Strings(duplicates)
	require.Equal(t, expectedDuplicates, duplicates, "Intersection should return the expected duplicates")
}

func BenchmarkMaps_StringIntersectionString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		maps.StringIntersection(slice1, slice2, slice3, slice4)
	}
}

func BenchmarkMaps_StringIntersection(b *testing.B) {
	for n := 0; n < b.N; n++ {
		maps.Intersection(slice1, slice2, slice3, slice4)
	}
}

func BenchmarkMaps_HasKeyDuplicationWithCountLoop_WithNoDuplication(b *testing.B) {
	for n := 0; n < b.N; n++ {
		maps.HasKeyDuplicationWithCountLoop(map1, map2, map3)
	}
}

func BenchmarkMaps_HasKeyDuplicationWithCountLoop_WithDuplication(b *testing.B) {
	for n := 0; n < b.N; n++ {
		maps.HasKeyDuplicationWithCountLoop(map1, map2, map4)
	}
}

func BenchmarkMaps_HasKeyDuplicationWithEarlyExit_WithNoDuplication(b *testing.B) {
	for n := 0; n < b.N; n++ {
		maps.HasKeyDuplicationWithEarlyExit(map1, map2, map3)
	}
}

func BenchmarkMaps_HasKeyDuplicationWithEarlyExit_WithDuplication(b *testing.B) {
	for n := 0; n < b.N; n++ {
		maps.HasKeyDuplicationWithEarlyExit(map1, map2, map4)
	}
}
