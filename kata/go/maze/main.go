package main

import (
	"fmt"
	"strconv"
)

/*
 --- --- ---
| 0   1   2 |
|    --- ---|
| 3   4   5 |
|        ---|
| 6 | 7   8
|--- --- ---

solution is: 0, 3, 4, 7, 8
*/

type Cell struct {
	ID          string
	IsEnd       bool
	Connections []*Cell
}

func SolveMaze(cells []*Cell) ([]string, error) {
	path, ok := solveMazeRecurse(cells[0], map[*Cell]bool{})
	if !ok {
		return nil, fmt.Errorf("No path found!")
	}

	return path, nil
}

func solveMazeRecurse(current *Cell, visited map[*Cell]bool) ([]string, bool) {
	visited[current] = true
	if current.IsEnd {
		return []string{current.ID}, true
	}

	for _, cell := range current.Connections {
		if _, ok := visited[cell]; ok {
			continue
		}

		path, ok := solveMazeRecurse(cell, visited)
		if ok {
			return append([]string{current.ID}, path...), true
		}
	}

	return nil, false
}

/*

Create and solve this maze:
 --- --- ---
| 0   1   2 |
|    --- ---|
| 3   4   5 |
|        ---|
| 6 | 7   8
|--- --- ---

Solution is: 0, 3, 4, 7, 8
*/

func main() {
	cells := make([]*Cell, 9)
	for i := range cells {
		cells[i] = &Cell{ID: strconv.Itoa(i)}
	}

	cells[0].Connections = []*Cell{cells[1], cells[3]}
	cells[1].Connections = []*Cell{cells[0], cells[2]}
	cells[2].Connections = []*Cell{cells[1]}
	cells[3].Connections = []*Cell{cells[0], cells[4], cells[6]}
	cells[4].Connections = []*Cell{cells[3], cells[5], cells[7]}
	cells[5].Connections = []*Cell{cells[4]}
	cells[6].Connections = []*Cell{cells[3]}
	cells[7].Connections = []*Cell{cells[4], cells[8]}
	cells[8].Connections = []*Cell{cells[7]}
	cells[8].IsEnd = true

	fmt.Println(SolveMaze(cells))
}
