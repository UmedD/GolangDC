package service

import (
	"GolangDC/Lesson9/internal/models"
	"GolangDC/Lesson9/internal/repository"
	"GolangDC/Lesson9/internal/repository/MemoryRepo"
	"fmt"
)

func Create(note models.Note) models.Note {
	if note.Title == "" {
		fmt.Println("Ошибка: Заголовок не может быть пустым.")
		return note
	}
	ok, err := repository.CreateNote(note)
	if err != nil {
		fmt.Println("Ошибка при удалении:", err)
		return note
	}

	return ok
}

func GetAll() ([]models.Note, error) {
	accounts, err := repository.GetAllAccounts()
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func GetByID(id int) (models.Note, error) {
	note, err := repository.GetAccountByID(id)
	if err != nil {
		return models.Note{}, err
	}

	return note, nil
}

func DeleteById(id int) bool {
	if id <= 0 {
		fmt.Println("Некорректный ID")
		return false
	}

	ok, err := repository.DeleteByID(id)
	if err != nil {
		fmt.Println("Ошибка при удалении:", err)
		return false
	}

	return ok
}

func UpdateAccount(note models.Note) error {
	return MemoryRepo.UpdateByID(note.ID, note)
}
