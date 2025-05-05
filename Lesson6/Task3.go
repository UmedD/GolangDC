package main

import "fmt"

type City struct {
	name       string
	population int
	isCapital  bool
}

func main() {
	var count int
	var sum int

	city1 := City{
		name:       "Dushanbe",
		population: 1000000,
		isCapital:  true,
	}

	city2 := City{
		name:       "Tashkent",
		population: 3000000,
		isCapital:  true,
	}

	city3 := City{
		name:       "Seoul",
		population: 8000000,
		isCapital:  true,
	}

	city4 := City{
		name:       "Hujand",
		population: 500000,
		isCapital:  false,
	}

	cities := []City{
		city1, city2, city3, city4,
	}

	for i := 0; i < len(cities); i++ {
		if cities[i].isCapital {
			count++
			sum += cities[i].population
		}
	}
	fmt.Printf("Количество столиц : %v Сумма насиление городов: %d\n", count, sum)
}
