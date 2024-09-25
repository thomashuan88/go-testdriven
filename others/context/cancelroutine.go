package context

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wait = sync.WaitGroup{}

func GetIp(ctx context.Context) (ip string, err error) {
	go func() {
		<-ctx.Done()
		fmt.Println("routine canceled ", ctx.Err())
		err = ctx.Err()
		wait.Done()
	}()

	time.Sleep(4 * time.Second)

	ip = "127.0.0.1"
	wait.Done()
	return
}
