package goroutine_test

import (
	"fmt"
	"testing"
	"time"
)

func count() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 1)
	}
}

func TestGoroutineCount(t *testing.T) {
	go count()
	time.Sleep(time.Millisecond * 2)
	t.Log("hello world")
	time.Sleep(time.Millisecond * 5)
}
