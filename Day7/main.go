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
	children []Dir
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	root := Dir{"/", 0, nil, []Dir{}}
	currDir := &root
	directoryList := make(map[string]*Dir)
	directoryList["/"] = &root
	fmt.Printf("%p\n", &root)
	fmt.Printf("%p\n", currDir)
	for scanner.Scan() {
		s := scanner.Text()
		if strings.Contains(s, "cd") {
			currDir = cd(currDir, s, directoryList)
		} else if strings.Contains(s, "ls") {
			fmt.Println(currDir)
			fmt.Printf("%p\n", &currDir)
			currDir = ls(scanner, currDir, directoryList)
		}
	}
	fmt.Println(root)
	for i := 0; i < len(root.children); i++ {
		fmt.Println(currDir.children[i])
	}
}

// function that change reference of passed in struct
func cd(currDir *Dir, s string, directoryList map[string]*Dir) *Dir {
	strs := strings.Split(s, " ")
	if strs[2] == ".." {
		return currDir.parent
	} else if val, ok := (directoryList)[strs[2]]; ok {
		fmt.Println("hi")
		fmt.Println(strs[2])
		fmt.Println(currDir)
		fmt.Printf("%p\n", currDir)
		fmt.Printf("%p\n", val)
		return val
	} else {
		temp := Dir{strs[2], 0, currDir, []Dir{}}
		currDir.children = append(currDir.children, temp)
		return &temp
	}
}

func ls(scanner *bufio.Scanner, currDir *Dir, directoryList map[string]*Dir) *Dir {
	for scanner.Scan() {
		s := scanner.Text()
		strs := strings.Split(s, " ")
		// fmt.Println("ls", currDir)
		// fmt.Printf("%p\n", currDir)
		if strs[1] == "cd" {
			currDir = cd(currDir, s, directoryList)
			return currDir
		}
		if strs[0] == "dir" {
			if directoryList[strs[1]] == nil {
				fmt.Println(directoryList)
				directoryList[strs[1]] = &Dir{strs[1], 0, *&currDir, []Dir{}}
			}
			continue
		}
		amount, _ := strconv.Atoi(strs[0])
		currDir.size += amount
		curr := (currDir).parent
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
