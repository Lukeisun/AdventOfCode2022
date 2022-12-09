package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var graph [][]int
	var row int
	c := 0
	for scanner.Scan() {
		s := scanner.Text()
		c++
		if graph == nil {
			graph = make([][]int, len(s))
			row = 0
			for i := range graph {
				graph[i] = make([]int, len(s))
			}
		}
		for i := range s {
			graph[row][i] = int(s[i] - '0')
		}
		row++
	}
	// amount of visible at start
	counter := len(graph[0])*4 - 4
	dirArray := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	max := math.MinInt
	for i := 1; i < len(graph)-1; i++ {
		for j := 1; j < len(graph[i])-1; j++ {
			currVal := graph[i][j]
			counter += part1(graph, i, j, currVal, dirArray)
			scenic := part2(graph, i, j, currVal, dirArray)
			if scenic > max {
				max = scenic
			}
		}
	}
	fmt.Println(counter, max)
}
func part2(graph [][]int, row int, col int, currVal int, dirArray [][]int) int {
	count := 1
	for k := range dirArray {
		currScore := score(graph, row, col, currVal, dirArray[k], 1)
		count *= currScore
	}
	return count
}
func part1(graph [][]int, row int, col int, currVal int, dirArray [][]int) int {
	for k := range dirArray {
		count := check(graph, row, col, currVal, dirArray[k])
		if count == 1 {
			return 1
		}
	}
	return 0
}

func score(graph [][]int, row int, col int, currVal int, dir []int, currScore int) int {
	length := len(graph)
	newRow := row + dir[0]
	newCol := col + dir[1]
	if newRow >= length || newRow < 0 || newCol >= length || newCol < 0 {
		return 0
	}
	// found a bigger number
	if graph[newRow][newCol] >= currVal {
		return currScore
	}
	return currScore + score(graph, newRow, newCol, currVal, dir, 1)
}
func check(graph [][]int, row int, col int, currVal int, dir []int) int {
	length := len(graph)
	newRow := row + dir[0]
	newCol := col + dir[1]
	// reached the end
	if newRow >= length || newRow < 0 || newCol >= length || newCol < 0 {
		return 1
	}
	// found a bigger number
	if graph[newRow][newCol] >= currVal {
		return 0
	}
	return check(graph, newRow, newCol, currVal, dir)
}
