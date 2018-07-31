package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func getStringInput(msg string, allowedResponses []string) (retString string) {
	allowedResponsesMap := make(map[string]bool)
	for _, allowedResponse := range allowedResponses {
		allowedResponsesMap[allowedResponse] = true
	}

	return getRestrictedStringInput(msg, allowedResponsesMap)
}

func getRestrictedStringInput(msg string, allowedResponses map[string]bool) (retString string) {
	done := false
	for ; !done; {
		fmt.Print(msg)
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.TrimSpace(input)

		if len(input) > 0 && allowedResponses[input]{
			retString = input
			done = true
		}
	}
	return
}

// Get integer input.
func getNonNegativeIntegerInput(msg string) (retInt int) {
	done := false
	for ; !done; {
		fmt.Print(msg)
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.TrimSpace(input)
		i,err := strconv.Atoi(input)

		if err == nil && i > 0 {
			retInt = i
			done = true
		}
	}
	return
}


func getInput() (amount int, country string) {
	amount = getNonNegativeIntegerInput("Amount: ")
	country = getStringInput("What currency [EU|SWE]: ", []string{"EU", "SE"})
	return
}

func calculateAmountOfDollars(amountInOtherCurrency int, exchangeRateToDollars float64) (amountDollars float64) {
	amountDollars = float64(amountInOtherCurrency) * exchangeRateToDollars / 100.0
	return
}

func printOutput(amountInOtherCurrency int, exchangeRateToDollars float64, amountDollars float64) {
	fmt.Printf("%d at an exchange rate of %.2f is %.2f U.S. dollars.\n", amountInOtherCurrency, exchangeRateToDollars, amountDollars)
}


func main() {
	countryToExchangeRateMap := make(map[string]float64)
	countryToExchangeRateMap["EU"] = 137.51
	countryToExchangeRateMap["SE"] = 15.28

	amount, currency := getInput()
	exchangeRate := countryToExchangeRateMap[currency]
	amountDollars := calculateAmountOfDollars(amount, exchangeRate)
	printOutput(amount, exchangeRate, amountDollars)
}
