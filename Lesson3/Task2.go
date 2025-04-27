package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var newSlice = []int{}

	for i := 0; i < len(slice); i++ {
		newSlice = append(newSlice, slice[len(slice)-1-i])
	}

	fmt.Println(newSlice)
}
