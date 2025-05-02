package main

import "fmt"

func main() {
	slice := []int{1, 0, 2, 0, 4, 5, 6}
	newSlice := make([]int, 0, len(slice))

	for _, el := range slice {
		if el != 0 {
			newSlice = append(newSlice, el)
		}
	}

	fmt.Println(slice)
	fmt.Println(newSlice)
}
