package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

var sum int64 = 0

func main() {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("file could not be open")
	}
	input := file

	praseList(input)

}

func convertAddSum(input byte) {
	i, err := strconv.Atoi(string(input))
	if err != nil {
		panic(err)
	}
	sum += int64(i)
}

// this function parses the list and checks if the current symbol + the next one is the same and if so adds to the sum variable
func praseList(input []byte) int64 {

	arrayLength := len(input)

	for index, value := range input {

		// if index is at max length we need to compare current and first index of list
		if index == (arrayLength - 1) {
			if value == input[0] {
				convertAddSum(value)
			}
		} else if value == input[index+1] {
			convertAddSum(value)
		}
	}

	return sum
}
