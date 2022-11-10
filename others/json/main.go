package main

import (
	"encoding/json"
	"fmt"
)

type OrderItem struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Order struct {
	ID         string      `json:"id"`
	Items      []OrderItem `json:"items"`
	TotalPrice float64     `json:"total_price"`
}

func unmarshal() {
	s := `{"id":"1234","items":[{"id":"item_1","name":"learn go","price":15},{"id":"item_2","name":"interview","price":10}],"total_price":30}`
	var o Order
	err := json.Unmarshal([]byte(s), &o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", o)
}

func marshal() {
	o := Order{
		ID:         "1234",
		TotalPrice: 30,
		Items: []OrderItem{
			{
				ID:    "item_1",
				Name:  "learn go",
				Price: 15,
			},
			{
				ID:    "item_2",
				Name:  "interview",
				Price: 10,
			},
		},
	}

	b, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)
}

func parseNLP() {
	res := `{
"data": [
    {
        "synonym":"",
        "weight":"0.6",
        "word": "真丝",
        "tag":"材质"
    },
    {
        "synonym":"",
        "weight":"0.8",
        "word": "韩都衣舍",
        "tag":"品牌"
    },
    {
        "synonym":"连身裙;联衣裙",
        "weight":"1.0",
        "word": "连衣裙",
        "tag":"品类"
    }
]
}`
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(res), &m)
	if err != nil {
		panic(err)
	}

	// fmt.Printf("%+v\n", m)

	// Type assertion
	fmt.Printf("%+v\n", m["data"].([]interface{})[2])
}

func main() {
	parseNLP()
}
