package main

import (
	"fmt"
	"math/rand"
	"time"
)

// message generator
// <-chan make sure channel just can receive
func msgGen() <-chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			c <- fmt.Sprintf("message %d", i)
			i++
		}
	}()
	return c
}

func main() {
	m := msgGen()
	for {
		fmt.Println(<-m)
		m <- "abc" // like this cannot, caused by <-chan
	}
}
