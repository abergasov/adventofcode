package main

import (
	"adventofcode/utils"
	"fmt"
	"log"
	"sort"
)

func main() {
	passTickets, err := utils.GetData("5")
	if err != nil {
		log.Fatal("error load file data", err)
	}
	maxId := 0
	ids := make([]int, 0, 900)
	for _, i := range passTickets {
		if len(i) == 0 {
			continue
		}

		id := processKey(i)
		if maxId < id {
			maxId = id
		}
		ids = append(ids, id)
	}
	println(maxId)
	sort.Ints(ids)
	prev := ids[0] - 1
	for _, i := range ids {
		if i != prev+1 {
			println(i - 1)
			//break
		}
		prev = i
	}
	println(maxId)
}

func processKey(i string) int {
	row := parseRow(i[0:7])
	col := parseRow(i[7:10])
	id := row*8 + col
	println(fmt.Sprintf("%s, %d * 8 + %d = %d", i, row, col, id))
	return id
}

func parseRow(key string) int {
	upLimit := 128
	lowerHalf := "F"
	upperHalf := "B"
	if len(key) == 3 {
		upLimit = 8
		upperHalf = "R"
		lowerHalf = "L"
	}
	downLimit := 0
	iteration := 0
	for {
		letter := string(key[iteration])
		if letter == lowerHalf {
			upLimit -= (upLimit - downLimit) / 2
		} else if letter == upperHalf {
			downLimit += (upLimit - downLimit) / 2
		}

		iteration++
		if len(key)-1 == iteration {
			letter = string(key[iteration])
			if letter == lowerHalf {
				return downLimit
			}
			return upLimit - 1
		}
	}
}
