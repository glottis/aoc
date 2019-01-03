package utils

import (
	"log"
	"strconv"
	"strings"
)

// ParseInts returns a list of valid ints derived from a string
func ParseInts(s string) []int {

	sLen := len(s)
	ints := make([]int, sLen)

	for i, v := range s {
		if (v-0x30) >= 0 && (v-0x30) <= 9 {
			ints[i] = int(v - 0x30)
		}
	}

	return ints
}

// Max returns the max value of x, y
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y

}

// Min returns the min value of x, y
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y

}

// intGrid creates a [][]int slice from a string
func IntGrid(s string) [][]int {

	lines := strings.Split(strings.TrimSpace(s), "\n")
	grid := make([][]int, len(lines))
	for i, line := range lines {

		subLines := strings.Fields(line)
		grid[i] = make([]int, len(subLines))

		for i2, ss := range subLines {
			value, err := strconv.Atoi(ss)
			if err != nil {
				log.Fatal(err)
			}
			grid[i][i2] = value
		}
	}
	return grid
}
