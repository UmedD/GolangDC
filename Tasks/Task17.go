package main

import "fmt"

type Student struct {
	Name   string
	Grades map[string]int
}

func AddGrade(student *Student, subject string, grade int) {
	if student.Grades == nil {
		student.Grades = make(map[string]int)
	}
	student.Grades[subject] = grade
}

func AverageGrade(student Student) float64 {
	var sum float64
	for _, grade := range student.Grades {
		sum += float64(grade)
	}
	return float64(sum) / float64(len(student.Grades))
}

func BestSubject(student Student) string {
	for subject, grade := range student.Grades {
		if grade == 55 {
			return subject
		}
	}
	return "Плохие оценки"
}

func filterTopStudents(students []Student, minAverage float64) []Student {
	slice := []Student{}
	for _, student := range students {
		if AverageGrade(student) > minAverage {
			slice = append(slice, student)
		}
	}
	return slice
}

func main() {
	student1 := Student{Name: "Алиса"}
	AddGrade(&student1, "Математика", 90)
	AddGrade(&student1, "Физика", 85)
	AddGrade(&student1, "Информатика", 95)

	student2 := Student{Name: "Боб"}
	AddGrade(&student2, "Математика", 60)
	AddGrade(&student2, "Физика", 70)

	fmt.Println(BestSubject(student1))
	fmt.Println(BestSubject(student2))

	students := []Student{student1, student2}

	for _, s := range students {
		fmt.Printf("Студент: %s\n", s.Name)
		fmt.Printf("Средняя оценка: %.2f\n", AverageGrade(s))
		fmt.Printf("Лучший предмет: %s\n", BestSubject(s))
		fmt.Println("------")
	}

}
