package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

func checkPhrase(s string) bool {

	input := strings.Fields(s)
	m := make(map[string]int)

	for _, v := range input {

		if m[v] != 0 {
			return false
		}

		m[v]++

	}

	return true
}

func checkPhraseRedux(s string) bool {

	input := strings.Fields(s)

	for i, v := range input {
		length := len(v)
		for ii, vv := range input {
			if length == len(vv) && i != ii {
				s1 := strings.Split(v, "")
				sort.Strings(s1)
				s2 := strings.Split(vv, "")
				sort.Strings(s2)
				s3 := strings.Join(s1, "")
				s4 := strings.Join(s2, "")

				if s3 == s4 {
					return false
				}
			}
		}

	}

	return true
}

func main() {

	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	inputSlice := strings.Split(string(input), "\n")

	sum := 0

	for _, v := range inputSlice {
		if checkPhraseRedux(v) {
			sum++
		}
	}

	fmt.Println("Number of correct phrases are: ", sum)

}
