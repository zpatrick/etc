package main

import "fmt"

// Given a 2D array, print it in spiral form
// Assume the array is always a rectangle

/*
end scenario?
	going right and no more down
	going down and no more left
	going left and no more up
	going up and no more right

hold maxLeft, maxDown, maxRight, maxUp

*/

//

type Direction int

const (
	Up    = 0
	Down  = 1
	Left  = 2
	Right = 3
)

type Cell struct {
	Row, Col int
}

func SpiralTraverse(numRows, numCols int, arr [][]int, fn func(v int)) {
	maxRightCol := numCols - 1
	maxLeftCol := 0
	maxBottomRow := numRows - 1
	maxTopRow := 0
	direction := Right
	c := Cell{0, 0}

	var directionSwitch bool
	walk := func() bool {
		if !directionSwitch {
			fn(arr[c.Row][c.Col])
		}

		switch direction {
		case Right:
			if c.Col == maxRightCol {
				maxTopRow++
				direction = Down
				directionSwitch = true
				return c.Row != maxBottomRow
			}

			c.Col++
		case Left:
			if c.Col == maxLeftCol {
				maxBottomRow--
				direction = Up
				directionSwitch = true
				return c.Row != maxLeftCol
			}

			c.Col--
		case Up:
			if c.Row == maxTopRow {
				maxLeftCol++
				direction = Right
				directionSwitch = true
				return c.Col != maxRightCol
			}

			c.Row--
		case Down:
			if c.Row == maxBottomRow {
				maxRightCol--
				direction = Left
				directionSwitch = true
				return c.Col != maxLeftCol
			}

			c.Row++
		}

		directionSwitch = false
		return true
	}

	for walk() {
	}
}

func main() {
	arr := [][]int{
		[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8},
		[]int{9, 10, 11, 12},
		[]int{13, 14, 15, 16},
	}

	// expected: 1 2 3 4 8 12 16 15 14 13 9 5 6 7 11 10
	SpiralTraverse(4, 4, arr, func(v int) {
		fmt.Printf("%d ", v)
	})
	
	fmt.Println()
}
