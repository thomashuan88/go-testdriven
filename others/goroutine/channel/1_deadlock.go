package channel

import "fmt"

// goroutines deadlock example

func chanDemo() {
	// var c chan int // c == nil, cannot use
	c := make(chan int)
	c <- 1 // send to channel
	c <- 2
	n := <-c // get from channel
	fmt.Println(n)
}
