package main

import (
	"fmt"
	"sync"
)

func doWork(id int, c chan int, wg *sync.WaitGroup) {
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
		wg.Done()
	}
}

type worker struct {
	in chan int
	wg *sync.WaitGroup
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		wg: wg,
	}
	go doWork(id, w.in, wg)
	return w
}

func chanDemo() {
	var workers [10]worker
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(20)
	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}
	wg.Wait()

}

func main() {
	chanDemo()

}
