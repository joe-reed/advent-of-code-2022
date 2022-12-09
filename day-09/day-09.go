package main

import (
	"math"
	"strings"
	. "utils"

	"github.com/samber/lo"
)

func puzzle1(input []string) (result int) {
	return solve(input, 2)
}

func puzzle2(input []string) (result int) {
	return solve(input, 10)
}

func solve(input []string, length int) int {
	s := Location{0, 0}

	rope := lo.Map(lo.Range(length), func(i, j int) Location {
		return s
	})

	instructions := lo.Map(input, func(line string, i int) Instruction {
		split := strings.Split(line, " ")
		return Instruction{split[0], ConvertToInt(split[1])}
	})

	visited := map[Location]bool{s: true}

	for _, instr := range instructions {
		for range lo.Range(instr.dist) {
			for i, knot := range rope {
				if i == 0 {
					rope[0] = knot.move(instr.dir)
					continue
				}
				rope[i] = knot.follow(rope[i-1])
			}
			visited[rope[length-1]] = true
		}
	}

	return len(visited)
}

type Location struct {
	x, y int
}

func (l Location) move(dir string) Location {
	switch dir {
	case "U":
		return Location{l.x, l.y + 1}
	case "R":
		return Location{l.x + 1, l.y}
	case "D":
		return Location{l.x, l.y - 1}
	case "L":
		return Location{l.x - 1, l.y}
	case "UR":
		return Location{l.x + 1, l.y + 1}
	case "UL":
		return Location{l.x - 1, l.y + 1}
	case "DR":
		return Location{l.x + 1, l.y - 1}
	case "DL":
		return Location{l.x - 1, l.y - 1}
	default:
		panic("unrecognised direction")
	}
}

func (f Location) follow(l Location) Location {
	if f.isAdjacentTo(l) {
		return f
	}
	if l.y > f.y && l.x > f.x {
		return f.move("UR")
	}
	if l.y > f.y && l.x < f.x {
		return f.move("UL")
	}
	if l.y < f.y && l.x > f.x {
		return f.move("DR")
	}
	if l.y < f.y && l.x < f.x {
		return f.move("DL")
	}
	if l.y > f.y {
		return f.move("U")
	}
	if l.y < f.y {
		return f.move("D")
	}
	if l.x > f.x {
		return f.move("R")
	}
	return f.move("L")
}

func (a Location) isAdjacentTo(b Location) bool {
	return math.Abs(float64(a.x-b.x)) <= 1 && math.Abs(float64(a.y-b.y)) <= 1
}

type Instruction struct {
	dir  string
	dist int
}
