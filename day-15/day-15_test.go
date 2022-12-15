package main

import (
	"fmt"
	"testing"
	. "utils"

	"github.com/stretchr/testify/assert"
)

func TestPuzzle1(t *testing.T) {
	tests := []struct {
		input     string
		rowNumber int
		expected  int
	}{
		{
			"./test-input-1.txt",
			10,
			26,
		},
		{
			"./input.txt",
			2000000,
			5394423,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, puzzle1(FileToStrings(test.input), test.rowNumber))
	}
}

func TestSolvePuzzle1(t *testing.T) {
	fmt.Println("Puzzle 1:", puzzle1(FileToStrings("./input.txt"), 2000000))
}

func TestPuzzle2(t *testing.T) {
	tests := []struct {
		input    string
		size     int
		expected int
	}{
		{
			"./test-input-1.txt",
			20,
			56000011,
		},
		{
			"./input.txt",
			4000000,
			11840879211051,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, puzzle2(FileToStrings(test.input), test.size))
	}
}

func TestSolvePuzzle2(t *testing.T) {
	fmt.Println("Puzzle 2:", puzzle2(FileToStrings("./input.txt"), 4000000))
}
