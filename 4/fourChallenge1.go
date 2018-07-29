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
	noun :=getNonEmptyInput("Enter a noun: ")
	noun2 :=getNonEmptyInput("Enter another noun: ")
	verb :=getNonEmptyInput("Enter a verb: ")
	adjective:=getNonEmptyInput("Enter a adjective: ")
	adjective2:=getNonEmptyInput("Enter another adjective: ")
	adverb:=getNonEmptyInput("Enter a adverb: ")
	fmt.Printf("Do you %s your %s %s and your %s %s %s? That's hillarious!\n", verb, adjective, noun, adjective2, noun2, adverb)
}



