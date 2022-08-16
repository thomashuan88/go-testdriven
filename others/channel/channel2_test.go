package channel

import (
	"fmt"
	"testing"
	"time"
)

func worker(id int, c chan int) {
	for {
		fmt.Printf("worker %d received %d\n", id, <-c)
	}
}

func chanDemo() {
	c := make(chan int)

	go worker(0, c)
	c <- 1
	c <- 2
	time.Sleep(time.Millisecond)
}

func TestChan(t *testing.T) {
	chanDemo()
}
