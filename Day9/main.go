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
	//input := "test.in"
	input := "input.in"
	part1(input)
	part2(input)
}
func part2(input string) {
	f, _ := os.Open(input)
	scanner := bufio.NewScanner(f)
	visited := make(map[[2]int]int)
	knots := [10][2]int{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
	visited[knots[0]] = 1
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
		//fmt.Println(strs)
		for i := 0; i < length; i++ {
			knots[0] = [2]int{knots[0][0] + traverseRL*modifier, knots[0][1] + traverseUD*modifier}
			for j := 1; j < 10; j++ {
				//fmt.Println("Knots: ", knots)
				curr := knots[j]
				parent := knots[j-1]
				deltaX := parent[0] - curr[0]
				deltaY := parent[1] - curr[1]
				absX := (math.Abs(float64(deltaX)))
				absY := (math.Abs(float64(deltaY)))
				flag := true
				if absX == 0 && absY == 2 {
					//fmt.Println("Vert")
					knots[j][1] += modifier
				} else if absY == 0 && absX == 2 {
					//fmt.Println("Horiz")
					knots[j][0] += modifier
				} else if absX > 0 && absY > 0 && !inDistance(curr, parent) {
					//fmt.Println("Diag")
					deltaX /= int(math.Abs(float64(deltaX)))
					deltaY /= int(math.Abs(float64(deltaY)))
					knots[j][0] += deltaX
					knots[j][1] += deltaY
				} else {
					flag = false
				}
				if flag && j == 10 {
					visited[curr] = 1
				}
				if !flag {
					break
				}
			}
		}
	}
	fmt.Println(len(visited))

}
func part1(input string) {
	f, _ := os.Open(input)
	scanner := bufio.NewScanner(f)
	visited := make(map[[2]int]bool)
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
