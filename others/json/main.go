package main

import "fmt"

type Order struct {
	ID         string
	Name       string
	Quantity   int
	TotalPrice float64
}

func main() {
	o := Order{
		ID:         "1234",
		Name:       "learn go",
		Quantity:   3,
		TotalPrice: 30,
	}

	fmt.Printf("%+v\n", o)
}
