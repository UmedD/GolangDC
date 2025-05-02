package main

import "fmt"

// Найти среднее значение всех элементов слайса.
func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 7}
	average := 0
	count := 0
	for _, el := range slice {
		average += el
		count++
	}
	fmt.Println(float64(average) / float64(count))
}
