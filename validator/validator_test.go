package main_test

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type MyStruct struct {
	Field1 string `validate:"required_without_all=Field2 Field3,required_with=Field4,max=256"`
	Field2 string `validate:"required_without_all=Field1 Field3,max=256"`
	Field3 string `validate:"required_without_all=Field1 Field2,max=256"`
	Field4 string `validate:"required_if=Field5 true,max=256"`
	Field5 bool
}

var (
	testValidator = validator.New()
	testStruct    = MyStruct{
		Field1: uuid.NewString(),
		Field2: uuid.NewString(),
		Field3: uuid.NewString(),
		Field4: uuid.NewString(),
		Field5: true,
	}
)

func (s *MyStruct) Validate() error {
	if s.Field1 == "" && s.Field2 == "" && s.Field3 == "" {
		return fmt.Errorf("Field1, Field2 or Field3 is required")
	}
	if s.Field4 != "" && s.Field1 == "" {
		return fmt.Errorf("Field1 is required if Field4 is set")
	}
	if s.Field5 && s.Field4 == "" {
		return fmt.Errorf("Field4 is required if Field5 is true")
	}
	if s.Field1 != "" && len(s.Field1) > 256 {
		return fmt.Errorf("Field1 max size is 256")
	}
	if s.Field2 != "" && len(s.Field2) > 256 {
		return fmt.Errorf("Field2 max size is 256")
	}
	if s.Field3 != "" && len(s.Field3) > 256 {
		return fmt.Errorf("Field3 max size is 256")
	}
	if s.Field4 != "" && len(s.Field4) > 256 {
		return fmt.Errorf("Field4 max size is 256")
	}
	return nil
}

func TestField1Or2Or3MustBePresent(t *testing.T) {
	IsValid(t, MyStruct{Field1: "foo"})
	IsValid(t, MyStruct{Field2: "foo"})
	IsValid(t, MyStruct{Field3: "foo"})
	NotValid(t, MyStruct{})
}

func TestOnlyField2(t *testing.T) {
	IsValid(t, MyStruct{Field2: "foo"})
}

func TestOnlyField3(t *testing.T) {
	IsValid(t, MyStruct{Field3: "foo"})
}

func TestFields123(t *testing.T) {
	IsValid(t, MyStruct{Field1: "foo", Field2: "bar", Field3: "baz"})
}

func TestField5IsTrueThenField4ThenField1(t *testing.T) {
	IsValid(t, MyStruct{Field5: true, Field4: "foo", Field1: "bar"})
}

func IsValid(t *testing.T, o MyStruct) {
	validate(t, o, false)
}

func NotValid(t *testing.T, o MyStruct) {
	validate(t, o, true)
}

func validate(t *testing.T, o MyStruct, isError bool) {
	err := testValidator.Struct(o)
	if (err != nil) != isError {
		t.Errorf(err.Error())
	}
	err = o.Validate()
	if (err != nil) != isError {
		t.Errorf(err.Error())
	}
}

func BenchmarkValidator(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testValidator.Struct(testStruct)
	}
}

func BenchmarkCustomIfs(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testStruct.Validate()
	}
}
