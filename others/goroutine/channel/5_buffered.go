package main

import (
	"fmt"
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

// goroutines deadlock example
// chan<- int (send only channel)
// <-chan int (receive only channel)
func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
		// n := <- channels[i] // this one not working,
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond) // give time to print 2

}

func bufferedChannel() {
	c := make(chan int, 3) // you can add 3 data , without hit deadlock
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

func main() {
	// chanDemo()
	bufferedChannel()
}
