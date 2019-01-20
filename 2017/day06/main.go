package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func countCycles(ns []int) int {

	completed := make(map[string]bool)
	cycles := 0

	for {
		// find the index of the highest bank (if tie, take the lowest index)
		bankIndex := 0
		for i, v := range ns[1:] {
			i++
			if v > ns[bankIndex] {
				bankIndex = i
			} else if v == ns[bankIndex] {
				if i < bankIndex {
					bankIndex = i
				}
			}
		}

		//save the blocks to distrubute and clear the current bank
		blocks := ns[bankIndex]
		ns[bankIndex] = 0
		length := len(ns)
		currentIndex := bankIndex

		for blocks > 0 {
			currentIndex = (currentIndex + 1) % length
			ns[currentIndex]++
			blocks--
		}

		//increment our cycles for each run
		cycles++

		//write each value to our result buffer in order to check if it exists in our result map
		var result strings.Builder
		for _, v := range ns {
			result.WriteString(string(v))
		}
		if _, ok := completed[result.String()]; ok {
			break
		}
		completed[result.String()] = true

	}

	return cycles
}

func countCyclesRedux(ns []int) int {

	completed := make(map[string]int)
	cycles := 0
	set := false

	for {
		// find the index of the highest bank (if tie, take the lowest index)
		bankIndex := 0
		for i, v := range ns[1:] {
			i++
			if v > ns[bankIndex] {
				bankIndex = i
			} else if v == ns[bankIndex] {
				if i < bankIndex {
					bankIndex = i
				}
			}
		}

		//save the blocks to distrubute and clear the current bank
		blocks := ns[bankIndex]
		ns[bankIndex] = 0
		length := len(ns)
		currentIndex := bankIndex

		for blocks > 0 {
			currentIndex = (currentIndex + 1) % length
			ns[currentIndex]++
			blocks--
		}

		//increment our cycles for each run
		cycles++

		//write each value to our result buffer in order to check if it exists in our result map
		var result strings.Builder
		for _, v := range ns {
			result.WriteString(string(v))
		}
		if v, ok := completed[result.String()]; ok {
			if v == 1 && !set {
				cycles = 0
				set = true
			}
			if v == 2 {
				break
			}
		}
		completed[result.String()]++

	}

	return cycles
}

func main() {
	s := `10	3	15	10	5	15	5	15	9	2	5	8	5	2	3	6`

	cycles := strings.Fields(s)

	ints := make([]int, len(cycles))

	for i, v := range cycles {
		var err error
		ints[i], err = strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("The number of cycles are: ", countCycles(ints))
	fmt.Println("The number of loops since the first matched are: ", countCyclesRedux(ints))

}
