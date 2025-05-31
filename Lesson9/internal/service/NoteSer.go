package service

import (
	"GolangDC/Lesson9/internal/models"
	"GolangDC/Lesson9/internal/repository"
	"fmt"
)

func Create(note models.Note) models.Note {
	if note.Title == "" {
		fmt.Println("Ошибка: Заголовок не может быть пустым.")
		return note
	}
	return repository.Create(note)
}

func GetAll() ([]models.Note, error) {
	accounts, err := repository.GetAll()
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func GetByID(id int) (models.Note, error) {
	note, err := repository.GetByID(id)
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
	return repository.DeleteById(id)
}

func UpdateAccount(note models.Note) error {
	return repository.UpdateByID(note.ID, note)
}
