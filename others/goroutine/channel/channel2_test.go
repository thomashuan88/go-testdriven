package channel

import (
	"fmt"
	"testing"
	"time"
)

func worker(id int, c chan int) {
	for {
		d, more := <-c
		if !more {
			return
		}
		fmt.Printf("worker %d received %c\n", id, d)
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
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

// indicate the channel already finish
func channelClose() {
	c := make(chan int, 3) // you can add 3 data , without hit deadlock
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	// close(c)
	time.Sleep(time.Millisecond)
}

func TestChan(t *testing.T) {
	// chanDemo()
	// bufferedChannel()
	channelClose()
}
