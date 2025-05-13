package controller

import "fmt"

var Commands = `
Выберите нужную команду:
0. Выход
1. Показать все заметки
2. Получить заметку по ID
3. Добавить заметку
4. Удалить заметку по ID`

func RunCommands() {
	for {
		fmt.Println(Commands)
		var choice string
		fmt.Scan(&choice)
		switch choice {
		case "0":
			fmt.Println("До свидания!")
			return
		case "1":
			GetAll()
		case "2":
			GetByID()
		case "3":
			Create()
		case "4":
			DeleteById()
		default:
			fmt.Println("Неверная команда. Попробуйте ещё раз.")
		}
	}
}
