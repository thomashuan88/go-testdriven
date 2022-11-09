package main

import (
	"encoding/json"
	"fmt"
)

type Order struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}

func main() {
	o := Order{
		ID:         "1234",
		Name:       "learn go",
		Quantity:   3,
		TotalPrice: 30,
	}

	b, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)
}
