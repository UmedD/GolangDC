package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	slice := make([]int, n)
	if n > 0 {
		slice[0] = 0
	}
	if n > 1 {
		slice[1] = 1
	}
	for i := 2; i < n; i++ {
		slice[i] = slice[i-1] + slice[i-2]
	}
	fmt.Println(slice)
}
