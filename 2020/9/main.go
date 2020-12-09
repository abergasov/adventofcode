package main

import (
	"adventofcode/utils"
	"log"
	"strconv"
)

const Limit = 25

func main() {
	printData, err := utils.GetData("9")
	if err != nil {
		log.Fatal("error load file data", err)
	}
	digits := convert2Int(printData)
	badNumber := 0
	for i := Limit; i < len(digits); i++ {
		valid := hasSum(digits[i], digits[i-Limit:i])
		if !valid {
			badNumber = digits[i]
			println("FOUND", badNumber)
			break
		}
	}
	for i := range digits {
		found, p := walkForward(i, badNumber, digits)
		if found {

			min, max := fundMinMax(digits[i:p])
			println("FOUND", min, max, min+max)
			break
		}
	}
}

func fundMinMax(digits []int) (min int, max int) {
	min = 999999999
	for _, v := range digits {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return
}

func walkForward(startPosition, needNumber int, digits []int) (bool, int) {
	counter := 0
	for i := startPosition; i < len(digits); i++ {
		counter += digits[i]
		if counter == needNumber {
			return true, i
		}
		if counter > needNumber {
			break
		}
	}
	return false, 0
}

func hasSum(i int, before []int) bool {
	elementMap := make(map[int]struct{})
	for _, s := range before {
		elementMap[s] = struct{}{}
	}
	for e := range elementMap {
		need := i - e
		_, ok := elementMap[need]
		if ok {
			return true
		}
	}
	return false
}

func convert2Int(s []string) []int {
	resp := make([]int, 0, len(s))
	for _, i := range s {
		ii, err := strconv.Atoi(i)
		if err != nil {
			continue
		}
		resp = append(resp, ii)
	}
	return resp
}
