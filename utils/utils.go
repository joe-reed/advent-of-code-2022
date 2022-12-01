package utils

import (
	"os"
	"strconv"
	"strings"

	lop "github.com/samber/lo/parallel"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ConvertToInt(s string) int {
	result, err := strconv.Atoi(s)
	Check(err)
	return result
}

func MapToInts(s []string) []int {
	return lop.Map(s, func(str string, index int) int {
		return ConvertToInt(str)
	})
}

func LoadFile(path string) []string {
	file, err := os.ReadFile(path)
	Check(err)
	lines := strings.Split(string(file), "\n")
	return lines[:len(lines)-1]
}
