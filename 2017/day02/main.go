package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/glottis/aoc/utils"
)

func main() {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(checkSumSheet(string(file)))
	fmt.Println(checkSumSheetRedux(string(file)))

}

func checkSumSheetRedux(s string) int {
	grid := utils.IntGrid(s)
	var sum int
	for _, v := range grid {
		sort.Sort(sort.Reverse(sort.IntSlice(v)))
		// slice is sorted so [0] should always be max value
		for ii, vv := range v {
			for _, vvv := range v[ii+1:] {
				if vv%vvv == 0 {
					sum += (vv / vvv)
					goto mainLoop
				}

			}

		}
	mainLoop:
	}
	return sum
}

func checkSumSheet(s string) int {
	var sum int

	s1 := strings.Split(s, "\n")

	for _, v := range s1 {
		var min int
		var max int
		s2 := strings.Fields(v)
		for i, v := range s2 {
			n, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			if i == 0 {
				min = n
				max = n
			} else if n > max {
				max = n
			} else if n < min {
				min = n
			}
		}
		sum += (max - min)

	}
	return sum
}
