package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {
	input := getInput("What is the rate of return? ")
	inputAsInt, _ := strconv.Atoi(input)
	numberOfYears := 72 / inputAsInt
	fmt.Printf("It will take %d years to double your initial investment.\n", numberOfYears)
}

func getInput(msg string) (input string) {
	done := false
	for ; !done; {
		fmt.Print(msg)
		input, _ = bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.TrimSpace(input)
		if len(input) > 0 {
			if ok, errorMessage := validateInput(input); ok {
				done = true
			} else {
				fmt.Println(errorMessage)

			}
		}
	}

	return input
}

func validateInput(input string) (ok bool, errorMessage string) {
	i, err := strconv.Atoi(input)

	ok = true
	if err != nil {
		ok = false
		errorMessage = "Sorry. That's not a valid input."
	} else if i <= 0 {
		ok = false
		errorMessage = "Please supply a positive number."
	}

	return ok, errorMessage
}