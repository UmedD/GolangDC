package repository

import (
	errors2 "GolangDC/Lesson9/internal/errors"
	"GolangDC/Lesson9/internal/models"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func loadData() ([]models.Note, error) {
	file, err := os.Open("storage.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var accounts []models.Note
	err = json.NewDecoder(file).Decode(&accounts)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func saveData(accounts []models.Note) error {
	file, err := os.OpenFile("storage.json", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	return encoder.Encode(accounts)
}

func GetAllAccounts() ([]models.Note, error) {
	accounts, err := loadData()
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func GetAccountByID(accountID int) (models.Note, error) {
	accounts, err := loadData()
	if err != nil {
		return models.Note{}, err
	}

	for _, account := range accounts {
		if account.ID == accountID {
			return account, nil
		}
	}
	return models.Note{}, errors2.ErrAccountNotFound
}

func DeleteByID(accountID int) (bool, error) {
	accounts, err := loadData()
	if err != nil {
		return false, err
	}
	var indexToDelete = -1

	for i, account := range accounts {
		if account.ID == accountID {
			indexToDelete = i
			break
		}
	}

	if indexToDelete == -1 {
		return false, fmt.Errorf("account with ID %d not found", accountID)
	}

	accounts = append(accounts[:indexToDelete], accounts[indexToDelete+1:]...)
	data, err := json.MarshalIndent(accounts, "", "  ")
	if err != nil {
		return false, err
	}

	err = os.WriteFile("storage.json", data, 0644)
	if err != nil {
		return false, err
	}

	return true, nil
}

func CreateNote(note models.Note) (models.Note, error) {
	accounts, err := loadData()
	if err != nil {
		return models.Note{}, err
	}

	maxID := 0
	for _, acc := range accounts {
		if acc.ID > maxID {
			maxID = acc.ID
		}
	}
	note.ID = maxID + 1
	note.CreatedAt = time.Now()

	accounts = append(accounts, note)

	data, err := json.MarshalIndent(accounts, "", "  ")
	if err != nil {
		return models.Note{}, err
	}

	err = os.WriteFile("storage.json", data, 0644)
	if err != nil {
		return models.Note{}, err
	}

	return note, nil
}
