package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

const EauClaireCountyExtraTax = 0.005
const DunnCountyExtraTax = 0.004
const WisconsinTax = 0.05
const IllinoisTax = 0.08
const SwededTax = 0.25

const Illinois = "ILLINOIS"
const IllinoisAbbreviation = "IL"
const Wisconsin = "WISCONSIN"
const WisconsinAbbreviation = "WI"
const EauClaire = "EAUCLAIRE"
const Dunn = "DUNN"
const Sweden = "SWEDEN"
const SwedenAbbreviation = "SWE"

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


func getInput() (orderAmount uint64, state string, county string) {
	orderAmount = getNonNegativeIntegerInput("What is the order amount? ")
	state = strings.ToUpper(getNonEmptyInput("What state do you live in? "))
	if state == Wisconsin || state == WisconsinAbbreviation {
		county = strings.ToUpper(getNonEmptyInput("What county do you live in? "))
	}  else {
		county = ""
	}

	return
}

func calculateTax(orderAmount uint64, state string, county string) (total float64, tax float64) {
	taxRate := 0.0
	if state == Wisconsin || state == WisconsinAbbreviation {
		taxRate = WisconsinTax

		if county == EauClaire {
			taxRate += EauClaireCountyExtraTax
		} else if county == Dunn  {
			taxRate += DunnCountyExtraTax
		}
	} else if state == Illinois || state == IllinoisAbbreviation {
		taxRate = IllinoisTax
	} else if state == Sweden || state == SwedenAbbreviation {
		taxRate = SwededTax
	}

	tax = float64(orderAmount) * taxRate
	total = float64(orderAmount) + tax
	return
}

func printOutput(totoal float64, tax float64) {
	outputMsg := ""
	if tax > 0.0 {
		outputMsg = fmt.Sprintf("The tax is $%.2f.\n", tax)
	}

	outputMsg += fmt.Sprintf("The total is $%.2f.\n", totoal)
	fmt.Print(outputMsg)
}

func main() {
	printOutput(calculateTax(getInput()))
}
