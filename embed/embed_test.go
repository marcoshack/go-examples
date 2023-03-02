package embed_test

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/google/uuid"
)

type Foo struct {
	attribute1 string
}

func (f *Foo) Method1() string {
	return fmt.Sprintf("Foo::Method1")
}

func (f *Foo) Method2() string {
	return fmt.Sprintf("Foo::Method2")
}

func (f *Foo) Method3() string {
	return f.attribute1
}

type Bar struct {
	Foo
}

func (b *Bar) Method2() string {
	return fmt.Sprintf("Bar::Method2")
}

func NewBar(attribute1 string) *Bar {
	b := &Bar{
		Foo{
			attribute1: attribute1,
		},
	}

	// alternatively
	// b := &Bar{}
	// b.attribute1 = attribute1

	return b
}

func TestEmbed(t *testing.T) {
	expectedAttribute1Value := uuid.NewString()
	b := NewBar(expectedAttribute1Value)
	assert.Equal(t, "Foo::Method1", b.Method1())
	assert.Equal(t, "Bar::Method2", b.Method2())
	assert.Equal(t, expectedAttribute1Value, b.Method3())
	b.Method2()
}
