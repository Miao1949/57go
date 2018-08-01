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
	subtotalString := ""
	taxString := ""
	const WI = "WI"
	const WISCONSIN = "WISCONSIN"
	if strings.ToUpper(state) == WI || strings.ToUpper(state) == WISCONSIN {
		tax := total * 0.055
		subtotalString = fmt.Sprintf("The subtotal is $%.2f\n", total)
		taxString = fmt.Sprintf("The tax is $%.2f\n", tax)
		total += tax
	}
	totalString := fmt.Sprintf("The total is %.2f.\n", total)

	fmt.Print(subtotalString + taxString + totalString)
}
