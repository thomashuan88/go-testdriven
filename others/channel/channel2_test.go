package channel

import (
	"fmt"
	"testing"
	"time"
)

func createWorker(id int) chan<- int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("worker %d received %d\n", id, <-c)
		}
	}()

	return c
}

func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int)
	c <- 1
}

func TestChan(t *testing.T) {
	// chanDemo()
	bufferedChannel()
}
