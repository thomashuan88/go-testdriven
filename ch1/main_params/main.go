package main

import (
	"fmt"
	"os"
)

// the way to receipt params when runtime

func main() {
	fmt.Println(os.Args)
	if len(os.Args) > 1 {
		fmt.Println("helle world", os.Args[1])
	}
	fmt.Println("hello world")
	os.Exit(-1)
}
