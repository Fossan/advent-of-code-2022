package main

import (
	_ "embed"
	"math"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string
var inputArr []string

type monkey struct {
	items     []int
	operation func(int) int
	divisible int
	passTrue  int
	passFalse int
	inspected int
}

func init() {
	input = strings.TrimRight(input, "\n")
	inputArr = strings.Split(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func catchMonkeys() ([]monkey, int) {
	monkeys := make([]monkey, 0)
	lcm := 1
	for i := 0; i < len(inputArr); i++ {
		numsStr := strings.Split(inputArr[i+1], ": ")[1]
		nums := make([]int, 0)
		for _, v := range strings.Split(numsStr, ", ") {
			vI, _ := strconv.Atoi(v)
			nums = append(nums, vI)
		}
		opStr := strings.Split(inputArr[i+2], "Operation: new = old ")[1]
		opParts := strings.Split(opStr, " ")
		var operation func(int) int
		if opParts[0] == "+" {
			if opParts[1] != "old" {
				num, _ := strconv.Atoi(opParts[1])
				operation = func(i int) int { return i + num }
			} else {
				operation = func(i int) int { return i + i }
			}
		} else {
			if opParts[1] != "old" {
				num, _ := strconv.Atoi(opParts[1])
				operation = func(i int) int { return i * num }
			} else {
				operation = func(i int) int { return i * i }
			}
		}
		divisible, _ := strconv.Atoi(strings.Split(inputArr[i+3], "Test: divisible by ")[1])
		ifTrue, _ := strconv.Atoi(strings.Split(inputArr[i+4], "If true: throw to monkey ")[1])
		ifFalse, _ := strconv.Atoi(strings.Split(inputArr[i+5], "If false: throw to monkey ")[1])
		monkey := monkey{items: nums, operation: operation, divisible: divisible, passTrue: ifTrue, passFalse: ifFalse, inspected: 0}
		monkeys = append(monkeys, monkey)

		lcm *= divisible
		i += 6
	}
	return monkeys, lcm
}

func performTests(monkeys []monkey, iterations int, testFunc func(monkey monkey, item int) int) {
	for i := 0; i < iterations; i++ {
		for idx, monkey := range monkeys {
			for _, item := range monkey.items {
				result := testFunc(monkey, item)
				if result%monkey.divisible == 0 {
					monkeys[monkey.passTrue].items = append(monkeys[monkey.passTrue].items, result)
				} else {
					monkeys[monkey.passFalse].items = append(monkeys[monkey.passFalse].items, result)
				}
				monkeys[idx].inspected += 1
			}
			monkeys[idx].items = make([]int, 0)
		}
	}
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspected > monkeys[j].inspected
	})
}

func part1() int {
	monkeys, _ := catchMonkeys()
	testFunc := func(monkey monkey, item int) int { return int(math.Floor(float64(monkey.operation(item) / 3))) }
	performTests(monkeys, 20, testFunc)
	return monkeys[0].inspected * monkeys[1].inspected
}

func part2() int {
	monkeys, lcm := catchMonkeys()
	testFunc := func(monkey monkey, item int) int { return monkey.operation(item) % lcm }
	performTests(monkeys, 10_000, testFunc)
	return monkeys[0].inspected * monkeys[1].inspected
}

func main() {
	println(part1())
	println(part2())
}
