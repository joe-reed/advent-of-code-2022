package main

import (
	"regexp"
	"sort"
	. "utils"

	"github.com/samber/lo"
)

const (
	DIR  = "dir"
	FILE = "file"
)

func puzzle1(input []string) (result int) {
	return lo.Sum(lo.Filter(getDirSizes(input), func(size int, i int) bool {
		return size < 100000
	}))
}

func puzzle2(input []string) (result int) {
	sizes := getDirSizes(input)
	sort.Ints(sizes)

	unusedSpace := (70000000 - sizes[len(sizes)-1])
	requiredSpace := 30000000 - unusedSpace

	size, _ := lo.Find(sizes, func(size int) bool {
		return size >= requiredSpace
	})

	return size
}

func getDirSizes(input []string) []int {
	currentNode := &Node{}
	dirs := []*Node{}
	for _, line := range input {
		if currentNode.nodeType == DIR && lo.IndexOf(dirs, currentNode) == -1 {
			dirs = append(dirs, currentNode)
		}

		dir, isCd := parseCd(line)
		if isCd {
			currentNode = currentNode.cd(dir)
			continue
		}

		if isLs(line) {
			continue
		}

		file, isFile := parseFile(line)
		if isFile {
			currentNode.addChild(file)
		}
	}

	return lo.Map(dirs, func(dir *Node, i int) int {
		return dir.calculateSize()
	})
}

func parseCd(line string) (string, bool) {
	matches := regexp.MustCompile(`\$ cd (.+)`).FindStringSubmatch(line)
	if len(matches) == 0 {
		return "", false
	}

	return matches[1], true
}

func isLs(line string) bool {
	return regexp.MustCompile(`\$ ls`).MatchString(line)
}

func parseFile(line string) (*Node, bool) {
	matches := regexp.MustCompile(`(\d+) (.+)`).FindStringSubmatch(line)
	if len(matches) == 0 {
		return nil, false
	}

	return &Node{nodeType: FILE, name: matches[2], size: ConvertToInt(matches[1])}, true
}

type Node struct {
	nodeType string
	name     string
	parent   *Node
	children []*Node
	size     int
}

func (n *Node) cd(dir string) *Node {
	if dir == ".." {
		return n.parent
	}
	return n.addChild(&Node{nodeType: DIR, name: dir})
}

func (n *Node) addChild(c *Node) *Node {
	c.parent = n
	n.children = append(n.children, c)
	return c
}

func (n *Node) calculateSize() int {
	if n.size == 0 {
		n.size = lo.Sum(lo.Map(n.children, func(c *Node, i int) int {
			return c.calculateSize()
		}))
	}
	return n.size
}
