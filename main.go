package main

import (
	"fmt"
	"math"
	"math/rand"
)

var testAmount int = 10000000

func main() {
	var diceSides int

	fmt.Println("How many sides on the dice?")
	fmt.Scan(&diceSides)

	total := 0
	for i := 0; i < testAmount; i++ {
		total += int(math.Max(float64(rand.Intn(diceSides))+1, float64(rand.Intn(diceSides))))
	}
}
