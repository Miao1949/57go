package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"math"
)

//What is the principal amount? 1500
//What is the rate? 4.3
//What is the number of years? 6
//What is the number of times the interest is compounded per year? 4
//$1500 invested at 4.3% for 6 years compounded 4 times per year is $1938.84.

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
			retFloat = i
			done = true
		}
	}
	return
}

func getInput() (principal int, rate float64, numberOfYears int, numberOfCompoundsPerYear int) {
	principal = getNonNegativeIntegerInput("What is the principal amount? ")
	rate = getNonNegativeFloatInput("What is the rate? ")
	numberOfYears = getNonNegativeIntegerInput("What is the number of years? ")
	numberOfCompoundsPerYear = getNonNegativeIntegerInput("What is the number of times the interest is compounded per year? ")
	return
}

func calculateInvestment(principal int, rateAsPercentage float64, numberOfYears int, numberOfCompoundsPerYear int) (valueAfterPeriod float64) {
	rate := rateAsPercentage/100.0
	valueAfterPeriod = float64(principal) * math.Pow(1+rate/float64(numberOfCompoundsPerYear), float64(numberOfYears)*float64(numberOfCompoundsPerYear))
	return

}

func printOutput(principal int, rateAsPercentage float64, numberOfYears int, numberOfCompoundsPerYear int, finalAmount float64) {
	fmt.Printf("$%d invested at %.2f%% for %d years compounded %d times per year is $%.2f.\n", principal, rateAsPercentage, numberOfCompoundsPerYear, numberOfYears, finalAmount)
}

func main() {
	principal, rate, numberOfYears, numberOfCompoundsPerYear := getInput()
	valueAfterPeriod := calculateInvestment(principal, rate, numberOfYears, numberOfCompoundsPerYear)
	printOutput(principal, rate, numberOfYears, numberOfCompoundsPerYear, valueAfterPeriod)
}
