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

		instr := instruction{direction, distance}
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

func render(knots [][]location) {
	for r := 5; r >= 0; r-- {
		line := ""
		for c := 0; c < 6; c++ {
			taken := false
			for k, knot := range knots {
				if !taken {
					location := knot[len(knot)-1]
					if location.x == c && location.y == r {
						line = line + fmt.Sprint(k)
						taken = true
					}
				}
			}
			if !taken {
				line = line + "."
			}
		}
		fmt.Println(line)
	}
	fmt.Println("")
}

func createApply(knots *[][]location) func(instruction) {
	iteration := 0
	return func(instr instruction) {
		heads := (*knots)[0]
		head := heads[len(heads)-1]

		nextHead := move(head, instr)
		(*knots)[0] = append(heads, nextHead)

		for t, trail := range (*knots)[1:] {
			prevPosns := (*knots)[t]
			prev := prevPosns[len(prevPosns)-1]

			curPosns := trail[len(trail)-1]
			cur := curPosns

			xDistance := math.Max(float64(prev.x), float64(cur.x)) -
				math.Min(float64(prev.x), float64(cur.x))

			yDistance := math.Max(float64(prev.y), float64(cur.y)) -
				math.Min(float64(prev.y), float64(cur.y))

			if xDistance > 1 || yDistance > 1 {
				if xDistance > 0 {
					if prev.x > cur.x {
						cur = move(cur, instruction{"R", 1})
					} else {
						cur = move(cur, instruction{"L", 1})
					}
				}

				if yDistance > 0 {
					if prev.y > cur.y {
						cur = move(cur, instruction{"U", 1})
					} else {
						cur = move(cur, instruction{"D", 1})
					}
				}
			}

			(*knots)[t+1] = append((*knots)[t+1], cur)
		}
		iteration++
	}
}

func puzzle01(instructions *[]instruction) int {
	knots := [][]location{
		{{x: 0, y: 0}},
		{{x: 0, y: 0}},
	}

	apply := createApply(&knots)

	for _, instr := range *instructions {
		for i := 0; i < instr.distance; i++ {
			// apply instructions one unit at a time
			apply(instruction{instr.direction, 1})
		}
	}

	return len(dedup(knots[len(knots)-1]))
}

func puzzle02(instructions *[]instruction) int {
	knots := [][]location{}

	for len(knots) < 10 {
		knots = append(knots, []location{{x: 0, y: 0}})
	}

	apply := createApply(&knots)

	for _, instr := range *instructions {
		fmt.Println("==", instr.direction, instr.distance, "==")
		for i := 0; i < instr.distance; i++ {
			// apply instructions one unit at a time
			apply(instruction{instr.direction, 1})
			render(knots)
		}
	}

	tailPosns := knots[len(knots)-1]

	return len(dedup(tailPosns))
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

	fmt.Println(puzzle01(&moves))
	fmt.Println(puzzle02(&moves))
}
