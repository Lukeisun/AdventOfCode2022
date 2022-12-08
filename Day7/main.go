package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Dir struct {
	name     string
	size     int
	parent   *Dir
	children []*Dir
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	root := Dir{"/", 0, nil, []*Dir{}}
	currDir := &root
	for scanner.Scan() {
		s := scanner.Text()
		if strings.Contains(s, "cd") {
			currDir = cd(currDir, s)
		} else if strings.Contains(s, "ls") {
			currDir = ls(scanner, currDir)
		}
	}
	sum := 0
	// Part 1
	findAllLessThan100k(&root, &sum)
	fmt.Println(sum)
	// Part 2
	available := 70000000 - root.size
	needed := 30000000
	max := math.MaxInt
	smallest := 0
	smallestToDelete(&root, available, &max, needed, &smallest)
	fmt.Println(smallest)
}
func findAllLessThan100k(currDir *Dir, sum *int) {
	if currDir.size < 100000 {
		*sum += currDir.size
	}
	for i := 0; i < len(currDir.children); i++ {
		findAllLessThan100k(currDir.children[i], sum)
	}
}
func smallestToDelete(currDir *Dir, available int, max *int, needed int, smallest *int) {
	if currDir.size > needed-available && currDir.size+available < *max {
		*max = currDir.size + available
		*smallest = currDir.size
	}
	for i := 0; i < len(currDir.children); i++ {
		smallestToDelete(currDir.children[i], available, max, needed, smallest)
	}
}

func cd(currDir *Dir, s string) *Dir {
	strs := strings.Split(s, " ")
	if strs[2] == ".." {
		return currDir.parent
	} else {
		for i := 0; i < len(currDir.children); i++ {
			if currDir.children[i].name == strs[2] {
				return currDir.children[i]
			}
		}
	}
	return nil
}

func ls(scanner *bufio.Scanner, currDir *Dir) *Dir {
	for scanner.Scan() {
		s := scanner.Text()
		strs := strings.Split(s, " ")
		if strs[1] == "cd" {
			currDir = cd(currDir, s)
			return currDir
		}
		if strs[0] == "dir" {
			name := strs[1]
			temp := Dir{name, 0, currDir, []*Dir{}}
			currDir.children = append(currDir.children, &temp)
			continue
		}
		amount, _ := strconv.Atoi(strs[0])
		currDir.size += amount
		curr := currDir.parent
		for curr != nil {
			curr.size += amount
			curr = curr.parent
		}
	}
	return nil
}
