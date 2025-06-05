package main

import (
	"GolangDC/Lesson9/internal/controller/gin_rest"
	"fmt"
	"os"
)

func main() {
	wd, _ := os.Getwd()
	fmt.Println("Current dir:", wd)

	gin_rest.StartServer()
	//http_rest.StartServer()
	//console.RunCommands()
}
