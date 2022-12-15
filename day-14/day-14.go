package main

import (
	"strings"
	. "utils"

	"github.com/samber/lo"
)

func puzzle1(input string) int {
	grid, maxY := getGrid(input)

	return simulateSand(grid, func(sand Coordinate) bool {
		return sand.y >= maxY
	})
}

func puzzle2(input string) int {
	grid, maxY := getGrid(input)

	for i := 0; i < len(grid[0]); i++ {
		grid[maxY+2][i] = true
	}

	return simulateSand(grid, func(sand Coordinate) bool {
		return sand.y == 0
	})
}

func getGrid(input string) ([][]bool, int) {
	grid := lo.Map(make([][]bool, 200), func(row []bool, i int) []bool {
		return make([]bool, 1000)
	})

	var maxY int
	lo.ForEach(strings.Split(input, "\n"), func(line string, i int) {
		path := getPath(line)
		lo.ForEach(path[1:], func(c Coordinate, i int) {
			prev := path[i]
			if c.y > maxY {
				maxY = c.y
			}
			if c.x > prev.x {
				for j := prev.x; j <= c.x; j++ {
					grid[c.y][j] = true
				}
			}
			if c.x < prev.x {
				for j := c.x; j <= prev.x; j++ {
					grid[c.y][j] = true
				}
			}
			if c.y > prev.y {
				for j := prev.y; j <= c.y; j++ {
					grid[j][c.x] = true
				}
			}
			if c.y < prev.y {
				for j := c.y; j <= prev.y; j++ {
					grid[j][c.x] = true
				}
			}
		})
	})

	return grid, maxY
}

func getPath(line string) []Coordinate {
	return lo.Map(strings.Split(line, " -> "), func(c string, i int) Coordinate {
		split := strings.Split(c, ",")
		return Coordinate{ConvertToInt(split[0]), ConvertToInt(split[1])}
	})
}

func simulateSand(grid [][]bool, terminate func(Coordinate) bool) (result int) {
	for {
		sand := Coordinate{500, 0}
		for {
			if sand.y == len(grid)-1 {
				break
			}
			if !grid[sand.y+1][sand.x] {
				sand = Coordinate{sand.x, sand.y + 1}
				continue
			}
			if !grid[sand.y+1][sand.x-1] {
				sand = Coordinate{sand.x - 1, sand.y + 1}
				continue
			}
			if !grid[sand.y+1][sand.x+1] {
				sand = Coordinate{sand.x + 1, sand.y + 1}
				continue
			}

			result++

			grid[sand.y][sand.x] = true
			break
		}
		if terminate(sand) {
			break
		}
	}
	return
}

type Coordinate struct {
	x, y int
}
