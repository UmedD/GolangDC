package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0
	for _, el := range slice {
		sum += el
	}
	fmt.Println(sum)
}
