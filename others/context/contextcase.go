package context

import (
	"fmt"
	"time"
)

func ContextCase() {
	// context
	// ctx := context.Background()
	// ctx = context.WithValue(ctx, "desc", "ContextCase")
	// ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	// // ctx, cancel = context.WithDeadline(ctx, time.Now())
	// defer cancel()

	done := make(chan struct{})

	go f1(done)
	go f1(done)
	time.Sleep(3 * time.Second)
	close(done)
}

func f1(done chan struct{}) {
	// for {
	// 	select {
	// 	case <-done:
	// 		fmt.Println("f1 done")
	// 		return
	// 	}
	// }

	<-done
	fmt.Println("f1 done")
}
