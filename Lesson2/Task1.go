package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d * %d = %d\n", i, num, i*num)
	}
}
