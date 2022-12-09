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
	s := Knot{0, 0}

	rope := lo.Map(lo.Range(length), func(i, j int) Knot {
		return s
	})

	instructions := lo.Map(input, func(line string, i int) Instruction {
		split := strings.Split(line, " ")
		return Instruction{split[0], ConvertToInt(split[1])}
	})

	visited := map[Knot]bool{s: true}

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

type Knot struct {
	x, y int
}

func (k Knot) move(dir string) Knot {
	switch dir {
	case "U":
		return Knot{k.x, k.y + 1}
	case "R":
		return Knot{k.x + 1, k.y}
	case "D":
		return Knot{k.x, k.y - 1}
	case "L":
		return Knot{k.x - 1, k.y}
	default:
		panic("unrecognised direction")
	}
}

func (f Knot) follow(k Knot) Knot {
	if f.isAdjacentTo(k) {
		return f
	}
	if k.y > f.y {
		f = f.move("U")
	}
	if k.x < f.x {
		f = f.move("L")
	}
	if k.y < f.y {
		f = f.move("D")
	}
	if k.x > f.x {
		f = f.move("R")
	}
	return f
}

func (a Knot) isAdjacentTo(b Knot) bool {
	return math.Abs(float64(a.x-b.x)) <= 1 && math.Abs(float64(a.y-b.y)) <= 1
}

type Instruction struct {
	dir  string
	dist int
}
