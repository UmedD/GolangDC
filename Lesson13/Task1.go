package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world from HelloHandler")
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Гость"
	}
	fmt.Fprintf(w, "Hello, %s!\n", name)
}

func sumHandler(w http.ResponseWriter, r *http.Request) {
	aStr := r.URL.Query().Get("a")
	bStr := r.URL.Query().Get("b")
	a, err1 := strconv.Atoi(aStr)
	b, err2 := strconv.Atoi(bStr)

	if err1 != nil || err2 != nil {
		http.Error(w, "Неверный параметры ", http.StatusBadRequest)
	}
	sum := a + b
	fmt.Fprintf(w, "sum, %d ", sum)
}
func jsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "успех"})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/greet", greetHandler)
	mux.HandleFunc("/sum", sumHandler)
	mux.HandleFunc("/json", jsonHandler)
	http.ListenAndServe(":8080", mux)
}
