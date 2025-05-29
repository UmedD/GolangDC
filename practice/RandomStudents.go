package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Student struct {
	ID   int
	Name string
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println(" СТУДЕНТЫ ===")
	students := []Student{
		{ID: 1, Name: "Одинаев Анушервон"},
		{ID: 2, Name: "Давлатов Умед"},
		{ID: 3, Name: "Рохила Иматова"},
		{ID: 4, Name: "Мардонов Фахриддин"},
		{ID: 5, Name: "Джураев Бехруз"},
		{ID: 6, Name: "Юсунов Джахонгир"},
		{ID: 7, Name: "Бозоров Анушервон"},
		{ID: 8, Name: "Гадоев Ахлиддин"},
		{ID: 9, Name: "Сатторов Хадис"},
		{ID: 10, Name: "Масрур Рахимов"},
		{ID: 11, Name: "Мухаммадджони Неъматулло"},
		{ID: 12, Name: "Халифхода Абубакр"},
		{ID: 13, Name: "Каримзода Фируз"},
		{ID: 14, Name: "Чавлиев Шералихон"},
		{ID: 15, Name: "Салимзода Умед"},
		{ID: 16, Name: "Мансуруллозода Минхоч"},
		{ID: 17, Name: "Шамсов Шохин"},
		{ID: 18, Name: "Умаров Шахзод"},
		{ID: 19, Name: "Хакимов Маъруф"},
		{ID: 20, Name: "Эркабоев Давлат"},
		{ID: 21, Name: "Рамадонов Некруз"},
		{ID: 22, Name: "Мусоев Хусрав"},
	}

	shuffleStudents(students)
	fmt.Println("Случайный порядок студентов:")
	for i, student := range students {
		fmt.Printf("%d. %s\n", i+1, student.Name)
	}

	randomStudent := students[rand.Intn(len(students))]
	fmt.Printf("Случайно выбранный студент: %s", randomStudent.Name)
}

// Функция для перемешивания слайса студентов (алгоритм Фишера-Йетса)
func shuffleStudents(students []Student) {
	for i := len(students) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		students[i], students[j] = students[j], students[i]
	}
}
