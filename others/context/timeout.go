package context

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wgg = sync.WaitGroup{}

func GetIp3(ctx context.Context, wg *sync.WaitGroup) (ip string, err error) {
	go func() {
		<-ctx.Done()
		fmt.Println("routine expired ", ctx.Err())
		err = ctx.Err()
		wgg.Done()
	}()

	time.Sleep(4 * time.Second)

	ip = "127.0.0.1"
	wgg.Done()
	return
}
