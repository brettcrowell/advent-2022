package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Square struct {
	x      int
	y      int
	label  string
	height int
	target bool
}

type Position struct {
	square *Square
}

func render(grid *[][]*Square) {
	for _, row := range *grid {
		for _, col := range row {
			fmt.Print(col.label)
		}
		fmt.Println()
	}
}

func hasVisited(path *[]*Square, target *Square) bool {
	for _, square := range *path {
		if square == target {
			return true
		}
	}
	return false
}

func isValid(path *[]*Square, target *Square, neighbor *Square) bool {
	gap := neighbor.height - (*target).height
	visited := hasVisited(path, neighbor)

	return gap < 2 && !visited
}

func getNeighbors(grid *[][]*Square, square *Square) []*Square {
	neighbors := []*Square{}

	if square.y < len(*grid)-1 {
		// down
		neighbors = append(neighbors, (*grid)[square.y+1][square.x])
	}

	if square.x < len((*grid)[0])-1 {
		// right
		neighbors = append(neighbors, (*grid)[square.y][square.x+1])
	}

	if square.y > 0 {
		// up
		neighbors = append(neighbors, (*grid)[square.y-1][square.x])
	}

	if square.x > 0 {
		// left
		neighbors = append(neighbors, (*grid)[square.y][square.x-1])
	}

	return neighbors
}

func getValid(
	path *[]*Square,
	neighbors *[]*Square,
	target *Square,
) []*Square {
	valid := []*Square{}
	for _, neighbor := range *neighbors {
		if isValid(path, target, neighbor) {
			valid = append(valid, neighbor)
		}
	}
	return valid
}

func step(grid *[][]*Square, target *Square) []*Square {
	path := []*Square{}
	next := target

	for next != nil {
		fmt.Println(next.x, next.y, next.label)

		neighbors := getNeighbors(grid, next)
		valid := getValid(&path, &neighbors, next)

		fmt.Println("neighbors")
		for _, neighbor := range valid {
			fmt.Println(neighbor.x, neighbor.y, neighbor.label)
		}
		fmt.Println("/neighbors")

		if len(valid) == 0 {
			break
		}

		if valid[0].target {
			break
		}

		next = valid[0]
		path = append(path, next)
	}

	fmt.Println(len(path))

	return path

	// if (*target).target {
	// 	return []*Square{}
	// }

	// return append(step(grid, neighbors[0]), target)
}

func puzzle01(squares *[]Square, grid *[][]*Square) int {
	square := (*squares)[0]
	path := step(grid, &square)

	fmt.Println(path)

	return 0
}

func main() {
	input, err := os.Open(os.Args[1])

	if err != nil {
		panic(err)
	}

	defer input.Close()
	scanner := bufio.NewScanner(input)

	heights := map[string]int{"S": -1}

	for c, char := range strings.Split("abcdefghijklmnopqrstuvwxyzE", "") {
		heights[char] = c
	}

	squares := []Square{}
	grid := [][]*Square{}

	r := 0
	for scanner.Scan() {
		currentRow := []*Square{}
		for c, col := range strings.Split(scanner.Text(), "") {
			height := heights[col]
			squares = append(squares, Square{c, r, col, height, col == "E"})
			currentRow = append(currentRow, &squares[len(squares)-1])
		}
		grid = append(grid, currentRow)
		r++
	}

	render(&grid)
	fmt.Println(puzzle01(&squares, &grid))
}
