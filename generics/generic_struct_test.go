package generics_test

import (
	"testing"

	"github.com/marcoshack/go-examples/generics"
	"github.com/stretchr/testify/require"
)

func TestGenerics_NewMyStruct(t *testing.T) {
	s := generics.NewMyStruct[string]("foo")
	require.NotNil(t, s)
}

func TestGenerics_Print(t *testing.T) {
	s := &generics.MyStruct[string]{Value: "foo"}
	s.Print()
}
