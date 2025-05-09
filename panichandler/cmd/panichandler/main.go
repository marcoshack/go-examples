package main

import (
	"context"
	"errors"
	"os"

	"github.com/marcoshack/go-examples/panichandler"
	"github.com/rs/zerolog"
)

func main() {
	//ctx := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).WithContext(context.Background())
	ctx := zerolog.New(os.Stdout).WithContext(context.Background())
	defer panichandler.NewPanicHandler().Recover(ctx)

	func1()
}

func func1() {
	func2()
}

func func2() {
	func3()
}

func func3() {
	panic(errors.New("panic: oops!"))
}
