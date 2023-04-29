package hooks_test

import (
	"strings"
	"testing"

	"github.com/marcoshack/go-examples/hooks"
	"github.com/stretchr/testify/require"
)

func TestHooks_ToLowerTransformer_NoHooks(t *testing.T) {
	result := hooks.ToLowerTransform(hooks.StringTransformInput{
		Value: " F o O ",
	})
	require.Equal(t, " f o o ", result)
}

func TestHooks_ToLowerTransformer_BeforeAndAfter(t *testing.T) {
	result := hooks.ToLowerTransform(hooks.StringTransformInput{
		Value: " F o O ",
		BeforeTransform: func(value string) string {
			return strings.TrimSpace(value)
		},
		AfterTransform: func(value string) string {
			return strings.ReplaceAll(value, " ", "_")
		},
	})
	require.Equal(t, "f_o_o", result)
}
