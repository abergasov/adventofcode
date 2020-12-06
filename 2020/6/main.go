package main

import (
	"adventofcode/utils"
	"log"
	"strings"
)

func main() {
	questions, err := utils.GetData("6")
	if err != nil {
		log.Fatal("error load file data", err)
	}

	println(stepOne(questions))
	println(stepTwo(questions))
}

func stepOne(questions []string) int {
	total := 0
	group := make([]string, 0, 10)
	for _, r := range questions {
		if len(r) == 0 {
			tmp := len(utils.UniqueSlice(group))
			total += tmp
			group = make([]string, 0, 10)
			continue
		}
		group = append(group, strings.Split(r, "")...)
	}
	if len(group) > 0 {
		tmp := len(utils.UniqueSlice(group))
		total += tmp
	}
	return total
}

func stepTwo(questions []string) int {
	total := 0
	group := make([]string, 0, 10)
	for _, r := range questions {
		if len(r) == 0 {
			tmp := calculateCount(group)
			total += tmp
			group = make([]string, 0, 10)
			continue
		}
		group = append(group, r)
	}
	if len(group) > 0 {
		tmp := calculateCount(group)
		total += tmp
	}
	return total
}

func calculateCount(answers []string) int {
	sp := strings.Split(strings.Join(answers, ""), "")
	keys := make(map[string]int)
	for _, entry := range sp {
		if _, value := keys[entry]; !value {
			keys[entry] = 0
		}
		keys[entry]++
	}

	total := 0
	need := len(answers)
	if need == 1 {
		return len(sp)
	}
	for i := range keys {
		if keys[i] == need {
			total++
		}
	}
	return total
}
