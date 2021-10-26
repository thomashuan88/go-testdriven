package err_test

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

/*
func GetFibonacci(n int) ([]int, error) {
	if n < 2 || n > 100 {
		return nil, errors.New("n should be in [2,100]")
	}
	fibList := []int{1, 1}

	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}


func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, errors.New("n should be in 2")
	}
	if n > 100 {
		return nil, errors.New("n should be in 100")
	}
	fibList := []int{1, 1}

	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}
*/

// variable of Error have to use errXXXX
var errLessThanTwoError = errors.New("n should be not less than 2")
var errLargeThanHundredError = errors.New("n should be in 100")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, errLessThanTwoError
	}
	if n > 100 {
		return nil, errLargeThanHundredError
	}
	fibList := []int{1, 1}

	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func GetFibonacci2(str string) {
	var (
		i    int
		err  error
		list []int
	)
	if i, err = strconv.Atoi(str); err != nil {
		fmt.Println("Error", err)
		return
	}
	if list, err = GetFibonacci(i); err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(list)
}

func TestGetFibonacci(t *testing.T) {
	if v, err := GetFibonacci(5); err != nil {
		if err == errLessThanTwoError {
			fmt.Println("It is less .")
		}
		t.Error(err)
	} else {
		t.Log(v)
	}
}
