package main

import (
	"fmt"
	"os"
)

// the way to return status of main function

func main() {
	a := 3
	b := 5
	a, b = b, a
	fmt.Println(a, b)
	fmt.Println("hello world")
	os.Exit(-1)
}
