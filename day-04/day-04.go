package main

import (
	"strings"
	. "utils"

	"github.com/samber/lo"
)

func puzzle1(input []string) int {
	pairs := getPairs(input)

	return len(lo.Filter(pairs, func(pair Pair, index int) bool {
		return pair.fullyIntersects()
	}))
}

func puzzle2(input []string) int {
	pairs := getPairs(input)

	return len(lo.Filter(pairs, func(pair Pair, index int) bool {
		return pair.intersects()
	}))
}

func getPairs(input []string) []Pair {
	return lo.Map(input, func(pairString string, index int) Pair {
		assignmentStrings := strings.Split(pairString, ",")

		return Pair{a1: newAssignment(assignmentStrings[0]), a2: newAssignment(assignmentStrings[1])}
	})
}

func newAssignment(assignmentString string) Assignment {
	ints := MapToInts(strings.Split(assignmentString, "-"))

	return Assignment{start: ints[0], end: ints[1]}
}

type Pair struct {
	a1 Assignment
	a2 Assignment
}

func (p Pair) fullyIntersects() bool {
	return len(p.intersection()) == p.a1.length() || len(p.intersection()) == p.a2.length()
}

func (p Pair) intersects() bool {
	return len(p.intersection()) > 0
}

func (p Pair) intersection() []int {
	return lo.Intersect(p.a1.section(), p.a2.section())
}

type Assignment struct {
	start int
	end   int
}

func (a Assignment) section() []int {
	return lo.RangeFrom(a.start, a.length())
}

func (a Assignment) length() int {
	return a.end - a.start + 1
}
