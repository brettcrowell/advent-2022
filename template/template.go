package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetCommonItems(first []string, second []string) []string {
	set := map[string]bool{}

	for _, firstItem := range first {
		for _, secondItem := range second {
			if firstItem == secondItem {
				set[firstItem] = true
			}
		}
	}

	common := []string{}
	for key := range set {
		common = append(common, key)
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
			firstCommon := GetCommonItems(contents[0], contents[1])
			secondCommon := GetCommonItems(firstCommon, contents[2])

			sum += itemValues[secondCommon[0]]
		}

		row++
	}

	fmt.Println(sum)
}
