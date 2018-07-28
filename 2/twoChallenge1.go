package main

import (
	"fmt"
	"strings"
)

func getNonEmptyInput(msg string) (input string) {
	done := false
	for ; !done; {
		fmt.Print(msg)
		fmt.Scanf("%s", &input)
		input = strings.TrimSpace(input)
		if len(input) > 0 {
			done = true
		}
	}

	return input
}

func main() {
	input := getNonEmptyInput("What is the input string? ")
	fmt.Println(input, "has", len(input), "characters.")
}
