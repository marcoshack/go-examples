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

func BenchmarkPlus(b *testing.B) {
	var s string
	for n := 0; n < b.N; n++ {
		s = uuid.New().String() + "#" + uuid.New().String() + "#" + uuid.New().String() + "#" + uuid.New().String()
	}
	result = s
}

func BenchmarkPlusMultiLine(b *testing.B) {
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

func BenchmarkSprintf(b *testing.B) {
	var s string
	for n := 0; n < b.N; n++ {
		s = fmt.Sprintf("%s#%s#%s#%s", uuid.New().String(), uuid.New().String(), uuid.New().String(), uuid.New().String())
	}
	result = s
}

func BenchmarkJoin(b *testing.B) {
	var s string
	for n := 0; n < b.N; n++ {
		elems := []string{uuid.New().String(), uuid.New().String(), uuid.New().String(), uuid.New().String()}
		s = strings.Join(elems, "#")
	}
	result = s
}

func BenchmarkJoinAppend(b *testing.B) {
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

func BenchmarkBuilder(b *testing.B) {
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
