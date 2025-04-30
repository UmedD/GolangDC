package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var el int
	flag := false
	fmt.Scan(&el)
	for _, element := range slice {
		if el == element {
			flag = true
			break
		}
	}

	fmt.Println(flag)
}
