package main

import (
	"bufio"
	"fmt"
	"os"
)

func allDifferent(chars ...string) bool {
	seen := map[string]bool{}

	for _, char := range chars {
		if seen[char] {
			return false
		}
		seen[char] = true
	}

	return true
}

func createIngester(marker_len int) func(string) (bool, int) {
	iterations := 0
	window := []string{}

	return func(char string) (bool, int) {
		iterations++

		window = append(window, char)

		if len(window) == marker_len {
			if allDifferent(window...) {
				return true, iterations
			} else {
				window = window[1:]
			}
		}

		return false, iterations
	}
}

func main() {
	input, err := os.Open(os.Args[1])

	if err != nil {
		panic(err)
	}

	defer input.Close()
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanRunes)

	ingest := createIngester(14)

	for scanner.Scan() {
		match, iterations := ingest(scanner.Text())

		if match {
			fmt.Println(iterations)
			break
		}
	}

}
