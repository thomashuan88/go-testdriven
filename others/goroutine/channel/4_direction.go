package main

import (
	"fmt"
	"time"
)

func createWorker(id int) chan<- int {

	c := make(chan int)
	go func() {
		for {
			fmt.Printf("Worker %d received %c\n", id, <-c)
		}
	}()
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

func main() {
	chanDemo()
}
