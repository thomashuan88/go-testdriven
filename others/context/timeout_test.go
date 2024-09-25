package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGetIp3(t *testing.T) {

	t1 := time.Now()

	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)

	wgg.Add(1)
	go func() {
		ip, err := GetIp3(ctx, &wg)
		fmt.Println(ip, err)
	}()

	wgg.Wait()
	fmt.Println("finished - ", time.Since(t1))
}
