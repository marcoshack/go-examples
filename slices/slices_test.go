package slices_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MyStruct struct {
	Values   []string
	Pointers []*string
}

func TestSlices_DefaultSliceObjects(t *testing.T) {
	myStruct := MyStruct{}
	myStruct.Values = append(myStruct.Values, "a")
	myStruct.Values = append(myStruct.Values, "b")
	myStruct.Values = append(myStruct.Values, "c")

	myStruct.Pointers = append(myStruct.Pointers, stringPtr("a"))
	myStruct.Pointers = append(myStruct.Pointers, stringPtr("b"))
	myStruct.Pointers = append(myStruct.Pointers, stringPtr("c"))

	assert.Len(t, myStruct.Values, 3)
	assert.Len(t, myStruct.Pointers, 3)
}

func stringPtr(s string) *string {
	return &s
}
