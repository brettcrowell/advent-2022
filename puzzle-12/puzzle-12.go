package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Node struct {
	x      int
	y      int
	label  string
	height int
}

func includes(path []*Node, node *Node) int {
	for s, square := range path {
		if node == square {
			return s
		}
	}
	return -1
}

func render(grid [][]*Node, path []*Node) {
	output := ""
	for _, row := range grid {
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

	// move cursor up to the 0,0 position in the grid
	// clear := "\033[" + fmt.Sprint(len(grid)+1) + "A"
	// output = clear + output

	fmt.Println(output)
}

func getNeighbors(grid [][]*Node, square *Node) []*Node {
	neighbors := []*Node{}

	if square.y < len(grid)-1 {
		// down
		neighbors = append(neighbors, (grid)[square.y+1][square.x])
	}

	if square.x < len((grid)[0])-1 {
		// right
		neighbors = append(neighbors, (grid)[square.y][square.x+1])
	}

	if square.y > 0 {
		// up
		neighbors = append(neighbors, (grid)[square.y-1][square.x])
	}

	if square.x > 0 {
		// left
		neighbors = append(neighbors, (grid)[square.y][square.x-1])
	}

	return neighbors
}

func getShortestPath(
	graph map[*Node][]*Node,
	start *Node,
	destination *Node,
) []*Node {
	// create a map of distances to each node
	distances := map[*Node]float64{}

	// keep track of all visited nodes
	visited := map[*Node]bool{}

	for node := range graph {
		distances[node] = math.Inf(1)
		visited[node] = false
	}

	// seed distances map with starting node
	distances[start] = 0

	// keep track of the closest neighbor for each node
	preceding := map[*Node]*Node{}

	for range graph {
		distance := math.Inf(1)

		// get next unprocessed node, initially will equal start
		var next *Node

		for node := range graph {
			if next != nil {
				distance = distances[next]
			}

			if !visited[node] && distances[node] < distance {
				next = node
			}
		}

		// update the node to visited
		visited[next] = true

		for _, neighbor := range graph[next] {
			// the distance to this neighbor following a path through "closest" node
			distance := distances[next] + 1.0

			if distance < distances[neighbor] {
				// if shorter, accept it
				distances[neighbor] = distance

				// mark this as the neighbor's preceding node
				preceding[neighbor] = next
			}
		}

		if distances[next] == math.Inf(1) {
			panic("No path found between nodes")
		}
	}

	shortest := []*Node{}
	node := destination

	for node != start {
		// reverse engineer the map of neighbors to find the shortest path
		shortest = append(shortest, preceding[node])
		node = preceding[node]
	}

	return shortest
}

func puzzle01(grid [][]*Node, graph map[*Node][]*Node) int {
	var start *Node
	var end *Node

	for node := range graph {
		switch (*node).label {
		case "S":
			{
				start = node
				break
			}

		case "E":
			{
				end = node
				break
			}
		}
	}

	path := getShortestPath(graph, start, end)
	render(grid, path)

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

	for c, char := range strings.Split("abcdefghijklmnopqrstuvwxyz", "") {
		heights[char] = c
	}

	heights["S"] = heights["a"]
	heights["E"] = heights["z"]

	grid := [][]*Node{}

	r := 0
	for scanner.Scan() {
		currentRow := []*Node{}
		for c, col := range strings.Split(scanner.Text(), "") {
			height := heights[col]
			node := Node{c, r, col, height}
			currentRow = append(currentRow, &node)
		}
		grid = append(grid, currentRow)
		r++
	}

	graph := map[*Node][]*Node{}

	for _, row := range grid {
		for _, node := range row {
			neighbors := getNeighbors(grid, node)
			for _, neighbor := range neighbors {
				gap := neighbor.height - node.height

				if gap < 2 {
					graph[node] = append(graph[node], neighbor)
				}
			}
		}
	}

	// render(&grid, &Nodes{})
	fmt.Println(puzzle01(grid, graph))
}
