package main

import (
	_ "embed"
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

func ranges(s1, s2 string) ([2]int, [2]int) {
	sec1 := strings.Split(s1, "-")
	sec2 := strings.Split(s2, "-")
	x1, _ := strconv.Atoi(sec1[0])
	x2, _ := strconv.Atoi(sec1[1])
	y1, _ := strconv.Atoi(sec2[0])
	y2, _ := strconv.Atoi(sec2[1])
	return [2]int{x1, x2}, [2]int{y1, y2}
}

func compare(s1, s2 string) bool {
	x, y := ranges(s1, s2)
	return x[0] >= y[0] && x[1] <= y[1]
}

func overlap(s1, s2 string) bool {
	x, y := ranges(s1, s2)
	return x[0] <= y[1] && y[0] <= x[1]
}

func part1() int {
	contained := 0
	for _, v := range inputArr {
		sections := strings.Split(v, ",")
		if compare(sections[0], sections[1]) || compare(sections[1], sections[0]) {
			contained = contained + 1
		}
	}
	return contained
}

func part2() int {
	overlapping := 0
	for _, v := range inputArr {
		sections := strings.Split(v, ",")
		if overlap(sections[0], sections[1]) || overlap(sections[1], sections[0]) {
			overlapping = overlapping + 1
		}
	}
	return overlapping
}

func main() {
	println(part1())
	println(part2())
}
