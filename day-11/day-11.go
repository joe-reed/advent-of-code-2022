package main

import (
	"regexp"
	"sort"
	"strings"
	. "utils"

	"github.com/samber/lo"
)

func puzzle1(input string) (result int) {
	monkeys := getMonkeys(input)

	for i := 0; i < 20; i++ {
		for _, monkey := range monkeys {
			for len(monkey.items) > 0 {
				monkey.inspectionCount++
				newItem := monkey.operation(monkey.items[0])
				newItem /= 3
				if newItem%monkey.test == 0 {
					monkey.throwTo(monkeys[monkey.trueMonkey], newItem)
				} else {
					monkey.throwTo(monkeys[monkey.falseMonkey], newItem)
				}
			}
		}
	}

	return calculateMonkeyBusiness(monkeys)
}

func puzzle2(input string) (result int) {
	monkeys := getMonkeys(input)

	mod := lo.Reduce(monkeys, func(carry int, m *Monkey, i int) int {
		return carry * m.test
	}, 1)

	for i := 0; i < 10000; i++ {
		for _, monkey := range monkeys {
			for len(monkey.items) > 0 {
				monkey.inspectionCount++
				newItem := monkey.operation(monkey.items[0])
				if newItem%monkey.test == 0 {
					monkey.throwTo(monkeys[monkey.trueMonkey], newItem%mod)
				} else {
					monkey.throwTo(monkeys[monkey.falseMonkey], newItem%mod)
				}
			}
		}
	}

	return calculateMonkeyBusiness(monkeys)
}

func getMonkeys(input string) []*Monkey {
	return lo.Map(strings.Split(input, "\n\n"), func(monkeyString string, i int) *Monkey {
		split := strings.Split(monkeyString, "\n")

		items := getItems(split[1])
		operation := getOperation(split[2])
		test := getTest(split[3])
		trueMonkey := getTrueMonkey(split[4])
		falseMonkey := getFalseMonkey(split[5])
		return &Monkey{items, operation, test, trueMonkey, falseMonkey, 0}
	})
}

func getItems(s string) []int {
	matches := regexp.MustCompile(`Starting items: (.+)`).FindStringSubmatch(s)
	return MapToInts(strings.Split(matches[1], ", "))
}

func getOperation(s string) func(int) int {
	matches := regexp.MustCompile(`Operation: new = (.+) (.) (.+)`).FindStringSubmatch(s)

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

func getTest(s string) int {
	matches := regexp.MustCompile(`Test: divisible by (\d+)`).FindStringSubmatch(s)
	return ConvertToInt(matches[1])
}

func getTrueMonkey(s string) int {
	matches := regexp.MustCompile(`If true: throw to monkey (\d+)`).FindStringSubmatch(s)
	return ConvertToInt(matches[1])
}

func getFalseMonkey(s string) int {
	matches := regexp.MustCompile(`If false: throw to monkey (\d+)`).FindStringSubmatch(s)
	return ConvertToInt(matches[1])
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

func (a *Monkey) throwTo(b *Monkey, item int) {
	items := make([]int, len(a.items)-1)
	copy(items, a.items[1:])
	a.items = items
	b.items = append(b.items, item)
}
