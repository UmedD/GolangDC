package main

import "fmt"

func main() {
	result := f(5, func(n int) int {
		return n + 1
	})
	fmt.Println(result)
}

func f(num int, f2 func(n int) int) int {
	return f2(num)
}
