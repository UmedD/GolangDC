package rest

import (
	"log"
	"net/http"
)

func StartServer() {
	mux := http.NewServeMux()

	RegisterRoutes(mux)

	log.Println("🚀 Сервер запущен на http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Сервер не запустился: %v", err)
	}

	http.HandleFunc("acc", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			GetAllAccount(writer, request)
		default:
			http.Error(writer, "Method not alowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/accounts/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			GetAccountByID(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			CreateAccount(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/delete/", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			DeleteAccount(writer, request)
		default:
			http.Error(writer, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
}
