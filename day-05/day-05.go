package main

import (
	"regexp"
	"strings"
	. "utils"

	"github.com/samber/lo"
)

func puzzle1(input string) string {
	stacks, instructions := parseInput(input)

	for _, instruction := range instructions {
		instruction.performOneByOne(stacks)
	}

	return getResult(stacks)
}

func puzzle2(input string) string {
	stacks, instructions := parseInput(input)

	for _, instruction := range instructions {
		instruction.performInChunks(stacks)
	}

	return getResult(stacks)
}

func getResult(stacks []*Stack) (result string) {
	for _, stack := range stacks {
		result += stack.items[0]
	}
	return
}

func parseInput(input string) ([]*Stack, []Instruction) {
	split := strings.Split(input, "\n\n")
	return parseStacks(split[0]), parseInstructions(split[1])
}

func parseStacks(input string) []*Stack {
	rows := strings.Split(input, "\n")

	numbersRow := rows[len(rows)-1]

	numbers := MapToInts(strings.Split(strings.ReplaceAll(numbersRow, " ", ""), ""))

	return lo.Map(numbers, func(n int, index int) *Stack {
		s := &Stack{}
		lo.ForEach(rows[:len(rows)-1], func(row string, index int) {
			crate := string(row[4*n-3])

			if crate != " " {
				s.items = append(s.items, crate)
			}
		})
		return s
	})
}

func parseInstructions(input string) []Instruction {
	re := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

	return lo.Map(strings.Split(input, "\n"), func(row string, index int) Instruction {
		matches := re.FindStringSubmatch(row)
		return Instruction{number: ConvertToInt(matches[1]), start: ConvertToInt(matches[2]) - 1, end: ConvertToInt(matches[3]) - 1}
	})
}

type Stack struct {
	items []string
}

func (s *Stack) addCrate(c string) {
	s.addCrates([]string{c})
}

func (s *Stack) removeCrate() string {
	return s.removeCrates(1)[0]
}

func (s *Stack) addCrates(c []string) {
	s.items = append(c, s.items...)
}

func (s *Stack) removeCrates(n int) []string {
	crates := make([]string, n)
	copy(crates, s.items[:n])

	items := make([]string, len(s.items)-n)
	copy(items, s.items[n:])
	s.items = items

	return crates
}

type Instruction struct {
	number int
	start  int
	end    int
}

func (instruction Instruction) performOneByOne(stacks []*Stack) {
	for i := 0; i < instruction.number; i++ {
		startStack := stacks[instruction.start]
		endStack := stacks[instruction.end]

		crate := startStack.removeCrate()
		endStack.addCrate(crate)
	}
}

func (instruction Instruction) performInChunks(stacks []*Stack) {
	startStack := stacks[instruction.start]
	endStack := stacks[instruction.end]

	crates := startStack.removeCrates(instruction.number)
	endStack.addCrates(crates)
}
