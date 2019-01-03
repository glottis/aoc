package main

import (
	"errors"
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
}
