package main

import (
	"fmt"
	"math/rand"
	"time"
)

// message generator
// <-chan make sure channel just can receive
func msgGen(name string) chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			c <- fmt.Sprintf("service %s: message %d", name, i)
			i++
		}
	}()
	return c
}

func main() {
	m1 := msgGen("service1") // similar like handle
	m2 := msgGen("service2")
	for {
		fmt.Println(<-m1)
		fmt.Println(<-m2)
	}
}
