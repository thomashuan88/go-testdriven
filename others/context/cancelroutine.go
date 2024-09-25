package context

import (
	"context"
	"time"
)

func GetIp(ctx context.Context) (ip string, err error) {
	time.Sleep(4 * time.Second)

	ip = "127.0.0.1"
	return
}
