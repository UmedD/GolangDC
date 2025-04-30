package main

import "fmt"

func main() {
	word := make(map[string]string)
	word["cat"] = "кошка"
	word["dog"] = "собака"
	word["fox"] = "лиса"
	word["frog"] = "лягушка"
	word["cow"] = "корова"

	search(word, "cow")
}

func search(word map[string]string, key string) {
	if value, ok := word[key]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("Слово не найдено")
	}
}
