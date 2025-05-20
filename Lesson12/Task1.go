package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/hello", helloHandler) // Регистрируем обработчик маршрута

	fmt.Println("Сервер запущен на http://localhost:8080")
	err := http.ListenAndServe(":8080", nil) // Запускаем сервер на порту 8080
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
