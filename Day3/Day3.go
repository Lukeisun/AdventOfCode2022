package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
	//"strings"
)

func part1() {
	fmt.Println("hello")
	scanner := bufio.NewScanner(os.Stdin)
	count := 0
	for scanner.Scan() {
		map1 := make(map[rune]int)
		map2 := make(map[rune]int)
		//map3 := make(map[string]int)
		s := scanner.Text()
		length := len(s) / 2
		firstString := s[:length]
		secondString := s[length:]
		for i := 0; i < length; i++ {
			map1[rune(firstString[i])]++
			map2[rune(secondString[i])]++
		}
		flagToBreak := false
		for k1 := range map1 {
			if flagToBreak {
				break
			}
			for k2 := range map2 {
				if k1 == k2 {
					flagToBreak = true
					// val, _ := strconv.Atoi(k1)
					// fmt.Println(val)
					if unicode.IsLower(k1) {
						count += int(k1) - int('a') + 1
					} else {
						count += int(k1) - int('A') + 27
					}
					break
				}
			}
		}
		fmt.Println(count)
	}
}
func part2() {
	scanner := bufio.NewScanner(os.Stdin)
	count := 0
	for scanner.Scan() {
		s1 := scanner.Text()
		scanner.Scan()
		s2 := scanner.Text()
		scanner.Scan()
		s3 := scanner.Text()
		flagToBreak := false
		for i := 0; i < len(s1); i++ {
			if flagToBreak {
				break
			}
			for j := 0; j < len(s2); j++ {
				if flagToBreak {
					break
				}
				if s1[i] == s2[j] {
					for k := 0; k < len(s3); k++ {
						if s1[i] == s3[k] {
							if unicode.IsLower(rune(s1[i])) {
								count += int(s1[i]) - int('a') + 1
							} else {
								count += int(s1[i]) - int('A') + 27
							}
							flagToBreak = true
							break
						}
					}
				}
			}
		}
		fmt.Println(count)
	}
}
func main() {
	part2()
}
