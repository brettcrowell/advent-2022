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
}

type Position struct {
	square *Square
}

func includes(path *[]*Square, node *Square) int {
	for s, square := range *path {
		if node == square {
			return s
		}
	}
	return -1
}

func render(grid *[][]*Square, path *[]*Square) {
	for _, row := range *grid {
		for _, col := range row {
			move := includes(path, col)
			if move > -1 {
				fmt.Print(move % 10)
			} else {
				fmt.Print(col.label)
			}
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

func getPath(
	grid *[][]*Square,
	start *Square,
	destination string,
) []*Square {

	// keep track of all of the visited nodes
	visited := []*Square{}

	next := start

	for next != nil {
		fmt.Println(next.x, next.y)

		// get all adjacent squares
		neighbors := getNeighbors(grid, next)

		// filter down to squares which are reachable
		valid := getValid(&visited, &neighbors, next)

		if valid[0].label == destination {
			// if we're at our destination, return
			break
		}

		if len(valid) == 0 {
			// if no valid edges, go back one step and try again
			next = visited[len(visited)-1]

			if next == start {
				// if we have exhausted all edges from the start
				break
			}

			continue
		}

		// check the next valid node
		next = valid[0]
		visited = append(visited, next)
	}

	return visited
}

func puzzle01(squares *[]Square, grid *[][]*Square) int {
	square := (*squares)[0]
	path := getPath(grid, &square, "E")

	render(grid, &path)

	return len(path)
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
			squares = append(squares, Square{c, r, col, height})
			currentRow = append(currentRow, &squares[len(squares)-1])
		}
		grid = append(grid, currentRow)
		r++
	}

	fmt.Println(puzzle01(&squares, &grid))
}
