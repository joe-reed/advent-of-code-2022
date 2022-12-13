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
			7,
		},
		{
			"./test-input-2.txt",
			5,
		},
		{
			"./test-input-3.txt",
			6,
		},
		{
			"./test-input-4.txt",
			10,
		},
		{
			"./test-input-5.txt",
			11,
		},
		{
			"./input.txt",
			1760,
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
			19,
		},
		{
			"./test-input-2.txt",
			23,
		},
		{
			"./test-input-3.txt",
			23,
		},
		{
			"./test-input-4.txt",
			29,
		},
		{
			"./test-input-5.txt",
			26,
		},
		{
			"./input.txt",
			2974,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, puzzle2(FileToString(test.input)))
	}
}

func TestSolvePuzzle2(t *testing.T) {
	fmt.Println("Puzzle 2:", puzzle2(FileToString("./input.txt")))
}
