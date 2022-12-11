package main

import (
	"regexp"
	"sort"
	"strings"
	. "utils"

	"github.com/samber/lo"
)

func puzzle1(input string) (result int) {
	return solve(input, func(i int, monkeys []*Monkey) int {
		return i / 3
	}, 20)
}

func puzzle2(input string) (result int) {
	return solve(input, func(i int, monkeys []*Monkey) int {
		return i % lo.Reduce(monkeys, func(carry int, m *Monkey, i int) int {
			return carry * m.test
		}, 1)
	}, 10000)
}

func solve(input string, adjustForWorry func(int, []*Monkey) int, rounds int) int {
	monkeys := getMonkeys(input)

	for i := 0; i < rounds; i++ {
		for _, m := range monkeys {
			for len(m.items) > 0 {
				m.inspectionCount++
				item := m.operation(m.items[0])
				item = adjustForWorry(item, monkeys)
				m.throwTo(monkeys[m.getRecipient(item)], item)
			}
		}
	}

	return calculateMonkeyBusiness(monkeys)
}

func getMonkeys(input string) []*Monkey {
	return lo.Map(strings.Split(input, "\n\n"), func(monkeyString string, i int) *Monkey {
		split := strings.Split(monkeyString, "\n")
		return &Monkey{
			items:           getItems(split[1]),
			operation:       getOperation(split[2]),
			test:            getNumber(split[3]),
			trueMonkey:      getNumber(split[4]),
			falseMonkey:     getNumber(split[5]),
			inspectionCount: 0,
		}
	})
}

func getItems(s string) []int {
	matches := regexp.MustCompile(`Starting items: (.+)`).FindStringSubmatch(s)
	return MapToInts(strings.Split(matches[1], ", "))
}

func getOperation(s string) func(int) int {
	matches := regexp.MustCompile(`new = (.+) (.) (.+)`).FindStringSubmatch(s)

	return func(old int) int {
		a, b := old, old
		if matches[1] != "old" {
			a = ConvertToInt(matches[1])
		}
		if matches[3] != "old" {
			b = ConvertToInt(matches[3])
		}

		switch matches[2] {
		case "+":
			return a + b
		case "*":
			return a * b
		default:
			panic("unrecognised operation")
		}
	}
}

func getNumber(s string) int {
	return ConvertToInt(regexp.MustCompile(`(\d+)`).FindStringSubmatch(s)[1])
}

func calculateMonkeyBusiness(monkeys []*Monkey) int {
	sort.Slice(monkeys, func(a, b int) bool {
		return monkeys[a].inspectionCount > monkeys[b].inspectionCount
	})
	return monkeys[0].inspectionCount * monkeys[1].inspectionCount
}

type Monkey struct {
	items                   []int
	operation               func(int) int
	test                    int
	trueMonkey, falseMonkey int
	inspectionCount         int
}

func (m *Monkey) getRecipient(item int) int {
	if item%m.test == 0 {
		return m.trueMonkey
	}
	return m.falseMonkey
}

func (a *Monkey) throwTo(b *Monkey, item int) {
	a.items = a.items[1:]
	b.items = append(b.items, item)
}
