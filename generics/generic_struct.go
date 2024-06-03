package generics

import "fmt"

func NewMyStruct[T any](value T) *MyStruct[T] {
	return &MyStruct[T]{
		Value: value,
	}
}

// MyStruct declares a genefic type that can be used in any of its methods.
type MyStruct[T any] struct {
	Value T
}

func (s *MyStruct[T]) Print() {
	fmt.Printf("MyStruct[value=%v]", s.Value)
}

// But struct methods cannot declare other types, it can only use the types declared by the struct
// Error: method must have no type parameters
//func (s *MyStruct[T]) PrintAnotherType[S any](anotherValue S) {
//}

// AnotherStruct is a non-generic struct, but even so a method of this struct cannot declare types
type AnotherStruct struct {
	// ...
}

// error: method must have no type parameters
//func (s *AnotherStruct) Print[T any](value T) {
//	fmt.Printf("AnotherStruct[value=%v]", value)
//}
