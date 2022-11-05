package main

import (
	"fmt"
	"math/rand"
	"time"
)

// message generator
// <-chan make sure channel just can receive
func msgGen() chan string {
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
	m1 := msgGen() // similar like handle
	m2 := msgGen()
	for {
		fmt.Println(<-m1)
		fmt.Println(<-m2)
	}
}
