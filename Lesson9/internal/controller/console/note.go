package console

import (
	"GolangDC/Lesson9/internal/models"
	"GolangDC/Lesson9/internal/service"
	"fmt"
)

func Create() {
	var title, content string
	var id int
	fmt.Println("Введите ID")
	fmt.Scan(&id)
	fmt.Println("Введите заголовок: ")
	fmt.Scan(&title)
	fmt.Println("Введите содержание: ")
	fmt.Scan(&content)

	note := models.Note{
		ID:      id,
		Title:   title,
		Content: content,
		Done:    false,
	}

	createdNote := service.Create(note)
	fmt.Println("заметка создана: ", createdNote)
}

func GetAll() {
	notes, err := service.GetAll()
	if err != nil {
		fmt.Println("не удалось получить заметки, причина: ", err.Error())
		return
	}

	for _, note := range notes {
		fmt.Printf("ID: %d, Title: %s\n", note.ID, note.Title)
	}
}

func GetByID() {
	fmt.Println("Введите ID ")
	var noteId int
	fmt.Scan(&noteId)

	if noteId <= 0 {
		fmt.Println("ID invalid")
		return
	}

	note, err := service.GetByID(noteId)
	if err != nil {
		fmt.Println("не удалось получить заметку ", err.Error())
		return
	}

	fmt.Printf("ID: %d, Title: %s\n", note.ID, note.Title)
	fmt.Printf("CreatedAt: %s, UserID: %s\n", note.CreatedAt.Format("2006/1/2, 15:04 \n"), note.Content)

}

func DeleteById() {
	var id int
	fmt.Println("Введите ID заметки для удаления: ")
	fmt.Scan(&id)

	success := service.DeleteById(id)
	if success {
		fmt.Println("Заметка удалена")
	} else {
		fmt.Println("Заметка не найдена")
	}
}
