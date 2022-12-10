package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	X := 1
	list := make(map[int]int)
	//var arr []int
	var screen [][]string
	for i := 0; i < 6; i++ {
		row := make([]string, 40)
		for j := 0; j < 40; j++ {
			row[j] = "."
		}
		screen = append(screen, row)
	}

	numCycles := 1
	for scanner.Scan() {
		s := scanner.Text()
		strs := strings.Split(s, " ")
		if strs[0] == "noop" {
			list[numCycles] = X
			numCycles++
			continue
		} else {
			val, _ := strconv.Atoi(strs[1])
			list[numCycles] = X
			list[numCycles+1] = X
			X += val
			numCycles += 2
			list[numCycles] = X
		}
	}
	for i := 0; i < len(list); i++ {
		spritePos := list[i+1]
		pixel := (i) % 40
		x := spritePos - pixel
		if x == 1 || x == 0 || x == -1 {
			screen[(i)/40][pixel] = "#"
		}
	}

	for row := 0; row < 6; row++ {
		for col := 0; col < 40; col++ {
			fmt.Print(screen[row][col], "")
		}
		fmt.Print("\n")
	}

	sum := 0
	for i := 20; i <= 220; i += 40 {
		sum += i * list[i]
	}
	fmt.Println(sum)
}
