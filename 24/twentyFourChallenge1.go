package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func createCharacterCountMap(s string) (characterCountMap map[rune]int) {
	characterCountMap = make(map[rune]int)
	for _, c := range s {
		characterCountMap[c] += 1
	}

	return
}

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	characterCountMap1 := createCharacterCountMap(s1)
	characterCountMap2 := createCharacterCountMap(s2)

	for k, v := range characterCountMap1 {
		if characterCountMap2[k] != v {
			return false
		}
	}

	return true
}

func getNonEmptyInput(msg string) (input string) {
	done := false
	for ; !done; {
		fmt.Print(msg)
		input, _ = bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.TrimSpace(input)
		if len(input) > 0 {
			done = true
		}
	}

	return input
}

func main() {
	fmt.Println("Enter two strings and I'll tell you if they are anagrams:")
	s1 := getNonEmptyInput("Enter the first string: ")
	s2 := getNonEmptyInput("Enter the second string: ")
	if isAnagram(s1, s2) {
		fmt.Printf("\"%s\" and \"%s\" are anagrams\n", s1, s2)
	} else {
		fmt.Printf("\"%s\" and \"%s\" are not anagrams\n", s1, s2)
	}
}
