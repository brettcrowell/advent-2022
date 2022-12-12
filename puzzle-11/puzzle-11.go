package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	items       []int  // slice of items
	operator    string // * | +
	value       int    // value for operator
	self        bool   // re-use old instead of value
	modulo      int    // basis for divisble test
	pass        int    // target monkey index when test passes
	fail        int    // target monkey index when test fails
	inspections int    // number of items inspected
}

func (monkey *Monkey) operation(old int) int {
	value := monkey.value

	if monkey.self {
		value = old
	}

	switch monkey.operator {
	case "+":
		{
			new := old + value
			fmt.Println("    Worry level is increased by", value, "to", new)
			return new
		}
	case "*":
		{
			new := old * value
			fmt.Println("    Worry level is multiplied by", value, "to", new)
			return new
		}
	}

	panic("unrecognized operator")
}

func (monkey *Monkey) test(worryLevel int) int {
	remainder := worryLevel % monkey.modulo
	divisible := remainder == 0

	adjective := "is"

	if !divisible {
		adjective = "isnt"
	}

	fmt.Println("    Current worry level", adjective, "divisble by", monkey.modulo)

	if remainder == 0 {
		return monkey.pass
	}
	return monkey.fail
}

func (monkey *Monkey) catch(item int) {
	monkey.items = append(monkey.items, item)
}

func (monkey *Monkey) throw(to *Monkey, item int, divisor int) {
	to.catch(item % divisor)
	monkey.items = monkey.items[1:]
}

func (monkey *Monkey) inspect(
	monkeys []*Monkey,
	worryFactor int,
	divisor int,
) {
	monkey.inspections++

	item := monkey.items[0]
	fmt.Println("  Monkey inspects an item with a worry level of", item)

	worry := monkey.operation(item)

	worry /= worryFactor
	fmt.Println("    Monkey gets bored with item. Worry level is divided by", worryFactor, "to", worry)

	passTo := monkey.test(worry)

	passToMonkey := monkeys[passTo]
	monkey.throw(passToMonkey, worry, divisor)
	fmt.Println("    item with worry level", worry, "passed to monkey", passTo)
}

func puzzle01(monkeys []*Monkey) int {
	for round := 0; round < 20; round++ {
		fmt.Println("==", round, "==")
		for m, monkey := range monkeys {
			fmt.Println("Monkey ", m)
			for range monkey.items {
				monkey.inspect(monkeys, 3, 1)
			}
		}
	}

	inspections := []int{}

	for m, monkey := range monkeys {
		fmt.Println("Monkey", m, "inspected", monkey.inspections, "items")
		inspections = append(inspections, monkey.inspections)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(inspections)))

	return inspections[0] * inspections[1]
}

func puzzle02(monkeys []*Monkey) int {
	// inspired/stolen from @gamache
	divisor := 1
	for _, monkey := range monkeys {
		divisor *= int(monkey.modulo)
	}

	for round := 0; round < 10000; round++ {
		fmt.Println("==", round, "==")
		for m, monkey := range monkeys {
			fmt.Println("Monkey ", m)
			for range monkey.items {
				monkey.inspect(monkeys, 1, divisor)
			}
		}
	}

	inspections := []int{}

	for m, monkey := range monkeys {
		fmt.Println("Monkey", m, "inspected", monkey.inspections, "items")
		inspections = append(inspections, monkey.inspections)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(inspections)))

	return inspections[0] * inspections[1]
}

func main() {
	input, err := os.Open(os.Args[1])

	if err != nil {
		panic(err)
	}

	defer input.Close()
	scanner := bufio.NewScanner(input)

	// input parsers

	parseItems, _ := regexp.Compile(
		"Starting items: ([\\d]+[,\\s(\\d+)]*)",
	)

	parseOperation, _ := regexp.Compile(
		"Operation: new = old ([*+]) (\\d+|old)",
	)

	parseTest, _ := regexp.Compile(
		"Test: divisible by (\\d+)",
	)

	parseOutcome, _ := regexp.Compile(
		"If (true|false): throw to monkey (\\d+)",
	)

	// globals

	monkeys := []*Monkey{}
	iter := 0

	for scanner.Scan() {
		line := scanner.Text()
		row := iter % 7

		if row == 0 {
			newMonkey := Monkey{}
			monkeys = append(monkeys, &newMonkey)
		}

		currentMonkey := monkeys[len(monkeys)-1]

		switch row {
		case 1:
			{
				parsed := parseItems.FindAllStringSubmatch(line, -1)[0][1]
				items := strings.Split(parsed, ", ")

				for _, item := range items {
					currentItem, _ := strconv.Atoi(item)
					currentMonkey.items = append(currentMonkey.items, currentItem)
				}
			}
		case 2:
			{
				operation := parseOperation.FindAllStringSubmatch(line, -1)[0]

				operator := operation[1]
				currentMonkey.operator = operator

				switch operation[2] {
				case "old":
					currentMonkey.self = true
					break
				default:
					{
						value, _ := strconv.Atoi(operation[2])
						currentMonkey.value = value
					}
				}

			}
		case 3:
			{
				test := parseTest.FindAllStringSubmatch(line, -1)[0]
				modulo, _ := strconv.Atoi(test[1])

				currentMonkey.modulo = modulo
			}
		case 4:
			{
				test := parseOutcome.FindAllStringSubmatch(line, -1)[0]
				index, _ := strconv.Atoi(test[2])

				currentMonkey.pass = index
			}
		case 5:
			{
				test := parseOutcome.FindAllStringSubmatch(line, -1)[0]
				index, _ := strconv.Atoi(test[2])

				currentMonkey.fail = index
			}
		}

		iter++
	}

	// fmt.Println(puzzle01(monkeys))
	fmt.Println(puzzle02(monkeys))

}
