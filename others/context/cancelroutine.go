package context

import (
	"context"
	"fmt"
	"time"
)

func GetIp(ctx context.Context) (ip string, err error) {
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("routine canceled ", ctx.Err())
			err = ctx.Err()
			return
		}
	}()

	time.Sleep(4 * time.Second)

	ip = "127.0.0.1"
	return
}
