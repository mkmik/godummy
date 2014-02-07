package godummy

import (
	"testing"
	"testing/quick"
)

func TestFoo(t *testing.T) {
	f := func(x int) bool {
		y := Foo(x)
		return y == x*2
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
