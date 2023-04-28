package modules_test

import (
	"testing"

	v1 "github.com/marcoshack/go-examples/modules"
	v2 "github.com/marcoshack/go-examples/modules/v2"
	"github.com/stretchr/testify/require"
)

func TestModulesV1(t *testing.T) {
	require.Equal(t, "foo/v1", v1.Foo())
}

func TestModulesV2(t *testing.T) {
	require.Equal(t, "foo/v2", v2.SayFoo())
}
