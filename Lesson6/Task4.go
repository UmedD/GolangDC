package main

import "fmt"

type Game struct {
	title       string
	platform    string
	hoursPlayed int
}

func main() {
	gamer1 := Game{
		title:       "Call Of Duty",
		platform:    "Steam",
		hoursPlayed: 500,
	}

	gamer2 := Game{
		title:       "Call Of Duty",
		platform:    "Steam",
		hoursPlayed: 1000,
	}

	gamer3 := Game{
		title:       "Call Of Duty",
		platform:    "Steam",
		hoursPlayed: 350,
	}

	gamerSlice := []Game{gamer1, gamer2, gamer3}
	fmt.Printf("Количество сыгранных часов: %d ", totalPlaytime(gamerSlice))
}

func totalPlaytime(game []Game) int {
	sum := 0
	for i := 0; i < len(game); i++ {
		sum += game[i].hoursPlayed
	}
	return sum
}
