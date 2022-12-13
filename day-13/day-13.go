package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/samber/lo"
)

const (
	LIST = "list"
	INT  = "int"
)

func puzzle1(input string) (result int) {
	pairs := lo.Map(strings.Split(input, "\n\n"), func(pairString string, i int) Pair {
		return getPair(pairString)
	})

	for i, pair := range pairs {
		if pair.isInRightOrder() {
			result += (i + 1)
		}
	}

	return
}

func puzzle2(input string) int {
	items := lo.FlatMap(strings.Split(input, "\n\n"), func(pairString string, i int) []ListItem {
		pair := getPair(pairString)
		return []ListItem{pair.a, pair.b}
	})
	items = append(items, getListItem(unmarshalJson("[[2]]")))
	items = append(items, getListItem(unmarshalJson("[[6]]")))

	sort.Slice(items, func(i, j int) bool {
		return items[i].isSmallerThan(items[j])
	})

	result := 1
	lo.ForEach(items, func(item ListItem, i int) {
		itemString := fmt.Sprint(item)
		if itemString == "{list 0 [{list 0 [{int 6 []}]}]}" || itemString == "{list 0 [{list 0 [{int 2 []}]}]}" {
			result *= (i + 1)
		}
	})

	return result
}

func getPair(s string) Pair {
	packetStrings := strings.Split(s, "\n")
	return Pair{
		getListItem(unmarshalJson(packetStrings[0])),
		getListItem(unmarshalJson(packetStrings[1])),
	}
}

func unmarshalJson(s string) []interface{} {
	var arr []interface{}
	json.Unmarshal([]byte(s), &arr)
	return arr
}

func getListItem(arr []interface{}) ListItem {
	data := []ListItem{}
	for i := range arr {
		switch item := arr[i].(type) {
		case []interface{}:
			data = append(data, getListItem(item))
		case float64:
			data = append(data, ListItem{itemType: INT, i: int(item)})
		default:
			panic("unrecognised item type")
		}
	}
	return ListItem{itemType: LIST, data: data}
}

type Pair struct {
	a, b ListItem
}

func (p Pair) isInRightOrder() bool {
	return p.a.isSmallerThan(p.b)
}

type ListItem struct {
	itemType string
	i        int
	data     []ListItem
}

func (l ListItem) isSmallerThan(r ListItem) bool {
	if l.itemType == LIST && r.itemType == LIST {
		for i := range l.data {
			if i > len(r.data)-1 {
				return false
			}
			if l.data[i].isSmallerThan(r.data[i]) {
				return true
			}
			if r.data[i].isSmallerThan(l.data[i]) {
				return false
			}
		}
		return len(l.data) < len(r.data)
	}

	if l.itemType == INT && r.itemType == LIST {
		return ListItem{itemType: LIST, data: []ListItem{l}}.isSmallerThan(r)
	}

	if l.itemType == LIST && r.itemType == INT {
		return l.isSmallerThan(ListItem{itemType: LIST, data: []ListItem{r}})
	}

	return l.i < r.i
}
