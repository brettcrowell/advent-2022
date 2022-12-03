package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	items := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	itemValues := map[string]int{}

	for l, letter := range strings.Split(items, "") {
		itemValues[letter] = l + 1
	}

	input, err := os.Open(os.Args[1])

	if err != nil {
		panic(err)
	}

	defer input.Close()
	scanner := bufio.NewScanner(input)

	sum := 0

	for scanner.Scan() {
		contents := strings.Split(scanner.Text(), "")
		half := len(contents) / 2

		first := contents[0:half]
		second := contents[half:]

	outer:
		for _, firstItem := range first {
			for _, secondItem := range second {
				if firstItem == secondItem {
					sum += itemValues[firstItem]
					break outer
				}
			}
		}
	}

	fmt.Println(sum)
}
