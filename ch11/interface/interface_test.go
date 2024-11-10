package interface_test

import (
	"testing"
)

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

// 1. not depend on interface
// 2. can declare interface at user package

func TestClient(t *testing.T) {
	// var p Programmer  -- should merge variable declaration with assignment on next line (S1021)
	// p = new(GoProgrammer)
	var p Programmer = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}
