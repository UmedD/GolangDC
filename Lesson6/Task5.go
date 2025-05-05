package main

import "fmt"

type Order struct {
	Item         string
	Quantity     int
	PricePerUnit float64
}

func main() {
	order1 := Order{
		Item:         "PC",
		Quantity:     12,
		PricePerUnit: 3,
	}

	order2 := Order{
		Item:         "Keyboard",
		Quantity:     10,
		PricePerUnit: 2,
	}

	order3 := Order{
		Item:         "Mouse",
		Quantity:     9,
		PricePerUnit: 1,
	}

	orders := [3]Order{order1, order2, order3}
	var sum float64
	for i := 0; i < len(orders); i++ {
		sum += float64(orders[i].Quantity) * orders[i].PricePerUnit
	}

	fmt.Printf("summ all orders: %v", sum)
}
