package main

import (
	_ "embed"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string
var inputArr []string

func init() {
	input = strings.TrimRight(input, "\n")
	inputArr = strings.Split(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func sortedElvesArr() []int {
	elvesIdx := 0
	elves := make([]int, 1)
	for _, value := range inputArr {
		if value == "" {
			elves = append(elves, 0)
			elvesIdx = elvesIdx + 1
		} else {
			calories, err := strconv.Atoi(value)
			if err != nil {
				panic(err)
			}
			elves[elvesIdx] = elves[elvesIdx] + calories
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(elves)))
	return elves
}

func part1() int {
	elves := sortedElvesArr()
	return elves[0]
}

func part2() int {
	elves := sortedElvesArr()
	calories := 0
	for _, elvesCalories := range elves[:3] {
		calories += elvesCalories
	}
	return calories
}

func main() {
	println(part1())
	println(part2())
}
