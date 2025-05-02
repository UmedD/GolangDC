package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	result := mapInt(slice, func(n int) int {
		return n * 2
	})
	fmt.Println(result)
}

func mapInt(slice []int, f func(n int) int) []int {
	newSlice := make([]int, len(slice))
	for i, el := range slice {
		newSlice[i] = f(el)
	}
	return newSlice
}
