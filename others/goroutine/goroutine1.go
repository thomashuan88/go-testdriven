package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				a[i]++
				runtime.Gosched() // give away process time
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
