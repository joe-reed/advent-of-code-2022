package main

import (
	"strings"

	"github.com/samber/lo"
)

var losingMatchups = map[string]string{
	"A": "C",
	"B": "A",
	"C": "B",
}

var winningMatchups = lo.Invert(losingMatchups)

var xyzAbcMap = map[string]string{
	"X": "A",
	"Y": "B",
	"Z": "C",
}

func puzzle1(input []string) int {
	return playRounds(getRounds(input))
}

func puzzle2(input []string) int {
	return playRounds(mapOutcomesToRounds(input))
}

func getRounds(input []string) []Round {
	return lo.Map(input, func(line string, index int) Round {
		split := strings.Split(line, " ")
		return Round{you: split[0], me: xyzAbcMap[split[1]]}
	})
}

func playRounds(rounds []Round) int {
	scores := lo.Map(rounds, func(round Round, index int) int {
		return round.play()
	})

	return lo.Sum(scores)
}

func mapOutcomesToRounds(input []string) []Round {
	return lo.Map(input, func(line string, index int) Round {
		split := strings.Split(line, " ")

		outcome := split[1]
		me := split[0]
		you := split[0]

		if outcome == "X" {
			me = losingMatchups[you]
		} else if outcome == "Z" {
			me = winningMatchups[you]
		}

		return Round{you: you, me: me}
	})
}

type Round struct {
	you string
	me  string
}

func (r Round) play() int {
	return r.outcomeScore() + r.shapeScore()
}

func (r Round) outcomeScore() int {
	if r.you == r.me {
		return 3
	}
	if winningMatchups[r.you] == r.me {
		return 6
	}
	return 0
}

func (r Round) shapeScore() int {
	return map[string]int{"A": 1, "B": 2, "C": 3}[r.me]
}
