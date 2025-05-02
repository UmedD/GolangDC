package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}

	for i, el := range slice {
		slice[i] = el + 1
	}
	fmt.Println(slice)
}
