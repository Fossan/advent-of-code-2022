package main

import (
	_ "embed"
	"fmt"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string
var inputArr []string

type Stack []string

func (s *Stack) push(str ...string) {
	*s = append(*s, str...)
}

func (s *Stack) pop() string {
	i := s.len() - 1
	v := (*s)[i]
	*s = (*s)[:i]
	return v
}

func (s *Stack) popN(n int) []string {
	i := s.len() - n
	v := (*s)[i:]
	*s = (*s)[:i]
	return v
}

func (s *Stack) peek() string {
	i := s.len() - 1
	return (*s)[i]
}

func (s *Stack) len() int {
	return len(*s)
}

func init() {
	input = strings.TrimRight(input, "\n")
	inputArr = strings.Split(input, "\n\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func readStackAndCommands() ([]Stack, []string) {
	stackLines := strings.Split(inputArr[0], "\n")
	commands := strings.Split(inputArr[1], "\n")
	allStacks := make([]Stack, len(stackLines[len(stackLines)-1]))
	for i := len(stackLines) - 1; i >= 0; i-- {
		for i, c := range stackLines[i] {
			if unicode.IsLetter(c) {
				allStacks[i].push(string(c))
			}
		}
	}

	var stacksNe []Stack
	for _, s := range allStacks {
		if s.len() != 0 {
			stacksNe = append(stacksNe, s)
		}
	}
	return stacksNe, commands
}

func readTop(stacks []Stack) string {
	res := ""
	for _, s := range stacks {
		res = res + s.peek()
	}
	return res
}

func part1() string {
	stacks, commands := readStackAndCommands()

	for _, inp := range commands {
		var amount, from, to int
		fmt.Sscanf(inp, "move %d from %d to %d", &amount, &from, &to)
		i := 1
		for i <= amount {
			stacks[to-1].push(stacks[from-1].pop())
			i = i + 1
		}
	}

	return readTop(stacks)
}

func part2() string {
	stacks, commands := readStackAndCommands()

	for _, inp := range commands {
		var amount, from, to int
		fmt.Sscanf(inp, "move %d from %d to %d", &amount, &from, &to)
		stacks[to-1].push(stacks[from-1].popN(amount)...)
	}

	return readTop(stacks)
}

func main() {
	println(part1())
	println(part2())
}
