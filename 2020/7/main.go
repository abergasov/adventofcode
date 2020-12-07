package main

import (
	"adventofcode/utils"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var re = regexp.MustCompile("[0-9]+ shiny gold bag")

func main() {
	bags, err := utils.GetData("7")
	if err != nil {
		log.Fatal("error load file data", err)
	}

	parentBags := make([]string, 0, 100)
	childBags := make(map[string]int)
	for _, b := range bags {
		dt := strings.Split(b, "contain")
		if len(dt) != 2 {
			continue
		}
		if dt[0] == "shiny gold bags " {
			ch := strings.Split(dt[1], ", ")
			for _, chB := range ch {
				chBData := cleanStr(chB)
				cnt := chBData[0:2]
				cnt = strings.TrimSpace(cnt)
				cntI, _ := strconv.Atoi(cnt)
				bg := chBData[2:]
				childBags[cleanStr(bg)] = cntI
			}
		}
		if !re.MatchString(b) {
			continue
		}
		parentBags = append(parentBags, cleanStr(dt[0]))
	}
	parentBags = utils.UniqueSlice(parentBags)
	for _, pb := range parentBags {
		parentBags = append(parentBags, createNestedData(pb, bags)...)
	}
	parentBags = utils.UniqueSlice(parentBags)
	println(len(parentBags))
	counter := 0
	for i := range childBags {
		counter += childBags[i]
		counter += childBags[i] * calculateChilds(i, bags)
	}
	println(counter)
}

func createNestedData(bag string, bags []string) []string {
	res := make([]string, 0, 100)
	for _, b := range bags {
		dt := strings.Split(b, "contain")
		if len(dt) != 2 {
			continue
		}

		if strings.Contains(dt[1], bag) {
			res = append(res, createNestedData(cleanStr(dt[0]), bags)...)
		}
	}
	res = append(res, bag)
	return res
}

func calculateChilds(bag string, bags []string) int {
	counter := 0
	for _, b := range bags {
		dt := strings.Split(b, "contain")
		if len(dt) != 2 {
			continue
		}

		if !strings.Contains(dt[0], bag) {
			continue
		}

		ch := strings.Split(dt[1], ", ")
		for _, chB := range ch {
			if chB == " no other bags." {
				return 0
			}
			chBData := cleanStr(chB)
			cnt := chBData[0:2]
			cnt = strings.TrimSpace(cnt)
			cntI, _ := strconv.Atoi(cnt)
			bg := chBData[2:]
			counter += cntI
			counter += cntI * calculateChilds(bg, bags)
		}

	}

	return counter
}

func cleanStr(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, ".", "")
	return strings.ReplaceAll(s, "bags", "bag")
}
