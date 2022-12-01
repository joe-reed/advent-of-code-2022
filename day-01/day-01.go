package main

import (
	"sort"
	"strings"
	. "utils"

	"github.com/samber/lo"
)

func puzzle1(file string) int {
	sums := getSums(file)
	return lo.Max(sums)
}

func puzzle2(file string) int {
	sums := getSums(file)
	sort.Ints(sums)
	return lo.Sum(sums[len(sums)-3:])
}

func getSums(file string) []int {
	lists := strings.Split(file, "\n\n")

	elves := lo.Map(lists, func(str string, index int) []int {
		strs := strings.Split(str, "\n")
		return MapToInts(strs)
	})

	return lo.Map(elves, func(items []int, index int) int {
		return lo.Sum(items)
	})
}
