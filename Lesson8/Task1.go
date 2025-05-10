package main

import "fmt"

type Storage interface {
	Save(data string)
	Load() string
}

type MemoryStorage struct {
	data string
}

func (m *MemoryStorage) Save(data string) {
	m.data = data
}

func (m MemoryStorage) Load() string {
	return m.data
}

type FileStorage struct {
	file string
}

func (f *FileStorage) Save(data string) {
	f.file = data
}

func (f FileStorage) Load() string {
	return f.file
}

func UseStroge(s Storage) {
	fmt.Println(s.Load())
}

func main() {
	memory := &MemoryStorage{data: "Memory Load"}
	storage := &FileStorage{file: "File Load"}

	UseStroge(memory)
	UseStroge(storage)
}
