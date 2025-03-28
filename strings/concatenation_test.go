package strings_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/google/uuid"
)

var (
	result string
)

func BenchmarkStringsConcatenation_Plus(b *testing.B) {
	var s string
	for n := 0; n < b.N; n++ {
		s = uuid.New().String() + "#" + uuid.New().String() + "#" + uuid.New().String() + "#" + uuid.New().String()
	}
	result = s
}

func BenchmarkStringsConcatenation_PlusMultiLine(b *testing.B) {
	var s string
	for n := 0; n < b.N; n++ {
		s = uuid.New().String()
		s += "#"
		s += uuid.New().String()
		s += "#"
		s += uuid.New().String()
		s += "#"
		s += uuid.New().String()
	}
	result = s
}

func BenchmarkStringsConcatenation_Sprintf(b *testing.B) {
	var s string
	for n := 0; n < b.N; n++ {
		s = fmt.Sprintf("%s#%s#%s#%s", uuid.New().String(), uuid.New().String(), uuid.New().String(), uuid.New().String())
	}
	result = s
}

func BenchmarkStringsConcatenation_Join(b *testing.B) {
	var s string
	for n := 0; n < b.N; n++ {
		elems := []string{uuid.New().String(), uuid.New().String(), uuid.New().String(), uuid.New().String()}
		s = strings.Join(elems, "#")
	}
	result = s
}

func BenchmarkStringsConcatenation_JoinAppend(b *testing.B) {
	var s string
	for n := 0; n < b.N; n++ {
		elems := make([]string, 0, 4)
		elems = append(elems, uuid.New().String())
		elems = append(elems, uuid.New().String())
		elems = append(elems, uuid.New().String())
		elems = append(elems, uuid.New().String())
		s = strings.Join(elems, "#")
	}
	result = s
}

func BenchmarkStringsConcatenation_Builder(b *testing.B) {
	var s string
	for n := 0; n < b.N; n++ {
		var sb strings.Builder
		sb.WriteString(uuid.New().String())
		sb.WriteString("#")
		sb.WriteString(uuid.New().String())
		sb.WriteString("#")
		sb.WriteString(uuid.New().String())
		sb.WriteString("#")
		sb.WriteString(uuid.New().String())
		s = sb.String()
	}
	result = s
}
