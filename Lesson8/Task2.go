package main

import "fmt"

type Action interface {
	Do()
}

type JustDoIt struct {
	str string
}

func (j JustDoIt) Do() {
	fmt.Println("jus do it")
}

type Computer struct {
	str string
}

func (c Computer) Do() {
	fmt.Println("my PC")
}
func LogAction(a Action) {
	a.Do()
}
func main() {
	just := JustDoIt{}
	pc := Computer{}

	LogAction(just)
	LogAction(pc)

}
