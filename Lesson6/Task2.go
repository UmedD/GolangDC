package main

import (
	"fmt"
)

func printExellentStudents(students []Student) {
	for i := 0; i < len(students); i++ {
		if students[i].GPA > 3.5 {
			fmt.Printf("Name: %v, Age: %v, Group: %v, GPA: %v\n",
				students[i].name, students[i].age, students[i].Group, students[i].GPA)
		}
	}
}
