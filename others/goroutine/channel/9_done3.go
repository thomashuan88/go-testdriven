package main

import (
	"fmt"
)

func doWork(id int, c chan int, done chan bool) {
	// for {
	// 	d, more := <-c
	// 	if !more {
	// 		return
	// 	}
	// 	fmt.Printf("worker %d received %c\n", id, d)
	// }
	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
		// go func() { done <- true }()
		done <- true
	}
}

type worker struct {
	in   chan int
	done chan bool
}

func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWork(id, w.in, w.done)
	return w
}

func chanDemo() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	// for i := 0; i < 10; i++ {
	for i, worker := range workers {
		worker.in <- 'a' + i
		// <-workers[i].done
	}
	for _, worker := range workers {
		<-worker.done
	}
	// for i := 0; i < 10; i++ {
	for i, worker := range workers {
		worker.in <- 'A' + i
		// <-workers[i].done
	}

	// wait for all of them
	for _, worker := range workers {
		<-worker.done
	}
}

func main() {
	chanDemo()

}
