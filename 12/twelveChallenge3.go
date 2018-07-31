package  main

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

func getNonNegativeFloatInput(msg string) (retFloat float64) {
	done := false
	for ; !done; {
		fmt.Print(msg)
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.TrimSpace(input)
		i, err := strconv.ParseFloat(input, 64)

		if err == nil && i > 0 {
			retFloat= i
			done = true
		}
	}
	return
}

func getInput() (principal int, rate float64, numberOfYears int) {
	principal = getNonNegativeIntegerInput("Enter the principal: ")
	rate = getNonNegativeFloatInput("Enter the rate of interest: ")
	numberOfYears = getNonNegativeIntegerInput("Enter the number of years: ")
	return
}


func printOutput(numberOfYears int, rateAsPercentage float64, finalAmount float64) {
	fmt.Printf("After %d years at %.1f%%, the investment will be worth $%.2f.\n", numberOfYears, rateAsPercentage, finalAmount)
}

func printYearlyAmountAndReturnFinalAmount(principal int, rateAsPercentage float64, numberOfYears int) (valueAfterPeriod float64) {
	currentAmount := float64(principal)
	rate := 1 + rateAsPercentage / 100.0
	for year := 1; year <= numberOfYears; year++ {
		currentAmount *= rate
		fmt.Printf("The amouont after year %d is %.2f\n", year, currentAmount)
	}

	valueAfterPeriod = currentAmount
	return
}

func main() {
	principal, rateAsPercentage, numberOfYears := getInput()
	valueAfterPeriod := printYearlyAmountAndReturnFinalAmount(principal, rateAsPercentage, numberOfYears)
	printOutput(numberOfYears, rateAsPercentage, valueAfterPeriod)
}
