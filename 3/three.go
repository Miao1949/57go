package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

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
	quote := getNonEmptyInput("What is the quote? ")
	who := getNonEmptyInput("Who said it? ")
	output := who + " says " + "\"" + quote + "\""
	fmt.Println(output)
}
