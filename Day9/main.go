package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
}
func part1() {
	visited := make(map[[2]int]bool)
	scanner := bufio.NewScanner(os.Stdin)
	head := [2]int{0, 0}
	tail := [2]int{0, 0}
	visited[head] = true
	for scanner.Scan() {
		s := scanner.Text()
		strs := strings.Split(s, " ")
		to := [2]int{0, 0}
		switch strs[0] {
		case "R":
			to[0], _ = strconv.Atoi(strs[1])
		case "L":
			to[0], _ = strconv.Atoi(strs[1])
			to[0] *= -1
		case "U":
			to[1], _ = strconv.Atoi(strs[1])
		case "D":
			to[1], _ = strconv.Atoi(strs[1])
			to[1] *= -1
		}
		modifier := 1
		if to[0] < 0 || to[1] < 0 {
			modifier = -1
		}
		var temp [2]int
		var length int
		traverseRL := 0
		traverseUD := 0
		if to[0] != 0 {
			length = int(math.Abs(float64(to[0])))
			traverseRL = 1
		} else if to[1] != 0 {
			length = int(math.Abs(float64(to[1])))
			traverseUD = 1
		}
		for i := 0; i < length; i++ {
			temp = [2]int{head[0] + traverseRL*modifier, head[1] + traverseUD*modifier}
			if !inDistance(tail, temp) {
				tail = head
				visited[tail] = true
			}
			head = temp
		}
	}
	fmt.Println(len(visited))
}
func inDistance(a [2]int, b [2]int) bool {
	dist := math.Sqrt(math.Pow(float64(a[0]-b[0]), 2) + math.Pow(float64(a[1]-b[1]), 2))
	if dist > math.Sqrt2 {
		return false
	}
	return true
}
