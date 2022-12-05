package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

func main() {
	input, err := os.Open(os.Args[1])

	if err != nil {
		panic(err)
	}

	defer input.Close()
	scanner := bufio.NewScanner(input)

	r, _ := regexp.Compile("move (\\d+) from (\\d+) to (\\d+)")
	stacks := map[int][]string{}
	stacks_complete := false

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			stacks_complete = true
			continue
		}

		if !stacks_complete {

		stacks:
			for c, char := range line {
				if unicode.IsDigit(char) {
					break stacks
				}

				if (c+1)%2 == 0 && !unicode.IsSpace(char) {
					stack_num := c/4 + 1
					stack := stacks[stack_num]

					if stack == nil {
						stacks[stack_num] = []string{}
					}

					stacks[stack_num] = append(stacks[stack_num], string(char))
				}
			}
			continue
		}

		moves := r.FindAllStringSubmatch(line, -1)[0][1:]

		num, _ := strconv.Atoi(moves[0])
		from, _ := strconv.Atoi(moves[1])
		to, _ := strconv.Atoi(moves[2])

		for i := 0; i < num; i++ {
			crate := stacks[from][0]
			stacks[from] = stacks[from][1:]
			stacks[to] = append([]string{crate}, stacks[to]...)
		}
	}

	answer := ""

	for i := 1; i <= len(stacks); i++ {
		crates := stacks[i]
		answer += crates[0]

	}

	fmt.Println(answer)
}
