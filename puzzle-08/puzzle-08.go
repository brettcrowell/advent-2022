package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func createTopographyReader(forest [][]int) func(string) [][]int {
	return func(input string) [][]int {
		// create an empty row
		row := []int{}

		// parse the input
		heights := strings.Split(input, "")

		for _, height := range heights {
			intHeight, _ := strconv.Atoi(height)

			// populate the row
			row = append(row, intHeight)
		}

		// add row to the forest
		forest = append(forest, row)

		return forest
	}
}

func max(heights ...int) int {
	max := heights[0]

	for _, height := range heights {
		if height > max {
			max = height
		}
	}

	return max
}

func getCol(col int, forest [][]int) []int {
	column := []int{}

	for _, row := range forest {
		column = append(column, row[col])
	}

	return column
}

func puzzle01(forest [][]int) int {
	// map of rows and columns
	visible := map[int]map[int]bool{}

	for r, trees := range forest {
		visible[r] = map[int]bool{}

		for c, height := range trees {
			// is first row
			if r == 0 {
				visible[r][c] = true
				continue
			}

			// is last row
			if r == len(forest)-1 {
				visible[r][c] = true
				continue
			}

			// is first column
			if c == 0 {
				visible[r][c] = true
				continue
			}

			// is last column
			if c == len(trees)-1 {
				visible[r][c] = true
				continue
			}

			// is taller than those before
			if height > max(trees[0:c]...) {
				visible[r][c] = true
				continue
			}

			// is taller than those after
			if height > max(trees[c+1:]...) {
				visible[r][c] = true
				continue
			}

			column := getCol(c, forest)

			// is taller than those before
			if height > max(column[0:r]...) {
				visible[r][c] = true
				continue
			}

			// is taller than those after
			if height > max(column[r+1:]...) {
				visible[r][c] = true
				continue
			}
		}
	}

	total := 0

	for _, row := range visible {
		for range row {
			total++
		}
	}

	return total

}

func getViewingScore(row int, col int, forest [][]int) int {
	tree := forest[row][col]

	// exclude the bounderies
	if row == 0 ||
		col == 0 ||
		row == len(forest)-1 ||
		col == len(forest[row])-1 {
		return 0
	}

	up := 1
	left := 1
	down := 1
	right := 1

	for r := row - 1; r > 0; r-- {
		if tree > forest[r][col] {
			up++
		} else {
			break
		}
	}

	for r := row + 1; r < len(forest)-1; r++ {
		if tree > forest[r][col] {
			down++
		} else {
			break
		}
	}

	for c := col - 1; c > 0; c-- {
		if tree > forest[row][c] {
			left++
		} else {
			break
		}
	}

	for c := col + 1; c < len(forest[row])-1; c++ {
		if tree > forest[row][c] {
			right++
		} else {
			break
		}
	}

	return up * left * down * right
}

func puzzle02(forest [][]int) int {
	maxScenicScore := 0

	for r, trees := range forest {
		for c := range trees {
			score := getViewingScore(r, c, forest)

			if score > maxScenicScore {
				maxScenicScore = score
			}
		}
	}

	return maxScenicScore
}

func main() {
	input, err := os.Open(os.Args[1])

	if err != nil {
		panic(err)
	}

	defer input.Close()
	scanner := bufio.NewScanner(input)

	var forest [][]int
	sample := createTopographyReader([][]int{})

	for scanner.Scan() {
		forest = sample(scanner.Text())
	}

	fmt.Println(puzzle01(forest))
	fmt.Println(puzzle02(forest))

}
