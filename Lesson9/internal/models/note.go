package models

import "time"

type Note struct {
	ID        int
	Title     string
	Content   string
	Done      bool
	CreatedAt time.Time
}
