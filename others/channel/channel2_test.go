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

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
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

// channel with buffer can help improve performance
func bufferedChannel() {
	c := make(chan int, 3) // you can add 3 data , without hit deadlock
	go worker(0, c)
	c <- 1
	c <- 2
	c <- 3
	time.Sleep(time.Millisecond)
}

func TestChan(t *testing.T) {
	// chanDemo()
	bufferedChannel()
}
