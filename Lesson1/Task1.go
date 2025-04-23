package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Print("Enter your name please - ")
	fmt.Scan(&name)
	fmt.Print("Enter your age please - ")
	fmt.Scan(&age)

	fmt.Printf("Your name is: %s\n ", name)
	fmt.Printf("Your age now: %d\n ", age)
	fmt.Printf("After 10 years you will be: %d ", (age + 10))
}
