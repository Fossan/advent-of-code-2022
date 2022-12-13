package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"sort"
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

func parsePackets() []any {
	var packets []any
	for _, s := range strings.Split(strings.TrimSpace(input), "\n\n") {
		s := strings.Split(s, "\n")
		var left, right any
		json.Unmarshal([]byte(s[0]), &left)
		json.Unmarshal([]byte(s[1]), &right)
		packets = append(packets, left, right)
	}
	return packets
}

func equal(first, second any) int {
	f, okF := first.(float64)
	s, okS := second.(float64)
	if okF && okS {
		return int(f) - int(s)
	}

	fList := packetsList(first)
	sList := packetsList(second)

	for i := range fList {
		if len(sList) <= i {
			return 1
		}
		r := equal(fList[i], sList[i])
		if r != 0 {
			return r
		}
	}
	if len(sList) == len(fList) {
		return 0
	}
	return -1
}

func packetsList(packet any) []any {
	switch packet.(type) {
	case []any, []float64:
		return packet.([]any)
	case float64:
		return []any{packet}
	}
	panic("Wrong packet type")
}

func part1() int {
	packets := parsePackets()
	orderedPairs := 0
	for i := 0; i < len(packets); i += 2 {
		left, right := packets[i], packets[i+1]
		if equal(left, right) <= 0 {
			orderedPairs += i/2 + 1
		}
	}

	return orderedPairs
}

func part2() int {
	packets := parsePackets()
	packets = append(packets, []any{[]any{2.0}}, []any{[]any{6.0}})
	sort.Slice(packets, func(i, j int) bool { return equal(packets[i], packets[j]) < 0 })

	decoderKey := 1
	for i, packet := range packets {
		if fmt.Sprint(packet) == "[[2]]" || fmt.Sprint(packet) == "[[6]]" {
			decoderKey *= i + 1
		}
	}
	return decoderKey
}

func main() {
	println(part1())
	println(part2())
}
