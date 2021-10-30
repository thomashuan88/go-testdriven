package cancel_test

import (
	"fmt"
	"testing"
	"time"
)

func isCancelled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

// func cancel_1(cancelChan chan struct{}) {
// 	cancelChan <- struct{}{}
// }

func cancel_2(cancelChan chan struct{}) {
	close(cancelChan) // use close() to broadcast to every goroutine
}

func TestCancel(t *testing.T) {
	cancelChan := make(chan struct{})
	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCancelled(cancelChan) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Canccelled")
		}(i, cancelChan)
	}
	cancel_2(cancelChan)
	time.Sleep(time.Second * 1)
}
