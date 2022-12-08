package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Node struct {
	parent   *Node
	name     string
	size     int
	children []*Node
}

func createFsReader(fs *Node) func(string) {
	cd, _ := regexp.Compile("\\$\\scd\\s(.*)")
	resp, _ := regexp.Compile("(\\d+)\\s(.+)")

	wd := fs

	return func(line string) {
		if cd.MatchString(line) {
			parts := cd.FindAllStringSubmatch(line, -1)[0][1:]
			arg := parts[0]

			switch arg {
			case "..":
				wd = wd.parent
			default:
				node := Node{
					name:     arg,
					parent:   wd,
					size:     0,
					children: []*Node{},
				}

				wd.children = append(wd.children, &node)
				wd = &node
			}
		}

		if resp.MatchString(line) {
			parts := resp.FindAllStringSubmatch(line, -1)[0][1:]

			name := parts[1]
			size, _ := strconv.Atoi(parts[0])

			file := Node{
				name:   name,
				parent: wd,
				size:   size,
			}

			wd.children = append(wd.children, &file)
		}
	}
}

func getAbsolutePath(node *Node) string {
	if node.parent != nil {
		return getAbsolutePath(node.parent) + "/" + node.name
	}

	return node.name
}

func getDirSizes(sizes *map[string]int, node *Node) {
	if len(node.children) == 0 {
		parent := node.parent

		for parent != nil {
			(*sizes)[getAbsolutePath(parent)] += node.size
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

	fs := Node{}

	read := createFsReader(&fs)

	for scanner.Scan() {
		read(scanner.Text())
	}

	dirSizes := map[string]int{}
	getDirSizes(&dirSizes, &fs)

	fmt.Println(puzzle01(&dirSizes))
	fmt.Println(puzzle02(&dirSizes))

}
