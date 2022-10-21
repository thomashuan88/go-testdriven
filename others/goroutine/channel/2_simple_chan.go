package main

import (
	"fmt"
	"time"
)

// goroutines deadlock example

func chanDemo() {
	// var c chan int // c == nil, cannot use
	c := make(chan int)

	go func() {
		for {
			n := <-c
			fmt.Println(n)
		}
	}()
	c <- 1 // send to channel
	c <- 2
	time.Sleep(time.Millisecond) // give time to print 2

}

func main() {
	chanDemo()
}
