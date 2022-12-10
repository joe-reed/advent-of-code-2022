package main

import (
	"strings"
	. "utils"

	"github.com/samber/lo"
)

func puzzle1(input []string) int {
	instructions := getInstructions(input)

	x := 1
	c := 0
	toRecord := []int{20, 60, 100, 140, 180, 220}
	results := []int{}
cycle:
	for _, instruction := range instructions {
		for i := 1; i <= instruction.cycleTime; i++ {
			c++

			if c == toRecord[0] {
				results = append(results, c*x)
				toRecord = toRecord[1:]
			}

			if len(toRecord) == 0 {
				break cycle
			}
		}
		x += instruction.value
	}

	return lo.Sum(results)
}

func puzzle2(input []string) (result string) {
	instructions := getInstructions(input)

	x := 1
	c := 0
	for _, instruction := range instructions {
		for i := 1; i <= instruction.cycleTime; i++ {
			if x-1 <= c%40 && c%40 <= x+1 {
				result += "#"
			} else {
				result += "."
			}

			c++
			if c%40 == 0 {
				result += "\n"
			}
		}
		x += instruction.value
	}

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

type Instruction struct {
	cycleTime, value int
}
