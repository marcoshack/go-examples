package atomic_test

import (
	"sync/atomic"
	"testing"
)

func TestAtomic_Bool(t *testing.T) {
	var a atomic.Bool

	if a.Load() {
		t.Error("a.Load() should be false")
	}

	a.Store(true)

	if !a.Load() {
		t.Error("a.Load() should be true")
	}
}
