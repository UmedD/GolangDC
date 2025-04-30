package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := make([]int, 0, len(slice))

	for _, el := range slice {
		if el%2 == 0 {
			result = append(result, el)
		}
	}
	fmt.Println(result)
}
