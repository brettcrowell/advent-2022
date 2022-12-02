package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func sum(rations []int) int {
	total := 0

	for _, calories := range rations {
		total += calories
	}

	return total
}

func tally(logs [][]int) []int {
	sums := make([]int, len(logs))

	for _, log := range logs {
		sums = append(sums, sum(log))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sums)))

	return sums
}

func top(tallies []int, n int) int {
	partialSum := 0

	for i := 0; i < 3; i++ {
		partialSum += tallies[i]
	}

	return partialSum
}

func main() {
	input, err := os.Open(os.Args[1])

	if err != nil {
		panic(err)
	}

	defer input.Close()

	scanner := bufio.NewScanner(input)
	logs := [][]int{{}}

	for scanner.Scan() {
		if scanner.Text() != "" {
			calories, err := strconv.Atoi(scanner.Text())

			if err != nil {
				panic(err)
			}

			lastElf := len(logs) - 1
			logs[lastElf] = append(logs[lastElf], calories)
		} else {
			logs = append(logs, make([]int, 0))
		}
	}

	tallies := tally(logs)

	fmt.Println(tallies[0])
	fmt.Println(top(tallies, 3))

}
