package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

// Get non-empty input.
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
	noun := getNonEmptyInput("Enter a noun: ")
	verb := getNonEmptyInput("Enter a verb: ")
	adjective := getNonEmptyInput("Enter a adjective: ")
	adverb := getNonEmptyInput("Enter a adverb: ")

	if noun == "cat" {
		fmt.Printf("Do you %s your %s %s %s? That's great!\n", verb, adjective, noun, adverb)
	} else {
		fmt.Printf("Do you %s your %s %s %s? That's hillarious!\n", verb, adjective, noun, adverb)
	}
}



