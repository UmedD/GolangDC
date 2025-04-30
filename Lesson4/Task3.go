package main

import "fmt"

func main() {
	fruit := map[string]int{
		"banana": 20,
		"apple":  25,
		"qiwi":   30,
		"grape":  15,
	}

	fmt.Println(change(fruit))
}

func change(product map[string]int) map[int]string {
	fruit := make(map[int]string)
	for key, value := range product {
		fruit[value] = key
	}

	return fruit
}
