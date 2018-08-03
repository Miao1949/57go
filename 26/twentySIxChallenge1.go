package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"math"
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

func getNonNegativeIntegerInput(msg string) (retInt uint64) {
	done := false
	for ; !done; {
		fmt.Print(msg)
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.TrimSpace(input)
		i, err := strconv.ParseUint(input, 10, 64)

		if err == nil {
			retInt = i
			done = true
		}
	}
	return
}

func getInputForNumberOfMonths() (balance uint64, apr uint64, monthlyPayment uint64) {
	balance = getNonNegativeIntegerInput("What is your balance? ")
	apr = getNonNegativeIntegerInput("What is the APR on the card (as a percent)? ")
	monthlyPayment = getNonNegativeIntegerInput("What is the monthly payment you can make? ")
	return
}

func calculateMonthsUntilPaidOff(balance uint64, apr uint64, monthlyPayment uint64) (numberOfMonths int) {
	dailyRate := float64(apr) / 100 / 365
	nominator := math.Log10(1 + float64(balance) / float64(monthlyPayment) * (1 - math.Pow(1 + dailyRate, 30)))
	denominator := math.Log10(1.0 + dailyRate)
	numberOfMonths = int(math.Ceil(-1.0 / 30.0 * nominator / denominator))
	return
}

func printNumberOfMonthsResult(numberOfMonths int) {
	fmt.Printf("It will take you %d months to pay off this card.\n", numberOfMonths)
}

func getInputForMontlyPayment() (balance uint64, apr uint64, numberOfMonths uint64) {
	balance = getNonNegativeIntegerInput("What is your balance? ")
	apr = getNonNegativeIntegerInput("What is the APR on the card (as a percent)? ")
	numberOfMonths = getNonNegativeIntegerInput("How many months would you like to pay off the debt in? ")
	return
}


func calculateMontlyPayment(balance uint64, apr uint64, numberOfMonths uint64) (monthlyPayment int) {
	dailyRate := float64(apr) / 100 / 365

	denominator := math.Pow(-30 * float64(numberOfMonths) * math.Log10(1 + dailyRate), 10) - 1
	nominator := float64(balance) * (1 - math.Pow(1 + dailyRate, 30))
	monthlyPayment = int(math.Ceil(nominator / denominator))
	return
}

func printMonthlyPaymentResult(monthlyPayment int) {
	fmt.Printf("The monthly payment is %d.\n", monthlyPayment)
}


func main() {
	choise := getStringInput("Calculate number of months(m) or monthly payment(p)?", []string{"m", "p"})
	if choise == "m" {
		printNumberOfMonthsResult(calculateMonthsUntilPaidOff(getInputForNumberOfMonths()))
	} else {
		printMonthlyPaymentResult(calculateMontlyPayment(getInputForMontlyPayment()))
	}

}

