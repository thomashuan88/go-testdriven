package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGetIp(t *testing.T) {
	// t.Run("success", func(t *testing.T) {
	// 	ctx := context.Background()
	// 	ip, err := GetIp(ctx)
	// 	if err != nil {
	// 		t.Errorf("expected no error, got %v", err)
	// 	}
	// 	if ip != "127.0.0.1" {
	// 		t.Errorf("expected ip to be 127.0.0.1, got %s", ip)
	// 	}
	// })

	// t.Run("timeout", func(t *testing.T) {
	// 	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	// 	defer cancel()
	// 	ip, err := GetIp(ctx)
	// 	if err != context.DeadlineExceeded {
	// 		t.Errorf("expected DeadlineExceeded error, got %v", err)
	// 	}
	// 	if ip != "" {
	// 		t.Errorf("expected ip to be empty, got %s", ip)
	// 	}
	// })

	t1 := time.Now()

	ctx, cancel := context.WithCancel(context.Background())

	wait.Add(1)
	go func() {
		ip, err := GetIp(ctx)
		fmt.Println(ip, err)
	}()

	go func() {
		time.Sleep(2 * time.Second)

		// cancel routine
		cancel()
	}()

	wait.Wait()
	fmt.Println("finished - ", time.Since(t1))
}
