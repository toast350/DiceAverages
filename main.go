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

	var one, two, three, four, five, six, seven, eight, nine, ten, eleven, twelve, thirteen, fourteen, fifteen, sixteen, seventeen, eighteen, nineteen, twenty int = 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0

	for i := 0; i < len(theSliceWithTheNumbers); i++ {
		switch theSliceWithTheNumbers[i] {
		case 1:
			one += 1
		case 2:
			two += 1
		case 3:
			three += 1
		case 4:
			four += 1
		case 5:
			five += 1
		case 6:
			six += 1
		case 7:
			seven += 1
		case 8:
			eight += 1
		case 9:
			nine += 1
		case 10:
			ten += 1
		case 11:
			eleven += 1
		case 12:
			twelve += 1
		case 13:
			thirteen += 1
		case 14:
			fourteen += 1
		case 15:
			fifteen += 1
		case 16:
			sixteen += 1
		case 17:
			seventeen += 1
		case 18:
			eighteen += 1
		case 19:
			nineteen += 1
		case 20:
			twenty += 1
		}
	}

	// Divide the thingies by 4 later
	fmt.Printf("Percentages:\n1: %v\n2: %v\n3: %v\n4: %v\n5:% v\n6: %v\n7: %v\n8: %v\n9: %v\n10: %v\n11: %v\n12: %v\n13: %v\n14: %v\n15: %v\n16: %v\n17: %v\n18: %v\n19: %v\n20: %v\n", one, two, three, four, five, six, seven, eight, nine, ten, eleven, twelve, thirteen, fourteen, fifteen, sixteen, seventeen, eighteen, nineteen, twenty)
}
