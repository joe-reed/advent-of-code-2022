package main

import (
	"strings"
	. "utils"

	"github.com/samber/lo"
)

func puzzle1(input []string) (result int) {
	runCycles(getInstructions(input), func(instruction Instruction, x, c int) {
		if (c+1)%40 == 20 {
			result += (c + 1) * x
		}
	})

	return
}

func puzzle2(input []string) (result string) {
	runCycles(getInstructions(input), func(instruction Instruction, x, c int) {
		if x-1 <= c%40 && c%40 <= x+1 {
			result += "#"
		} else {
			result += "."
		}

		if c%40 == 39 {
			result += "\n"
		}
	})

	return
}

func getInstructions(input []string) []Instruction {
	return lo.Map(input, func(line string, i int) Instruction {
		if line == "noop" {
			return Instruction{cycleTime: 1}
		}
		return Instruction{cycleTime: 2, value: ConvertToInt(strings.Split(line, " ")[1])}
	})
}

func runCycles(instructions []Instruction, callback func(instruction Instruction, x, c int)) {
	c, x := 0, 1
	for _, instruction := range instructions {
		for i := 1; i <= instruction.cycleTime; i++ {
			callback(instruction, x, c)
			c++
		}
		x += instruction.value
	}
}

type Instruction struct {
	cycleTime, value int
}
