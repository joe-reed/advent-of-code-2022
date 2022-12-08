package main

import (
	"strings"
	. "utils"

	"github.com/samber/lo"
)

func puzzle1(input []string) (result int) {
	grid := getGrid(input)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if isVisible(grid, i, j) {
				result++
			}
		}
	}
	return
}

func puzzle2(input []string) int {
	grid := getGrid(input)

	scenicScores := []int{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			scenicScores = append(scenicScores, scenicScore(grid, i, j))
		}
	}
	return lo.Max(scenicScores)
}

func getGrid(input []string) [][]int {
	return lo.Map(input, func(row string, i int) []int {
		return MapToInts(strings.Split(row, ""))
	})
}

func isVisible(grid [][]int, i, j int) bool {
	return isOnEdge(grid, i, j) ||
		isVisibleFromLeft(grid, i, j) ||
		isVisibleFromRight(grid, i, j) ||
		isVisibleFromTop(grid, i, j) ||
		isVisibleFromBottom(grid, i, j)
}

func isVisibleFromLeft(grid [][]int, i, j int) bool {
	for k := j - 1; k >= 0; k-- {
		if grid[i][k] >= grid[i][j] {
			return false
		}
	}
	return true
}

func isVisibleFromRight(grid [][]int, i, j int) bool {
	for k := j + 1; k < len(grid); k++ {
		if grid[i][k] >= grid[i][j] {
			return false
		}
	}
	return true
}

func isVisibleFromTop(grid [][]int, i, j int) bool {
	for k := i - 1; k >= 0; k-- {
		if grid[k][j] >= grid[i][j] {
			return false
		}
	}
	return true
}

func isVisibleFromBottom(grid [][]int, i, j int) bool {
	for k := i + 1; k < len(grid); k++ {
		if grid[k][j] >= grid[i][j] {
			return false
		}
	}
	return true
}

func scenicScore(grid [][]int, i, j int) int {
	if isOnEdge(grid, i, j) {
		return 0
	}

	return viewingDistanceLeft(grid, i, j) *
		viewingDistanceRight(grid, i, j) *
		viewingDistanceTop(grid, i, j) *
		viewingDistanceBottom(grid, i, j)
}

func viewingDistanceLeft(grid [][]int, i, j int) int {
	k := j - 1
	for k > 0 {
		if grid[i][k] >= grid[i][j] {
			break
		}
		k--
	}
	return j - k
}

func viewingDistanceRight(grid [][]int, i, j int) int {
	k := j + 1
	for k < len(grid)-1 {
		if grid[i][k] >= grid[i][j] {
			break
		}
		k++
	}
	return k - j
}

func viewingDistanceTop(grid [][]int, i, j int) int {
	k := i - 1
	for k > 0 {
		if grid[k][j] >= grid[i][j] {
			break
		}
		k--
	}
	return i - k
}

func viewingDistanceBottom(grid [][]int, i, j int) int {
	k := i + 1
	for k < len(grid)-1 {
		if grid[k][j] >= grid[i][j] {
			break
		}
		k++
	}
	return k - i
}

func isOnEdge(grid [][]int, i, j int) bool {
	return i == 0 || j == 0 || i == len(grid)-1 || j == len(grid)-1
}
