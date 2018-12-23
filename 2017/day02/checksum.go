package day02

import (
	"encoding/csv"
	"log"
	"sort"
	"strconv"
	"strings"
)

/*
	CheckSum calculates the sum of the min value and the max value from the rows of a spreadsheet.
	Needs a string input and the delimter in form of a rune.
*/
func CheckSum(input string, option rune) int {

	r := csv.NewReader(strings.NewReader(input))
	r.Comma = option

	var sum int

	record, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range record {

		var min int
		var max int

		for index, value := range v {

			i, _ := strconv.Atoi(value)
			if index == 0 {

				min = i
				max = i

			} else if i > max {
				max = i
			} else if i < min {
				min = i
			}
		}

		sum += (max - min)
	}

	return sum
}

/*
CheckSumRedux acts like CheckSum but sums up the evenly divisible values from each row
*/
func CheckSumRedux(input string, option rune) int {

	r := csv.NewReader(strings.NewReader(input))
	r.Comma = option

	var sum int

	record, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range record {

		sliceLength := len(v)
		intSlice := make([]int, sliceLength)

		// populate the slice in order to sort it
		for index, value := range v {

			i, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal(err)
			}

			intSlice[index] = i
		}
		// reverse sort it
		sort.Sort(sort.Reverse(sort.IntSlice(intSlice)))

		// now will we start at the top of the list (desc sorted) and check if the first
		// value is evenly divided by the next value, if not we will continue with the next one
		for indexBase, valueBase := range intSlice {
			var found bool
			for _, valueNew := range intSlice[indexBase+1:] {
				if valueBase%valueNew == 0 {
					sum += valueBase / valueNew
					found = true
					break
				}
			}
			if found {
				break
			}
		}

	}

	return sum
}
