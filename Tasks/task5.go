package main

import "fmt"

// Объединить два слайса в один.
func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{6, 7, 8, 9, 10}
	result := make([]int, 0, len(slice1)+len(slice2))
	result = append(result, slice1...)
	result = append(result, slice2...)
	fmt.Println(result)
}
