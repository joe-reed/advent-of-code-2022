package main

import (
	"regexp"
	. "utils"

	"github.com/samber/lo"
)

func puzzle1(input []string, rowNumber int) int {
	row := make(map[int]bool)

	sensors := getSensors(input)
	for _, s := range sensors {
		if abs(s.position.y-rowNumber) <= s.distance {
			for i := 0; i <= s.distance-abs(s.position.y-rowNumber); i++ {
				row[s.position.x+i] = true
				row[s.position.x-i] = true
			}
		}
		if s.beacon.y == rowNumber {
			row[s.beacon.x] = false
		}
	}
	return lo.Count(lo.Values(row), true)
}

func puzzle2(input []string, size int) int {
	justOutOfReach := lo.RepeatBy(size, func(i int) map[int]int {
		return make(map[int]int)
	})

	lo.ForEach(getSensors(input), func(s Sensor, _ int) {
		for i := -(s.distance + 1); i <= s.distance+1; i++ {
			d := abs((s.distance + 1) - abs(i))
			if 0 < s.position.y+i && s.position.y+i < size && 0 < s.position.x+d && s.position.x+d < size {
				justOutOfReach[s.position.y+i][s.position.x+d]++
				if d != 0 {
					justOutOfReach[s.position.y+i][s.position.x-d]++
				}
			}
		}
	})

	max := 0
	yMax := 0
	for y, row := range justOutOfReach {
		m := lo.Max(lo.Values(row))
		if m > max {
			max = m
			yMax = y
		}
	}

	xMax, _ := lo.FindKey(justOutOfReach[yMax], max)
	return yMax + xMax*4000000
}

func getSensors(input []string) []Sensor {
	return lo.Map(input, func(line string, i int) Sensor {
		re := regexp.MustCompile(`Sensor at x=(.+), y=(.+): closest beacon is at x=(.+), y=(.+)`)
		matches := re.FindStringSubmatch(line)
		ints := MapToInts(matches[1:])
		distance := abs(ints[0]-ints[2]) + abs(ints[1]-ints[3])
		return Sensor{Position{ints[0], ints[1]}, Position{ints[2], ints[3]}, distance}
	})
}

type Sensor struct {
	position, beacon Position
	distance         int
}

type Position struct {
	x, y int
}

func abs(i int) int {
	if i > 0 {
		return i
	}
	return -i
}
