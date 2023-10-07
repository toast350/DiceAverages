package main

import (
	"fmt"
	"math"
)

func main() {
	var diceSides int

	fmt.Println("How many sides on the dice?")
	fmt.Scan(&diceSides)

	var mainSlice = make([][]int, diceSides)

	for i := 0; i < diceSides; i++ {
		for ii := 1; ii <= diceSides; ii++ {
			mainSlice[i] = append(mainSlice[i], ii)
		}
	}

	theSliceWithTheNumbers := make([]int, 0, diceSides*diceSides)
	theNumber := 0
	for i := 0; i < diceSides; i++ {
		for ii := 0; ii < diceSides; ii++ {
			theNumber = int(math.Max(float64(i+1), float64(mainSlice[i][ii])))

			theSliceWithTheNumbers = append(theSliceWithTheNumbers, theNumber)
		}
	}

	counters := make([][]int, diceSides+1)

	for i := 0; i < len(theSliceWithTheNumbers); i++ {
		counters[theSliceWithTheNumbers[i]] = append(counters[theSliceWithTheNumbers[i]], theSliceWithTheNumbers[i])
	}
	counters = append(counters[1:])

	for i := 0; i < len(counters); i++ {
		fmt.Printf("%v: %v\n", i+1, len(counters[i]))
	}
}
