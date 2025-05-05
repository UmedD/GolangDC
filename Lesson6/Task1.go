package main

import "fmt"

type Student struct {
	name  string
	age   int
	Group string
	GPA   float64
}

var StudentSlice = []Student{
	{name: "Umed", age: 18, Group: "1331", GPA: 4.2},
	{name: "Rustam", age: 28, Group: "1341", GPA: 4.1},
	{name: "Ahmad", age: 24, Group: "1351", GPA: 3.5},
}

func main() {
	for i := 0; i < len(StudentSlice); i++ {
		fmt.Printf("Name: %v, Age: %v, Group: %v, GPA: %v\n",
			StudentSlice[i].name, StudentSlice[i].age, StudentSlice[i].Group, StudentSlice[i].GPA)
	}

	printExellentStudents(StudentSlice)
}
