package main

import (
	"errors"
	"fmt"
	"math"
)

func calcSteps(point int) (int, error) {
	if point == 1 {
		return 0, nil
	}

	length := point * 2
	mainCorners := make([]int, length)

	for i := range mainCorners {
		if i == 0 {
			mainCorners[i] = 9
		} else {
			sqrt := math.Sqrt(float64(mainCorners[i-1]))
			newValue := int((sqrt + 2) * (sqrt + 2))
			mainCorners[i] = newValue
		}

	}

	rowCount := 0
	extraCorners := make([]int, 4)
	stepsVert := 0
	stepsMiddle := 0
	stepsCorner := 0

	for i, v := range mainCorners {

		stepsVert = i + 1
		if v >= point {
			rowCount = int(math.Sqrt(float64(v))) - 1
			extraCorners[3] = v
			extraCorners[2] = extraCorners[3] - rowCount
			extraCorners[1] = extraCorners[2] - rowCount
			extraCorners[0] = extraCorners[1] - rowCount

			for _, corner := range extraCorners {
				if corner >= point {
					stepsCorner = corner - point
					stepsMiddle = rowCount / 2

					// if the point is on the middle of the row
					if stepsCorner == stepsMiddle {
						return stepsVert, nil
					}
					stepsHort := int(math.Abs(float64(stepsMiddle - stepsCorner)))
					return (stepsVert + stepsHort), nil
				}
			}
		}
	}
	return 0, errors.New("Not found")
}

func main() {

	m := [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}

	fmt.Println(m)

	for _, v := range m {
		for _, vv := range v {
			fmt.Printf("%v ", vv)
		}
		fmt.Println()
	}
	length := len(m) * len(m[0])

	steps := 0

	rightBound := len(m) - 1
	downBound := len(m) - 1
	leftBound := 0
	upBound := 1
	stepsPlay := 0

	for steps < length {
		for i := 0; i < rightBound; i++ {
			fmt.Print(stepsPlay + 1)
		}
		rightBound--
		for i := 0; i < downBound; i-- {
			fmt.Print(stepsPlay + 1)
		}
		downBound--

	}

}
