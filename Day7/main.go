package main

import (
	"bufio"
	"fmt"
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
	directoryList := make(map[string]*Dir)
	directoryList["/"] = &root
	for scanner.Scan() {
		s := scanner.Text()
		if strings.Contains(s, "cd") {
			currDir = cd(currDir, s, directoryList)
		} else if strings.Contains(s, "ls") {
			currDir = ls(scanner, currDir, directoryList)
		}
	}
	sum := 0 
	findAllLessThan100k(&root, &sum)
	fmt.Println(sum)
}
func findAllLessThan100k(currDir *Dir, sum *int){
	if currDir.size < 100000 {
		*sum += currDir.size
	}
	for i := 0; i < len(currDir.children); i++ {
		findAllLessThan100k(currDir.children[i], sum)
	}
}
// function that change reference of passed in struct
func cd(currDir *Dir, s string, directoryList map[string]*Dir) *Dir {
	strs := strings.Split(s, " ")
	if strs[2] == ".." {
		return currDir.parent
	} else if val, ok := (directoryList)[strs[2]]; ok {
		// fmt.Println("hi")
		// fmt.Println(strs[2])
		// fmt.Println(currDir)
		// fmt.Printf("%p\n", currDir)
		// fmt.Printf("%p\n", val)
		return val
	}
	fmt.Println("err")
	return currDir
}

func ls(scanner *bufio.Scanner, currDir *Dir, directoryList map[string]*Dir) *Dir {
	for scanner.Scan() {
		s := scanner.Text()
		strs := strings.Split(s, " ")
		// fmt.Printf("%p\n", currDir)
		if strs[1] == "cd" {
			currDir = cd(currDir, s, directoryList)
			return currDir
		}
		if strs[0] == "dir" {
			if directoryList[strs[1]] == nil {
				temp := Dir{strs[1], 0, currDir, []*Dir{}}
				currDir.children = append(currDir.children, &temp)
				directoryList[strs[1]] = &temp
			}
			continue
		}
		amount, _ := strconv.Atoi(strs[0])
		currDir.size += amount
		curr := currDir.parent
		// fmt.Println("hi", currDir)
		// fmt.Printf("%p\n", currDir)
		// fmt.Println(currDir.parent)
		// fmt.Println(curr)
		for curr != nil {
			curr.size += amount
			curr = curr.parent
		}
	}
	return currDir
}
