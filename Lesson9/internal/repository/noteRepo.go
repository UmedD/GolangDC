package repository

import (
	"GolangDC/Lesson9/internal/errors"
	"GolangDC/Lesson9/internal/models"
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
