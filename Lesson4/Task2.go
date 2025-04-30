package main

import "fmt"

type productMap map[string]int

func main() {
	product := make(productMap)
	product["phone"] = 500
	product["computer"] = 1000
	product["buds 2"] = 400
	product["keyboard"] = 200

	add(product)
	output(product)
	fmt.Println(price(product, "computer"))
}

func add(product productMap) productMap {
	product["mouse"] = 340
	product["hard disk"] = 440
	return product
}

func price(product productMap, name string) int {
	return product[name]
}

func output(product productMap) {
	for key, value := range product {
		fmt.Printf("product : %s | price : %d TJS\n", key, value)
	}
}
