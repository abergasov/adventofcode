package main

import (
	"adventofcode/utils"
	"log"
	"sort"
)

var cache map[int]int

func main() {
	voltageData, err := utils.GetData("10")
	if err != nil {
		log.Fatal(err)
	}
	voltage := utils.Convert2Int(voltageData)
	sort.Ints(voltage)

	voltage = append(voltage, voltage[len(voltage)-1]+3)
	cache = make(map[int]int, len(voltage))

	currentGroup := []int{0}

	diffThree := 0
	diffOne := 0
	prev := 0
	pt2 := 1
	for _, i := range voltage {
		diff := i - prev
		if diff == 1 {
			diffOne++
		} else if diff == 3 {
			diffThree++
			pt2 *= countGroupCombinations(currentGroup)
			currentGroup = nil
		}
		currentGroup = append(currentGroup, i)
		prev = i
	}

	println(diffOne*diffThree, pt2)
}

func countGroupCombinations(group []int) int {
	if len(group) == 1 {
		return 1
	}

	target := group[len(group)-1]
	if v, ok := cache[target]; ok {
		return v
	}

	var combinations int

	for i := len(group) - 2; i >= 0 && target-group[i] <= 3; i-- {
		combinations += countGroupCombinations(group[:i+1])
	}
	cache[target] = combinations
	return combinations
}
