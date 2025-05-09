package panichandler

import (
	"context"
	"fmt"
	"runtime/debug"
	"strings"

	"github.com/rs/zerolog/log"
)

type PanicHandler struct {
	// ...
}

func NewPanicHandler() *PanicHandler {
	// initialize dependencies, if needed
	return &PanicHandler{}
}

func (h *PanicHandler) Recover(ctx context.Context) {
	// Just in case the Recover iself panic. Here we don't use context, logger or anything from the handler.
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("PANIC HANDLER: panic while recovering from panic: %v\n", r)
		}
	}()

	if r := recover(); r != nil {
		if ctx == nil {
			fmt.Printf("PANIC HANDLER: invalid nil context in recover function. panic: %v\n", r)
			return
		}

		debug.PrintStack()

		stackStr := string(debug.Stack())
		stackStr = strings.ReplaceAll(stackStr, "\t", ": ")
		stack := strings.Split(stackStr, "\n")

		formattedStack := make([]string, len(debug.Stack()))
		for i, line := range stack {
			formattedStack[i] = strings.ReplaceAll(line, "\t", "")
		}
		log.Ctx(ctx).Error().Err(errorFromPanic(r)).Bytes("stack", debug.Stack()).Msgf("recovered from panic")
	}
}

func errorFromPanic(r interface{}) error {
	if err, ok := r.(error); ok {
		return err
	}
	return fmt.Errorf("panic: %v", r)
}
