package day01

import (
	"math"
)

// ParseList : this function parses the list and checks if the current symbol + the next one is the same and if so adds to the sum variable
func ParseList(input []byte) int {

	var sum int

	arrayLength := len(input)

	for index, value := range input {
		// if index is at max length we need to compare the current (e.g. last value) and first value of list
		if value-0x30 >= 0 && value-0x30 <= 9 {

			if index == (arrayLength - 1) {
				if value == input[0] {
					sum += int(value - 0x30)
				}
			} else if value == input[index+1] {
				sum += int(value - 0x30)
			}
		}

	}

	return sum
}

// ParseListRedux : this function parses the list and checks if the current symbol + (halfway around one) is the same and if so adds to the sum variable
func ParseListRedux(input []byte) int {

	var sum int

	arrayLength := len(input)
	lengthDiv := arrayLength / 2

	for index, value := range input {

		if value-0x30 >= 0 && value-0x30 <= 9 {

			if (index + lengthDiv) >= arrayLength {

				corrIndex := math.Abs(float64(arrayLength - (index + lengthDiv)))
				if value == input[int(corrIndex)] {

					sum += int(value - 0x30)
				}
			} else if value == input[index+lengthDiv] {
				sum += int(value - 0x30)
			}
		}

	}

	return sum
}
