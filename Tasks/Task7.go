package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	newSlice := make([]int, len(slice))
	for i := 0; i < len(slice); i++ {
		newSlice[i] = slice[len(slice)-1-i]
	}
	fmt.Println(newSlice)
}
