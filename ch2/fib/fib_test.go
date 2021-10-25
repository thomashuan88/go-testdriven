package fib

import (
	"testing"
)

func TestFibList(t *testing.T) {
	// var a int = 1
	// var b int = 1
	// var (
	// 	a int = 1
	// 	b = 1
	// )

	a := 1
	b := 1

	// fmt.Print(a)
	t.Log()
	for i := 0; i < 5; i++ {
		// fmt.Print(" ", b)
		t.Log(" ", b)
		temp := a
		a = b
		b = temp + a
	}
	// fmt.Println()

}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	// tmp := a
	// a = b
	// b = tmp
	a, b = b, a
	t.Log(a, b)
}

func TestConst(t *testing.T) {
	const (
		Open = 1 << iota
		Close
		Pending
	)
	t.Logf("open %b", Open)
	t.Logf("close %b", Close)
	t.Logf("pending %b", Pending)
}
