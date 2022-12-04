package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse(input string) [][]int {
	pairs := strings.Split(input, ",")
	ranges := [][]int{}

	for _, sections := range pairs {
		bounderies := strings.Split(sections, "-")

		low, err := strconv.Atoi(bounderies[0])

		if err != nil {
			panic(err)
		}

		high, err := strconv.Atoi(bounderies[1])

		if err != nil {
			panic(err)
		}

		ranges = append(
			ranges,
			[]int{low, high},
		)
	}

	return ranges
}

func main() {
	input, err := os.Open(os.Args[1])

	if err != nil {
		panic(err)
	}

	defer input.Close()
	scanner := bufio.NewScanner(input)

	contained := 0

	for scanner.Scan() {
		ranges := parse(scanner.Text())

		lowA := ranges[0][0]
		highA := ranges[0][1]
		lowB := ranges[1][0]
		highB := ranges[1][1]

		aContainsB := lowA <= lowB && highA >= highB
		bContainsA := lowB <= lowA && highB >= highA

		if aContainsB || bContainsA {
			contained++
		}

	}

	fmt.Println(contained)
}
