package channel

import (
	"fmt"
	"testing"
	"time"
)

func printCount(c chan int) {
	num := 0
	for num >= 0 {
		num = <-c
		fmt.Print(num, " ")
	}
}

func TestChanCount(t *testing.T) {
	c := make(chan int)
	a := []int{8, 6, 7, 5, 3, 0, 9, -1}

	go printCount(c)

	for _, v := range a {
		c <- v
	}
	time.Sleep(time.Millisecond * 1)
	t.Log("End of test")
}
