package main

import (
	"strings"

	"github.com/samber/lo"
)

var alphabet = []string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
}

func puzzle1(input []string) (result int) {
	grid, start, end := getGrid(input)

	return getShortestPath(grid, start, end)
}

func puzzle2(input []string) (result int) {
	grid, _, end := getGrid(input)

	starts := []Coordinate{}
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == 0 {
				starts = append(starts, Coordinate{x, y})
			}
		}
	}

	distances := lo.Map(starts, func(start Coordinate, i int) int {
		return getShortestPath(grid, start, end)
	})

	return lo.Min(lo.Filter(distances, func(d int, i int) bool {
		return d != -1
	}))
}

func getGrid(input []string) ([][]int, Coordinate, Coordinate) {
	var s Coordinate
	var e Coordinate
	g := lo.Map(input, func(line string, y int) []int {
		return lo.Map(strings.Split(line, ""), func(c string, x int) int {
			if c == "S" {
				s = Coordinate{x, y}
				return 0
			}
			if c == "E" {
				e = Coordinate{x, y}
				return 25
			}
			return lo.IndexOf(alphabet, c)
		})
	})
	return g, s, e
}

func getShortestPath(g [][]int, start Coordinate, end Coordinate) int {
	c := start
	d := 0
	visited := []Coordinate{}
	queue := []Location{{c, d}}
	for len(queue) > 0 {
		l := queue[0]
		c, d, queue = l.c, l.d, queue[1:]
		if c.x == end.x && c.y == end.y {
			return d
		}
		if contains(visited, c) {
			continue
		}
		visited = append(visited, c)
		currentVal := g[c.y][c.x]

		if c.y > 0 && g[c.y-1][c.x]-currentVal <= 1 {
			next := Coordinate{c.x, c.y - 1}
			queue = append(queue, Location{next, d + 1})
		}
		if c.x > 0 && g[c.y][c.x-1]-currentVal <= 1 {
			next := Coordinate{c.x - 1, c.y}
			queue = append(queue, Location{next, d + 1})
		}
		if c.y < len(g)-1 && g[c.y+1][c.x]-currentVal <= 1 {
			next := Coordinate{c.x, c.y + 1}
			queue = append(queue, Location{next, d + 1})
		}
		if c.x < len(g[0])-1 && g[c.y][c.x+1]-currentVal <= 1 {
			next := Coordinate{c.x + 1, c.y}
			queue = append(queue, Location{next, d + 1})
		}
	}
	return -1
}

func contains(haystack []Coordinate, needle Coordinate) bool {
	return lo.ContainsBy(haystack, func(c Coordinate) bool {
		return c.x == needle.x && c.y == needle.y
	})
}

type Coordinate struct {
	x, y int
}

type Location struct {
	c Coordinate
	d int
}
