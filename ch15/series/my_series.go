package series

import "fmt"

func init() {
	fmt.Println("init1")
}
func init() {
	fmt.Println("init2")
}

func GetFibonacci(n int) []int {
	fibList := []int{1, 1}

	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList
}

// ** lowercase cannot accessed by difference package
// func square(n int) int {
// 	return n * n
// }

func Square(n int) int {
	return n * n
}
