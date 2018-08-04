package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {
	input := getInput("What is the rate of return? ", "Sorry. That's not a valid input.")
	inputAsInt, _ := strconv.Atoi(input)
	numberOfYears := 72 / inputAsInt
	fmt.Printf("It will take %d years to double your initial investment.\n", numberOfYears)
}

func getInput(msg string, errorMessage string) (input string) {
	done := false
	for ; !done; {
		fmt.Print(msg)
		input, _ = bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.TrimSpace(input)
		if len(input) > 0 {
			if validateInput(input) {
				done = true
			} else {
				fmt.Println(errorMessage)

			}
		}
	}

	return input
}

func validateInput(input string) bool {
	i, err := strconv.Atoi(input)

	return err == nil && i > 0
}