package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)

	if number%3 == 0 && number%5 == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
