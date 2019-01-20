package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func traverseInstruction(intS []int) int {
	steps := 0
	length := len(intS)

	for i := 0; ; {
		current := intS[i]
		intS[i]++
		i = current + i
		steps++
		if i >= length {
			break
		}
	}

	return steps

}

func traverseInstructionRedux(intS []int) int {
	steps := 0
	length := len(intS)

	for i := 0; ; {
		current := intS[i]
		if intS[i] >= 3 {
			intS[i]--
		} else {
			intS[i]++
		}
		i = current + i
		steps++
		if i >= length {
			break
		}
	}

	return steps

}

func main() {
	input, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	inputSlice := strings.Fields(string(input))
	intSlice := make([]int, len(inputSlice))

	for i, v := range inputSlice {
		vv, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		intSlice[i] = vv
	}
	fmt.Println("the number of steps is: ", traverseInstructionRedux(intSlice))

}
