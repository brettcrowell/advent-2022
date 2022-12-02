// rock defeats scissors, scissors defeats paper, paper defeats rock
// first column: opponent, second column: response

// rock/a, paper/b, scissers/c
// rock/x, paper/y, scissers/z
// rock/1, paper/2, scissers/3
// lost/0, draw/3, win/6

// rock/1, paper/2, scissers/3
// 1 > 3, 3 > 2, 2 > 1

// score = shape score + outcome

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	baseline := map[string]int{
		"rock":     1,
		"paper":    2,
		"scissers": 3,
	}

	scoring := map[string]map[string]int{
		"rock":     {"rock": 3, "paper": 0, "scissers": 6},
		"paper":    {"rock": 6, "paper": 3, "scissers": 0},
		"scissers": {"rock": 0, "paper": 6, "scissers": 3},
	}

	expectedValue := map[int]map[string]string{}
	for _, desired := range [3]int{0, 3, 6} {
		expectedValue[desired] = map[string]string{}
		for ours, scores := range scoring {
			for theirs, score := range scores {
				if desired == score {
					expectedValue[desired][theirs] = ours
					break
				}
			}
		}
	}

	fmt.Println(expectedValue)

	parse := map[string]string{
		"A": "rock",
		"B": "paper",
		"C": "scissers",
	}

	outcomes := map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}

	input, err := os.Open(os.Args[1])

	if err != nil {
		panic(err)
	}

	defer input.Close()
	scanner := bufio.NewScanner(input)

	total := 0

	for scanner.Scan() {
		args := strings.Split(scanner.Text(), " ")

		theirs := parse[args[0]]
		outcome := outcomes[args[1]]
		ours := expectedValue[outcome][theirs]

		roundScore := baseline[ours] + scoring[ours][theirs]
		total += roundScore
	}

	fmt.Println(total)
}
