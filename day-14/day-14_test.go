package main

import (
	"fmt"
	"testing"
	. "utils"

	"github.com/stretchr/testify/assert"
)

func TestPuzzle1(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			"./test-input-1.txt",
			24,
		},
		{
			"./input.txt",
			755,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, puzzle1(FileToString(test.input)))
	}
}

func TestSolvePuzzle1(t *testing.T) {
	fmt.Println("Puzzle 1:", puzzle1(FileToString("./input.txt")))
}

func TestPuzzle2(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			"./test-input-1.txt",
			93,
		},
		{
			"./input.txt",
			29805,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, puzzle2(FileToString(test.input)))
	}
}

func TestSolvePuzzle2(t *testing.T) {
	fmt.Println("Puzzle 2:", puzzle2(FileToString("./input.txt")))
}
