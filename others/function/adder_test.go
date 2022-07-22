package function_test

import (
	"testing"
)

func TestAdder(t *testing.T) {
	a := adder()
	for i := 0; i < 10; i++ {
		t.Logf("0 + 1 + ... + %d = %d\n", i, a(i))
	}
}

func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}
