package main

import (
	"fmt"
	"os"
	"testing"
	. "utils"

	"github.com/stretchr/testify/assert"
)

func TestPuzzle1(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"./test-input-1.txt",
			"CMZ",
		},
		{
			"./input.txt",
			"SVFDLGLWV",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, puzzle1(openFile(test.input)))
	}
}

func TestSolvePuzzle1(t *testing.T) {
	fmt.Println("Puzzle 1:", puzzle1(openFile("./input.txt")))
}

func TestPuzzle2(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"./test-input-1.txt",
			"MCD",
		},
		{
			"./input.txt",
			"DCVTCVPCL",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, puzzle2(openFile(test.input)))
	}
}

func TestSolvePuzzle2(t *testing.T) {
	fmt.Println("Puzzle 2:", puzzle2(openFile("./input.txt")))
}

func openFile(path string) string {
	file, err := os.ReadFile(path)
	Check(err)
	return string(file)
}
