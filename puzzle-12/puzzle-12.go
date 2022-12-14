package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Node struct {
	x      int
	y      int
	label  string
	height int
}

type Nodes []*Node

func (n Nodes) Len() int {
	return len(n)
}

func (n Nodes) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n Nodes) Less(i, j int) bool {
	return (*n[j]).height < (*n[i]).height
}

func includes(path *Nodes, node *Node) int {
	for s, square := range *path {
		if node == square {
			return s
		}
	}
	return -1
}

func render(grid *[][]*Node, path *Nodes) {
	output := ""
	for _, row := range *grid {
		for _, col := range row {
			move := includes(path, col)
			if move > -1 {
				output += col.label
			} else {
				switch col.label {
				case "S":
					output += "S"
				case "E":
					output += "E"
				default:
					output += "."
				}
			}
		}
		output += "\n"
	}
	fmt.Print("\033[" + fmt.Sprint(len(*grid)+1) + "A" + output)
}

func hasVisited(path *Nodes, target *Node) bool {
	for _, square := range *path {
		if square == target {
			return true
		}
	}
	return false
}

func getNeighbors(grid *[][]*Node, square *Node) Nodes {
	neighbors := Nodes{}

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

func filterValid(visited *map[*Node]bool, neighbors *Nodes, target *Node) Nodes {
	valid := Nodes{}

	for _, neighbor := range *neighbors {
		gap := neighbor.height - (*target).height
		hasVisited := (*visited)[neighbor]

		isValid := gap < 2 && !hasVisited

		if isValid {
			valid = append(valid, neighbor)
		}
	}

	return valid
}

func getPath(
	grid *[][]*Node,
	start *Node,
	destination string,
) Nodes {

	visited := map[*Node]bool{}

	// keep track of the path
	path := Nodes{}

	next := start

	for next != nil {

		// if we've never been here, add it to the list
		visited[next] = true

		// get all adjacent squares
		neighbors := getNeighbors(grid, next)

		// filter down to squares which are reachable
		valid := filterValid(&visited, &neighbors, next)

		if len(valid) == 0 {
			if next == start {
				// if we're back at the start with no valid neighbors
				// we have nowhere to go!
				break
			}

			visited[path[len(path)-1]] = false

			// roll back the last node
			path = path[:len(path)-1]

			// if no valid edges, go back one step and try again
			next = path[len(path)-1]

			continue
		}

		// keep appending to path
		path = append(path, next)

		if valid[0].label == destination {
			// if we're at our destination, return
			return path
		}

		// sort reachable nodes by distance
		sort.Sort(Nodes(valid))

		// check the next valid node
		next = valid[0]
	}

	panic("No path found between nodes")
}

func puzzle01(nodes *[]Node, grid *[][]*Node) int {
	var start *Node

	for _, node := range *nodes {
		if node.label == "S" {
			start = &node
			break
		}
	}

	path := getPath(grid, start, "E")

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

	heights := map[string]int{}

	for c, char := range strings.Split("abcdefghijklmnopqrstuvwxyzE", "") {
		heights[char] = c
	}

	heights["S"] = heights["a"]
	heights["E"] = heights["z"]

	squares := []Node{}
	grid := [][]*Node{}

	r := 0
	for scanner.Scan() {
		currentRow := []*Node{}
		for c, col := range strings.Split(scanner.Text(), "") {
			height := heights[col]
			squares = append(squares, Node{c, r, col, height})
			currentRow = append(currentRow, &squares[len(squares)-1])
		}
		grid = append(grid, currentRow)
		r++
	}

	// render(&grid, &Nodes{})
	fmt.Println(puzzle01(&squares, &grid))
}
