package channel

import (
	"fmt"
	"sync"
)

func doWork10(id int, c chan int, wg *sync.WaitGroup) {
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

type worker10 struct {
	in chan int
	wg *sync.WaitGroup
}

func createWorker10(id int, wg *sync.WaitGroup) worker10 {
	w := worker10{
		in: make(chan int),
		wg: wg,
	}
	go doWork10(id, w.in, wg)
	return w
}

func chanDemo10() {
	var workers [10]worker10
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		workers[i] = createWorker10(i, &wg)
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
