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

// use hasValue to detect have send value
func main() {
	var c1, c2 = generator(), generator()
	w := createWorker(0)
	n := 0
	hasValue := false
	for {
		select {
		case n = <-c1:
			hasValue = true
		case n = <-c2:
			hasValue = true
		case w <- n:
			if hasValue {
				// send to w, but code cannot be like this
			}
		}
	}

}
