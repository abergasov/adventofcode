package main

import (
	"adventofcode/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	passwords, err := utils.GetData("2")
	if err != nil {
		log.Fatal("error load file data", err)
	}
	validPassCount := 0
	validPassCountSecond := 0
	for _, v := range passwords {
		p := strings.Split(v, " ")
		if len(p) != 3 {
			continue
		}
		letter := p[1][0:1]
		count := strings.Count(p[2], letter)
		mM := strings.Split(p[0], "-")
		min, _ := strconv.Atoi(mM[0])
		max, _ := strconv.Atoi(mM[1])
		if min <= count && count <= max {
			validPassCount++
		}

		first := string(p[2][min-1])
		second := string(p[2][max-1])
		if (first == letter || second == letter) && second != first {
			validPassCountSecond++
		}
	}
	println(fmt.Sprintf("fst pt: %d", validPassCount))
	println(fmt.Sprintf("snd pt: %d", validPassCountSecond))
}
