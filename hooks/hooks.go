package hooks

import (
	"strings"
)

type stringTransformer func(string) string

type StringTransformInput struct {
	Value           string
	BeforeTransform stringTransformer
	AfterTransform  stringTransformer
}

func ToLowerTransform(input StringTransformInput) string {
	transformed := input.Value

	if input.BeforeTransform != nil {
		transformed = input.BeforeTransform(input.Value)
	}

	transformed = strings.ToLower(transformed)

	if input.AfterTransform != nil {
		transformed = input.AfterTransform(transformed)
	}

	return transformed
}
