package rest

import (
	"GolangDC/Lesson9/internal/models"
	"GolangDC/Lesson9/internal/repository"
	"GolangDC/Lesson9/internal/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	idSeq        int
	originalNote *models.Note
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong")
	})
	mux.HandleFunc("/acc", GetAllAccount)
	mux.HandleFunc("/accounts/", GetAccountByID)
	mux.HandleFunc("/create/", CreateAccount)
	mux.HandleFunc("/delete/", DeleteAccount)
}

func GetAllAccount(w http.ResponseWriter, r *http.Request) {
	acc, err := service.GetAll()
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		handleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(acc)
}

func GetAccountByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/accounts/")
	id, err := strconv.Atoi(idStr)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Account ID is %d", id)
	account, err := service.GetByID(id)
	if err != nil {
		handleError(w, http.StatusNotFound, err.Error())
		return
	}

	json.NewEncoder(w).Encode(account)
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var note models.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	idSeq = repository.GetMaxID()
	idSeq++
	note.ID = idSeq
	note.CreatedAt = time.Now()
	service.Create(note)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(note)
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/delete/")
	id, err := strconv.Atoi(idStr)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	success := service.DeleteById(id)
	if !success {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{
		"message": fmt.Sprintf("Note with ID %d deleted", id),
		"id":      id,
		"status":  "success",
	}
	json.NewEncoder(w).Encode(response)
}
