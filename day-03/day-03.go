package main

import (
	"strings"

	"github.com/samber/lo"
)

var itemList = []string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
}

func puzzle1(input []string) int {
	priorities := lo.Map(input, func(bag string, index int) int {
		items := strings.Split(bag, "")
		compartment1, compartment2 := items[:len(items)/2], items[len(items)/2:]
		duplicateItem := lo.Intersect(compartment1, compartment2)[0]
		return lo.IndexOf(itemList, duplicateItem) + 1
	})

	return lo.Sum(priorities)
}

func puzzle2(input []string) int {
	teams := lo.Chunk(input, 3)

	priorities := lo.Map(teams, func(team []string, index int) int {
		bags := lo.Map(team, func(bag string, index int) []string {
			return strings.Split(bag, "")
		})
		intersection := lo.Intersect(lo.Intersect(bags[0], bags[1]), bags[2])[0]
		return lo.IndexOf(itemList, intersection) + 1
	})

	return lo.Sum(priorities)
}
