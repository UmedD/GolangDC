package main

import "fmt"

func main() {
	var arr = []int{1, 2, 3, 4, 5}
	var idx int
	fmt.Scan(&idx)
	arr = append(arr[:idx], arr[idx+1:]...)
	fmt.Println(arr)
}
