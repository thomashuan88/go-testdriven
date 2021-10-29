package csp_test

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

// normal process othertask have to wait for othertask finish
func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

// start another goroutine to run service()
func AsyncService() chan string {
	// retCh := make(chan string)
	retCh := make(chan string, 1) // with buffer , service no need to wait
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

// for using this async process the time of running will base on the longer time's task
func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)
	time.Sleep(time.Second * 1)
}
