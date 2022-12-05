package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, err := os.Open(os.Args[1])

	if err != nil {
		panic(err)
	}

	defer input.Close()
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		fmt.Println(scanner.Scan())
	}

}
