package share_mem_test

import (
	"sync"
	"testing"
	"time"
)

// unsafe multi thread
func TestCounter_unsafe(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

func TestCounter_safe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

func TestWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Logf("counter = %d", counter)
}
