package groutine_test

import (
	"fmt"
	"testing"
	"time"
)

func TestGroutine(t *testing.T) {
	for i := 1; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
			time.Sleep(time.Millisecond * time.Duration(i)) // cannot directly * i, have to * time.Duration(i)
		}(i)

	}
	time.Sleep(time.Millisecond * 50)
}
