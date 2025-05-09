package misc_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type Hooks struct {
	Func1 func() string
	Func2 func() string
	Func3 func() string
}

type MyHooks1 struct {
	value string
}

func (h *MyHooks1) Func1() string {
	return h.value
}

type MyHooks2 struct {
	value string
}

func (h *MyHooks2) Func2() string {
	return h.value
}

func TestFunctionAsAttributes(t *testing.T) {
	myHooks1 := MyHooks1{value: "MyHooks1"}
	myHooks2 := MyHooks2{value: "MyHooks2"}

	hooks := Hooks{
		Func1: myHooks1.Func1,
		Func2: myHooks2.Func2,
	}
	require.Equal(t, "MyHooks1", hooks.Func1())
	require.Equal(t, "MyHooks2", hooks.Func2())
	require.Nil(t, hooks.Func3)
}
