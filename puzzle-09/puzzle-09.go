package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	direction string
	distance  int
	axis      string
}

type location struct {
	x int
	y int
}

func createReader(steps *[]instruction) func(string) {
	return func(step string) {
		parsed := strings.Split(step, " ")
		direction := parsed[0]
		distance, _ := strconv.Atoi(parsed[1])

		axis := "X"

		if direction == "U" || direction == "D" {
			axis = "Y"
		}

		instr := instruction{direction, distance, axis}
		*steps = append(*steps, instr)
	}
}

func dedup(locations []location) []location {
	locos := map[location]bool{}

	for _, location := range locations {
		locos[location] = true
	}

	unique := []location{}

	for location := range locos {
		unique = append(unique, location)
	}

	return unique
}

func move(end location, instr instruction) location {
	switch instr.direction {
	case "R":
		return location{end.x + 1, end.y}
	case "L":
		return location{end.x - 1, end.y}
	case "U":
		return location{end.x, end.y + 1}
	case "D":
		return location{end.x, end.y - 1}
	default:
		return end
	}
}

func render(head location, tail location) {
	for r := 4; r >= 0; r-- {
		line := ""
		for c := 0; c < 6; c++ {
			if head.x == c && head.y == r {
				line = line + "H"
			} else if tail.x == c && tail.y == r {
				line = line + "T"
			} else {
				line = line + "."
			}
		}
		fmt.Println(line)
	}
	fmt.Println("")
}

func createApply(
	heads *[]location,
	tails *[]location,
) func(instruction) {

	return func(instr instruction) {
		head := (*heads)[len(*heads)-1]
		tail := (*tails)[len(*tails)-1]

		nextHead := move(head, instr)
		nextTail := tail

		xDistance := math.Max(float64(nextHead.x), float64(tail.x)) -
			math.Min(float64(nextHead.x), float64(tail.x))

		yDistance := math.Max(float64(nextHead.y), float64(tail.y)) -
			math.Min(float64(nextHead.y), float64(tail.y))

		if xDistance != yDistance && xDistance+yDistance > 1 {
			if xDistance > 0 {
				nextTail = location{head.x, nextTail.y}
			}

			if yDistance > 0 {
				nextTail = location{nextTail.x, head.y}
			}
		}

		*heads = append(*heads, nextHead)
		*tails = append(*tails, nextTail)

		fmt.Println("x", xDistance, "y", yDistance)
		render(nextHead, nextTail)
	}
}

func puzzle01(instructions *[]instruction) int {
	head := []location{{x: 0, y: 0}}
	tail := []location{{x: 0, y: 0}}

	apply := createApply(&head, &tail)

	for _, instr := range *instructions {
		for i := 0; i < instr.distance; i++ {
			// apply instructions one unit at a time
			apply(instruction{instr.direction, 1, instr.axis})
		}
	}

	return len(dedup(tail))
}

func main() {
	input, err := os.Open(os.Args[1])

	if err != nil {
		panic(err)
	}

	defer input.Close()
	scanner := bufio.NewScanner(input)

	moves := []instruction{}
	read := createReader(&moves)

	for scanner.Scan() {
		read(scanner.Text())
	}

	fmt.Println(moves)
	fmt.Println(puzzle01(&moves))
}
