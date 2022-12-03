package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getCommonItems(first []string, second []string) []string {
	common := []string{}

	for _, firstItem := range first {
		for _, secondItem := range second {
			if firstItem == secondItem {
				common = append(common, firstItem)
			}
		}
	}

	return common
}

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

	row := 1
	contents := [3][]string{}
	sum := 0

	for scanner.Scan() {
		content := strings.Split(scanner.Text(), "")
		contents[row%3] = content
		if row%3 == 0 {
			firstCommon := getCommonItems(contents[0], contents[1])
			secondCommon := getCommonItems(firstCommon, contents[2])

			sum += itemValues[secondCommon[0]]

			contents = [3][]string{}
		}
		row++
	}

	fmt.Println(sum)
}
