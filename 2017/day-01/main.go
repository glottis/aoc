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
	//fmt.Println(file)

	arrayLength := len(file)

	for index, value := range file {

		// if we have reached max length of the input we need to make a rebound check to the start of the array if the first element is the same as the last
		if index == (arrayLength - 1) {
			if value == file[0] {
				addSum(value)
			}
		} else if value == file[index+1] {
			addSum(value)
		}
	}

	fmt.Println(sum)

}

func addSum(input byte) {
	i, err := strconv.Atoi(string(input))
	if err != nil {
		panic(err)
	}
	sum += int64(i)
}
