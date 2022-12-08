package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Node struct {
	parent   *Node
	name     string
	size     int
	children []*Node
}

func find(name string, children ...*Node) *Node {
	for _, child := range children {
		if child.name == name {
			return child
		}
	}
	return nil
}

func mkdir(name string, parent *Node) Node {
	return Node{
		name:     name,
		parent:   parent,
		size:     0,
		children: []*Node{},
	}
}

func touch(name string, size int, parent *Node) Node {
	return Node{
		name:   name,
		parent: parent,
		size:   size,
	}
}

func createFsReader(root *Node) func(string) {
	cd, _ := regexp.Compile("\\$\\scd\\s(.*)")
	resp, _ := regexp.Compile("(\\d+)\\s(.+)")

	// keep a pointer to the working directory
	wd := root

	return func(line string) {
		if cd.MatchString(line) {
			// the cmd argument is the first capture group
			arg := cd.FindAllStringSubmatch(line, -1)[0][1:][0]

			switch arg {
			case "..":
				// update the working directory to the parent
				wd = wd.parent
			default:
				// check dir already exists
				dir := find(arg, wd.children...)

				if dir == nil {
					// create a new empty dir
					new := mkdir(arg, wd)
					dir = &new

					// create dir and add a pointer to the parent
					wd.children = append(wd.children, dir)
				}

				// update the working directory pointer
				wd = dir
			}
		}

		if resp.MatchString(line) {
			matches := resp.FindAllStringSubmatch(line, -1)[0][1:]

			name := matches[1]
			size, _ := strconv.Atoi(matches[0])

			// create a file with a pointer to the parent
			file := touch(name, size, wd)

			// append the child to the parent
			wd.children = append(wd.children, &file)
		}
	}
}

func getAbsolutePath(node *Node, path ...string) string {
	if node.parent != nil {
		return getAbsolutePath(
			node.parent,
			append([]string{node.name}, path...)...,
		)
	}

	return strings.Join(path, "/") + node.name
}

func getDirSizes(sizes *map[string]int, node *Node) {
	if len(node.children) == 0 {
		parent := node.parent

		for parent != nil {
			path := getAbsolutePath(parent)
			(*sizes)[path] += node.size
			parent = parent.parent
		}
	}

	for _, child := range node.children {
		getDirSizes(sizes, child)
	}
}

func puzzle01(dirSizes *map[string]int) int {
	sum := 0

	for _, size := range *dirSizes {
		if size <= 100000 {
			sum += size
		}
	}

	return sum
}

func puzzle02(dirSizes *map[string]int) int {
	total := 70000000
	required := 30000000

	unused := total - (*dirSizes)[""]
	var smallest int

	for _, size := range *dirSizes {
		if size >= required-unused && (smallest == 0 || size < smallest) {
			smallest = size
		}
	}

	return smallest
}

func main() {
	input, err := os.Open(os.Args[1])

	if err != nil {
		panic(err)
	}

	defer input.Close()
	scanner := bufio.NewScanner(input)

	var fs Node

	read := createFsReader(&fs)

	for scanner.Scan() {
		read(scanner.Text())
	}

	dirSizes := map[string]int{}
	getDirSizes(&dirSizes, &fs)

	fmt.Println(puzzle01(&dirSizes))
	fmt.Println(puzzle02(&dirSizes))

}
