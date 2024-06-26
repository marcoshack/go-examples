package json_test

import (
	"encoding/json"
	"os"
	"testing"

	myjson "github.com/marcoshack/go-examples/json"
	"github.com/stretchr/testify/require"
)

// TestJsonDecodeWithCustomJsonTags validates that json.Unmarshal uses the `json:"attrName"` tags and
// we're able to use a different attribute name in the json file then the attribute name of the struct.
func TestJsonDecodeWithCustomJsonTags(t *testing.T) {
	jsonContent, err := os.ReadFile("file.json")
	require.NoError(t, err)
	require.NotNil(t, jsonContent)

	var aStruct myjson.AStruct
	json.Unmarshal(jsonContent, &aStruct)
	require.Equal(t, "value1", aStruct.Attr1)

	require.NotNil(t, aStruct.Another)
	require.Equal(t, "anotherValue1", aStruct.Another.AnotherAttr1)
}
