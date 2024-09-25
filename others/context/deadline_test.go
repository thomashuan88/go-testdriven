package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGetIp2(t *testing.T) {

	t1 := time.Now()

	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))

	wg.Add(1)
	go func() {
		ip, err := GetIp2(ctx, &wg)
		fmt.Println(ip, err)
	}()

	wg.Wait()
	fmt.Println("finished - ", time.Since(t1))
}
