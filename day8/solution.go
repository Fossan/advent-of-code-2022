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

func readTrees() [][]int {
	trees := make([][]int, len(inputArr))
	for i, v := range inputArr {
		trees[i] = make([]int, len(v))
		for j, vv := range v {
			treeHeight, _ := strconv.Atoi(string(vv))
			trees[i][j] = treeHeight
		}
	}
	return trees
}

func checkTop(trees [][]int, i, j int) (bool, int) {
	treeHeight := trees[i][j]
	y := i
	visibleTrees := 0
	visibleUntil := 0
	for y > 0 {
		y--
		visibleUntil++
		if trees[y][j] < treeHeight {
			visibleTrees++
		} else {
			break
		}
	}
	if visibleTrees == i {
		return true, visibleUntil
	}
	return false, visibleUntil
}

func checkRight(trees [][]int, i, j int) (bool, int) {
	treeHeight := trees[i][j]
	x := j
	visibleTrees := 0
	visibleUntil := 0
	for x < len(trees[i])-1 {
		x++
		visibleUntil++
		if trees[i][x] < treeHeight {
			visibleTrees++
		} else {
			break
		}
	}
	if visibleTrees == len(trees[i])-j-1 {
		return true, visibleUntil
	}
	return false, visibleUntil
}

func checkBottom(trees [][]int, i, j int) (bool, int) {
	treeHeight := trees[i][j]
	y := i
	visibleTrees := 0
	visibleUntil := 0
	for y < len(trees)-1 {
		y++
		visibleUntil++
		if trees[y][j] < treeHeight {
			visibleTrees++
		} else {
			break
		}
	}
	if visibleTrees == len(trees[i])-i-1 {
		return true, visibleUntil
	}
	return false, visibleUntil
}

func checkLeft(trees [][]int, i, j int) (bool, int) {
	treeHeight := trees[i][j]
	x := j
	visibleTrees := 0
	visibleUntil := 0
	for x > 0 {
		x--
		visibleUntil++
		if trees[i][x] < treeHeight {
			visibleTrees++
		} else {
			break
		}
	}
	if visibleTrees == j {
		return true, visibleUntil
	}
	return false, visibleUntil
}

func isOnEdge(trees [][]int, i, j int) bool {
	return i == 0 || j == 0 || i == len(trees)-1 || j == len(trees[i])-1
}

func part1() int {
	trees := readTrees()
	visibleTrees := 0

	for i := 0; i < len(trees); i++ {
		for j := 0; j < len(trees[i]); j++ {
			onEdge := isOnEdge(trees, i, j)
			top, _ := checkTop(trees, i, j)
			right, _ := checkRight(trees, i, j)
			bottom, _ := checkBottom(trees, i, j)
			left, _ := checkLeft(trees, i, j)
			if onEdge || top || right || bottom || left {
				visibleTrees++
				continue
			}
		}
	}

	return visibleTrees
}

func part2() int {
	trees := readTrees()
	highestScenicScore := 0

	for i := 0; i < len(trees); i++ {
		for j := 0; j < len(trees[i]); j++ {
			if isOnEdge(trees, i, j) {
				continue
			}

			_, visibleTop := checkTop(trees, i, j)
			_, visibleRight := checkRight(trees, i, j)
			_, visibleBottom := checkBottom(trees, i, j)
			_, visibleLeft := checkLeft(trees, i, j)

			scenicScore := visibleTop * visibleRight * visibleBottom * visibleLeft

			if scenicScore > highestScenicScore {
				highestScenicScore = scenicScore
			}

		}
	}
	return highestScenicScore
}

func main() {
	println(part1())
	println(part2())
}
