package main

import (
	"strings"

	"github.com/samber/lo"
)

func puzzle1(input string) int {
	return solve(input, 4)
}

func puzzle2(input string) int {
	return solve(input, 14)
}

func solve(input string, length int) int {
	chars := strings.Split(input, "")

	for i := length; i < len(chars)-1; i++ {
		if len(lo.FindDuplicates(chars[i-length:i])) == 0 {
			return i
		}
	}
	return -1
}
