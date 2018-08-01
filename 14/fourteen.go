package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func getNonNegativeIntegerInput(msg string) (retInt int) {
	done := false
	for ; !done; {
		fmt.Print(msg)
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.TrimSpace(input)
		i, err := strconv.Atoi(input)

		if err == nil && i > 0 {
			retInt = i
			done = true
		}
	}
	return
}

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

func getInput() (orderAmount int, state string) {
	orderAmount = getNonNegativeIntegerInput("What is the order amount? ")
	state = getNonEmptyInput("What is the state? ")

	return
}

func main() {
	orderAmount, state := getInput()
	total := float64(orderAmount)
	if state == "WI" {
		tax := total * 0.055
		fmt.Printf("The subtotal is $%.2f\n", total)
		fmt.Printf("The tax is $%.2f\n", tax)
		total += tax
	}

	fmt.Printf("The total is %.2f.\n", total)
}
