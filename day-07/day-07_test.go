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
			95437,
		},
		{
			"./input.txt",
			1453349,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, puzzle1(LoadFile(test.input)))
	}
}

func TestSolvePuzzle1(t *testing.T) {
	fmt.Println("Puzzle 1:", puzzle1(LoadFile("./input.txt")))
}

func TestPuzzle2(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			"./test-input-1.txt",
			24933642,
		},
		{
			"./input.txt",
			2948823,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, puzzle2(LoadFile(test.input)))
	}
}

func TestSolvePuzzle2(t *testing.T) {
	fmt.Println("Puzzle 2:", puzzle2(LoadFile("./input.txt")))
}
