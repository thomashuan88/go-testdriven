package main

import (
	"fmt"
	"math/rand"
	"time"
)

// message generator
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
	m := msgGen()
	for {
		fmt.Println(<-m)
	}
}
