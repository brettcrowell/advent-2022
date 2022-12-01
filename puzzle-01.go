package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	input, err := os.Open(os.Args[1])

	if err != nil {
		panic(err)
	}

	defer input.Close()

	scanner := bufio.NewScanner(input)
	rations := make(map[int]int)

	curElf := 0
	rations[curElf] = 0

	for scanner.Scan() {
		if scanner.Text() != "" {
			calories, err := strconv.Atoi(scanner.Text())

			if err != nil {
				panic(err)
			}

			rations[curElf] = rations[curElf] + calories
		} else {
			curElf++
		}
	}

	calories := make([]int, len(rations))

	for elf, elfCalories := range rations {
		calories[elf] = elfCalories
	}

	sort.Ints(calories)

	length := len(calories)
	fmt.Println(calories[length-1] + calories[length-2] + calories[length-3])
}
