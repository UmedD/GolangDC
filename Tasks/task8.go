package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 2, 5, 2, 7}
	visited := make(map[int]bool)
	newSlice := []int{}

	for _, el := range slice {
		if !visited[el] {
			visited[el] = true
			newSlice = append(newSlice, el)
		}
	}

	fmt.Println("Оригинал:", slice)
	fmt.Println("Без дубликатов:", newSlice)
}
