package main

import (
	"fmt"
	"math"
)

func main() {
	slice := []int{234, 23, 1, 343, -5}
	mx := math.MinInt
	mn := math.MaxInt

	for _, el := range slice {
		if mx < el {
			mx = el
		}

		if mn > el {
			mn = el
		}
	}

	fmt.Printf("max : %d\n min : %d", mx, mn)
}
