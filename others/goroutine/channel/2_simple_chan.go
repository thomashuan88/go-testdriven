package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for {
		// n := <-c
		fmt.Printf("Worker %d received %c\n", id, <-c)
	}
}

// goroutines deadlock example

func chanDemo() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go worker(i, channels[i])
	}
	// var c chan int // c == nil, cannot use
	// c := make(chan int)

	// go worker(0, c)
	// c <- 1 // send to channel
	// c <- 2

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond) // give time to print 2

}

func main() {
	chanDemo()
}