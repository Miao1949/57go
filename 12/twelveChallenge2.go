package  main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"math"
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

func calculateSimpleInterest(principal int, rateAsPercentage float64, numberOfYears int) (valueAfterPeriod float64){
	rate := 1 + rateAsPercentage / 100.0
	valueAfterPeriod = float64(principal) * math.Pow(rate, float64(numberOfYears))
	return

}

func printOutput(numberOfYears int, rateAsPercentage float64, finalAmount float64) {
	fmt.Printf("After %d years at %.1f%%, the investment will be worth $%.2f.\n", numberOfYears, rateAsPercentage, finalAmount)
}

func main() {
	principal, rateAsPercentage, numberOfYears := getInput()
	valueAfterPeriod := calculateSimpleInterest(principal, rateAsPercentage, numberOfYears)
	printOutput(numberOfYears, rateAsPercentage, valueAfterPeriod)
}
