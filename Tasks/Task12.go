package main

import "fmt"

// Отсортировать слайс по возрастанию (вручную, без sort).
func main() {
	slice := []int{2, 3, 1, 5, 4, 7}
	tmp := slice[0]
	for i := 0; i < len(slice); i++ {
		for j := 0; j < i; j++ {
			if slice[i] < slice[j] {
				tmp = slice[j]
				slice[j] = slice[i]
				slice[i] = tmp
			}
		}
	}
	fmt.Println(slice)
}
