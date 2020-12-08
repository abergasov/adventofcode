package main

import (
	"adventofcode/utils"
	"log"
	"strconv"
	"strings"
)

func main() {
	commands, err := utils.GetData("8")
	if err != nil {
		log.Fatal("error load file data", err)
	}

	accumulator, _ := calcAccum(commands)
	println(accumulator)

	for i := 0; i < len(commands); i++ {
		accum, fixed := findSolution(i, commands)
		if fixed {
			println("FOUND", accum, i)
			break
		}
	}

}

func findSolution(i int, commands []string) (int, bool) {
	cmnd := strings.Split(commands[i], " ")
	if cmnd[0] == "acc" {
		return 0, false
	}

	if len(cmnd[0]) == 0 {
		return 0, false
	}

	val := getVal(cmnd[1])
	if val == 0 {
		return 0, false
	}

	newCommand := "jmp "
	if val > 0 {
		newCommand += "+" + strconv.Itoa(val)
	} else {
		newCommand += strconv.Itoa(val)
	}

	if cmnd[0] == "jmp" {
		newCommand = "nop +0"
	}
	b := make([]string, len(commands))
	copy(b, commands)
	b[i] = newCommand
	return calcAccum(b)
}

func calcAccum(commands []string) (int, bool) {
	iterators := make([]int, 0, 1000)
	accumulator := 0
	for i := 0; i < len(commands); i++ {
		if len(commands[i]) == 0 {
			break
		}
		cmnd := strings.Split(commands[i], " ")
		val := getVal(cmnd[1])
		if utils.IntInSlice(i, iterators) {
			return accumulator, false
		}
		iterators = append(iterators, i)
		if cmnd[0] == "acc" {
			accumulator += val
		} else if cmnd[0] == "jmp" {
			i += val - 1
		}
	}
	return accumulator, true
}

func getVal(val string) int {
	i, _ := strconv.Atoi(val[1:])
	if val[0:1] == "+" {
		return i
	}
	return -1 * i
}
