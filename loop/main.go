package main

import (
	"fmt"
)

type Foo struct {
	Max int
	Counter int
}

func (f *Foo) HasMore() bool {
	if (f.Counter <= f.Max) {
		return true
	}
	return false
}

func (f *Foo) Next() int {
	current := f.Counter
	f.Counter++
	return current
}

func main() {
	foo := Foo{Max: 10}
	for foo.HasMore() {
		fmt.Printf("Next: %d\n", foo.Next())
	}
}
