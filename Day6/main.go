package main

import (
	"bufio"
	"fmt"
	"os"
	//"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	for i := 0; i < len(s); i++ {
		set := make(map[rune]int)
		set[rune(s[i])]++
		set[rune(s[i+1])]++
		set[rune(s[i+2])]++
		set[rune(s[i+3])]++
		if len(set) == 4 {
			fmt.Println(i + 3 + 1)
			break
		}
	}
	// Part 2
	for i := 0; i < len(s); i++ {
		set := make(map[rune]int)
		j := i
		for ; j < i+14; j++ {
			set[rune(s[j])]++
		}
		if len(set) == 14 {
			fmt.Println(j)
			break
		}
	}
}
