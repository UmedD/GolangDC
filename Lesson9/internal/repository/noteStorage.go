package repository

import (
	"GolangDC/Lesson9/internal/models"
	"time"
)

var noteList = []models.Note{
	{
		ID:        1,
		Title:     "School",
		Content:   "meeting in school with friend",
		Done:      true,
		CreatedAt: time.Now(),
	},

	{
		ID:        2,
		Title:     "Store",
		Content:   "Buy a phone",
		Done:      false,
		CreatedAt: time.Now(),
	},

	{
		ID:        3,
		Title:     "LogPass",
		Content:   "root/1234",
		Done:      true,
		CreatedAt: time.Now(),
	},
}
