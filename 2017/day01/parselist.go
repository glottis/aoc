package day01

import (
	"strconv"
)

func convertAddSum(input byte) int64 {
	i, err := strconv.Atoi(string(input))
	if err != nil {
		panic(err)
	}
	return int64(i)
}

// ParseList : this function parses the list and checks if the current symbol + the next one is the same and if so adds to the sum variable
func ParseList(input []byte) int64 {

	var sum int64

	arrayLength := len(input)

	for index, value := range input {

		// if index is at max length we need to compare the current (e.g. last value) and first value of list
		if index == (arrayLength - 1) {
			if value == input[0] {
				sum += convertAddSum(value)
			}
		} else if value == input[index+1] {
			sum += convertAddSum(value)
		}
	}

	return sum
}
