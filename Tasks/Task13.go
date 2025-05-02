package main

import (
	"fmt"
	"reflect"
)

// Сравнить два слайса на равенство.
func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{1, 2, 3, 4, 4}
	fmt.Println(reflect.DeepEqual(slice1, slice2))
	fmt.Println(equal(slice1, slice2))
}

func equal(slice1 []int, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
