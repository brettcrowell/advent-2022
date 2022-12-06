package main

import (
	"bufio"
	"fmt"
	"os"
)

func all_different(chars ...string) bool {
	seen := map[string]bool{}

	for _, char := range chars {
		if seen[char] {
			return false
		}
		seen[char] = true
	}

	return true
}

func main() {
	input, err := os.Open(os.Args[1])

	if err != nil {
		panic(err)
	}

	defer input.Close()
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanRunes)

	marker_len := 14
	iterations := 1
	trailing := []string{}

	for scanner.Scan() {
		current := scanner.Text()
		trailing = append(trailing, current)

		if len(trailing) == marker_len {
			if all_different(trailing...) {
				fmt.Println(iterations)
				break
			} else {
				trailing = trailing[1:]
			}
		}

		iterations++
	}

}
