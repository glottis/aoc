package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type node struct {
	id     string
	weight int
	childs []string
}

var children map[string]bool

func parseFile(r io.Reader) []node {

	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)

	var result []node
	for s.Scan() {
		result = append(result, parseNode(s.Text()))
	}

	return result
}

func parseNode(s string) node {
	fields := strings.Fields(s)
	// id is the first field, weight is the second field in "()", children in the rest
	n := node{}
	n.id = fields[0]
	n.weight, _ = strconv.Atoi(fields[1][1 : len(fields[1])-1])

	// if the node does not have any children, the length wil be 2 after fields function. More if it has children
	if len(fields) > 2 {
		length := len(fields[3:])
		n.childs = make([]string, length)
		for i, v := range fields[3:] {
			v = strings.Replace(v, ",", "", -1)
			n.childs[i] = v
			children[v] = true

		}

	} else {
		children[n.id] = true
	}

	return n
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	children = make(map[string]bool)

	nodes := parseFile(f)

	for _, v := range nodes {
		if len(v.childs) > 0 {
			if _, ok := children[v.id]; !ok {
				fmt.Println("The bottom pillar is: ", v.id)
				break
			}

		}
	}

}
