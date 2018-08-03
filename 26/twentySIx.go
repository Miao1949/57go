package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"math"
)

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

func getInput() (balance uint64, apr uint64, monthlyPayment uint64) {
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

func printResult(numberOfMonths int) {
	fmt.Printf("It will take you %d months to pay off this card.\n", numberOfMonths)
}

func main() {
	printResult(calculateMonthsUntilPaidOff(getInput()))
}

