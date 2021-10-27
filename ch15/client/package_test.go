package client

import (
	"go-testdriven/ch15/series"
	"testing"
)

func TestPackage(t *testing.T) {
	// t.Log(Series.GetFibonacci(5))
	t.Log(series.GetFibonacci(5))
	t.Log(series.Square(5))
}
