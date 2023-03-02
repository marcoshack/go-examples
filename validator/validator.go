package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type MyStruct struct {
	Field1 string `validate:"required_without_all=Field2 Field3,required_with=Field4"`
	Field2 string `validate:"required_without_all=Field1 Field3"`
	Field3 string `validate:"required_without_all=Field1 Field2"`
	Field4 string `validate:"required_if=Field5 true"`
	Field5 bool
}

func main() {
	values := make([]MyStruct, 0, 5)
	values = append(values, MyStruct{Field1: "foo"})
	values = append(values, MyStruct{Field2: "bar"})
	values = append(values, MyStruct{Field3: "baz"})
	values = append(values, MyStruct{Field3: "baz", Field1: "foo", Field4: "qux"})
	values = append(values, MyStruct{Field1: "foo", Field5: true, Field4: "qux"})

	validate := validator.New()

	for _, v := range values {
		err := validate.Struct(v)
		if err != nil {
			fmt.Printf(err.Error())
			return
		}
	}

	fmt.Printf("All good.")
}
