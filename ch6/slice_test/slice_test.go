package slice_test

import (
	"fmt"
	"testing"
)

func TestSliceInit(t *testing.T) {
	var s0 []int // slice is difference with array
	t.Log(len(s0), cap(s0))

	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	// t.Log(s2[0], s2[1], s2[2], s2[3], s2[4]) hit error len just 3, then only can visit 3 value
	t.Log(s2[0], s2[1], s2[2])
	s2 = append(s2, 1)
	t.Log(s2[0], s2[1], s2[2], s2[3])
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s)) // capacity every time growing time 2
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "Unknow"
	t.Log(Q2)
	t.Log(year)
}

func TestSliceComparing(t *testing.T) {
	// a := []int{1, 2, 3, 4}
	// b := []int{1, 2, 3, 4}
	// if a == b {
	// 	t.Log("equal")
	// }
}

func TestAppendSlice(t *testing.T) {
	arr := make([]int, 0)
	for i := 0; i < 20; i++ {
		if len(arr) > 0 {
			fmt.Printf("len : %d, cap : %d, location : %p\n", len(arr), cap(arr), &arr[0])
		} else {
			fmt.Printf("len : %d, cap : %d, location : (empty)\n", len(arr), cap(arr))
		}
		arr = append(arr, i)
	}

}

// a func pass in channel and print value

// test print a string with two goroutines
func TestPrintString(t *testing.T) {
	// a int channel
	// ch := make(chan int)

	// create a array random 30 numbers
	arr := make([]int, 30)
	for i := 0; i < 30; i++ {
		arr[i] = i
	}
	fmt.Println(arr)

	// go printChar(ch)
	// go printChar(ch)

	// for
}

func swap(sw []int) {
	for a, b := 0, len(sw)-1; a < b; a, b = a+1, b-1 {
		fmt.Println(a, b)
		sw[a], sw[b] = sw[b], sw[a]
	}
}

func TestSwap(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	swap(a)
	t.Log(a)
}
