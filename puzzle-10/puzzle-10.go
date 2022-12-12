package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cpu struct {
	cycle     int
	registers map[string]int
	history   []int
}

func (cpu *cpu) tick() {
	cpu.cycle++
	cpu.history = append(cpu.history, cpu.getSignalStrength())
}

func (cpu *cpu) addx(value int) {
	cpu.tick()
	cpu.tick()
	cpu.registers["X"] += value
}

func (cpu *cpu) noop() {
	cpu.tick()
}

func (cpu *cpu) exec(instr instruction) {
	switch instr.operation {
	case "addx":
		cpu.addx(instr.argument)
	case "noop":
		cpu.noop()
	}
}

func (cpu *cpu) getSignalStrength() int {
	return cpu.cycle * cpu.registers["X"]
}

type instruction struct {
	operation string
	argument  int
}

func puzzle01(program *[]instruction) int {
	cpu := cpu{0, map[string]int{"X": 1}, []int{}}

	for _, instruction := range *program {
		cpu.exec(instruction)
	}

	total := 0

	for s, strength := range cpu.history {
		cycle := s + 1

		if (cycle-20)%40 == 0 {
			fmt.Println(cycle, strength)
			total += strength
		}
	}

	return total
}

func main() {
	input, err := os.Open(os.Args[1])

	if err != nil {
		panic(err)
	}

	defer input.Close()
	scanner := bufio.NewScanner(input)

	program := []instruction{}

	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		var argument int

		if len(tokens) > 1 {
			arg, _ := strconv.Atoi(tokens[1])
			argument = arg
		}

		program = append(program, instruction{tokens[0], argument})
	}

	fmt.Println(puzzle01(&program))
}
