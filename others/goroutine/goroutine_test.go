package goroutine_test

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func count() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 1)
	}
}

func TestGoroutineCount(t *testing.T) {
	go count()
	time.Sleep(time.Millisecond * 2)
	t.Log("hello world")
	time.Sleep(time.Millisecond * 5)
}

func TestRun(t *testing.T) {
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				fmt.Printf("hello from goroutine %d\n", i)
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
}
func TestWithoutIO(t *testing.T) {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				a[i]++
				runtime.Gosched()
			}
		}(i)

	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}

func TestWithoutIO2(t *testing.T) {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) { // race condition , 数据访问冲突
			for {
				a[i]++
				runtime.Gosched()
			}
		}(i)

	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
