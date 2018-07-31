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

func getInput() (amount int, exchangeRate float64) {
	amount = getNonNegativeIntegerInput("How many euros are you exchanging? ")
	exchangeRate = getNonNegativeFloatInput("What is the exchange rate? ")
	return
}

func calculateAmountOfDollars(amountEuros int, exchangeRateEuros float64) (amountDollars float64) {
	amountDollars = float64(amountEuros) * exchangeRateEuros / 100.0
	return
}

func printOutput(amountEuros int, exchangeRateEuros float64, amountDollars float64) {
	fmt.Printf("%d euros at an exchange rate of %.2f is %.2f U.S. dollars.\n", amountEuros, exchangeRateEuros, amountDollars)
}

func main() {
	amountEuros, exchangeRateEuros := getInput()
	amountDollars := calculateAmountOfDollars(amountEuros, exchangeRateEuros)
	printOutput(amountEuros, exchangeRateEuros, amountDollars)

}
