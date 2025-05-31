package repository

import (
	"GolangDC/Lesson9/internal/errors"
	"GolangDC/Lesson9/internal/models"
	"fmt"
)

func Create(note models.Note) models.Note {
	noteList = append(noteList, note)
	return note
}

func GetAll() ([]models.Note, error) {
	if noteList == nil {
		return nil, errors.ErrSomethingWentWrong
	}
	return noteList, nil
}

func GetByID(id int) (models.Note, error) {
	for _, note := range noteList {
		if note.ID == id {
			return note, nil
		}
	}

	return models.Note{}, errors.ErrAccountNotFound
}

func DeleteById(id int) bool {
	for i, note := range noteList {
		if note.ID == id {
			noteList = append(noteList[:i], noteList[i+1:]...)
			return true
		}
	}
	return false
}

func UpdateByID(id int, updated models.Note) error {
	for i, note := range noteList {
		if note.ID == id {
			noteList[i] = updated
			return nil
		}
	}
	return fmt.Errorf("note with ID %d not found", id)
}

func GetMaxID() int {
	maxID := 0
	for _, note := range noteList {
		if note.ID > maxID {
			maxID = note.ID
		}
	}
	return maxID
}
