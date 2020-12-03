package main

import (
	"adventofcode/utils"
	"fmt"
	"log"
)

type tmp struct {
	right int
	down  int
}

func main() {
	treeMap, err := utils.GetData("3")
	if err != nil {
		log.Fatal("error load file data", err)
	}
	treeCounter := calculateIt(3, 1, treeMap)
	println("task 1:", treeCounter)

	steps := []tmp{
		{right: 1, down: 1},
		{right: 3, down: 1},
		{right: 5, down: 1},
		{right: 7, down: 1},
		{right: 1, down: 2},
	}

	total := 1
	for _, i := range steps {
		t := calculateIt(i.right, i.down, treeMap)
		total = total * t
	}
	println("task 1:", total)
}

func calculateIt(rightStep, downStep int, treeMap []string) int {
	treeCounter := 0
	offset := 0
	for i := 0; i < len(treeMap); i++ {
		if i%downStep != 0 {
			continue
		}

		o := offset % 31
		logStr := fmt.Sprintf("%d,%d - v %s, offset %d", offset, i, string(treeMap[i][o]), o)
		println(logStr)
		if string(treeMap[i][o]) == "#" {
			treeCounter++
		}
		offset += rightStep
	}
	return treeCounter
}
