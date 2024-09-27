package context

import (
	"context"
	"fmt"
	"time"
)

func CcA() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "desc", "CcA")
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	data := [][]int{{1, 2}, {3, 2}}
	ch := make(chan []int)
	go calculate(ctx, ch)

	for i := 0; i < len(data); i++ {
		ch <- data[i]
	}
	time.Sleep(10 * time.Second)
}

func calculate(ctx context.Context, data <-chan []int) {

	desc := ctx.Value("desc").(string)
	for {
		select {
		case <-ctx.Done():
			// desc := ctx.Value("desc").(string)
			fmt.Printf("calculate quit, context desc : %s, err : %v\n", desc, ctx.Err())
			return
		case item := <-data:
			ctx = context.WithValue(ctx, "desc", "calculate")

			ch := make(chan []int)
			go sumContext(ctx, ch)
			ch <- item

			ch1 := make(chan []int)
			go multiContext(ctx, ch1)
			ch1 <- item

		}
	}
}

func sumContext(ctx context.Context, data <-chan []int) {
	for {
		select {
		case <-ctx.Done():
			desc := ctx.Value("desc").(string)
			fmt.Printf("sumContext quit, context desc : %s, err : %v\n", desc, ctx.Err())
			return
		case item := <-data:
			a, b := item[0], item[1]
			res := sum(a, b)
			fmt.Printf("%d + %d = %d\n", a, b, res)
		}
	}
}

func multiContext(ctx context.Context, data <-chan []int) {
	for {
		select {
		case <-ctx.Done():
			desc := ctx.Value("desc").(string)
			fmt.Printf("multiContext quit, context desc : %s, err : %v\n", desc, ctx.Err())
			return
		case item := <-data:
			a, b := item[0], item[1]
			res := multi(a, b)
			fmt.Printf("%d * %d = %d\n", a, b, res)
		}
	}
}

func sum(a, b int) int {
	return a + b
}

func multi(a, b int) int {
	time.Sleep(5 * time.Second)
	return a * b
}
