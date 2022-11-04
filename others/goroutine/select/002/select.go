package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	for n := range c {
		fmt.Printf("worker %d received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

// like this can send and receive in select, but still need to handle n
func main() {
	var c1, c2 = generator(), generator()
	w := createWorker(0)
	for {
		select {
		case n := <-c1:
		case n := <-c2:
		case w <- n:
		}
	}

}
