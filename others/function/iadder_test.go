package function_test

import "testing"

type iAdder func(int) (int, iAdder)

// original functional programming
func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func TestAdder2(t *testing.T) {
	a := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, a = a(i)
		t.Logf("0 + 1 + ... + %d = %d\n", i, s)
	}
}
