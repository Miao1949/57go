package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

const TaxRate = 0.055

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

func getInput() (p1, q1, p2, q2, p3, q3 int){
	p1 = getNonNegativeIntegerInput("Enter the price of item 1: ")
	q1 = getNonNegativeIntegerInput("Enter the quantity of item 1: ")
	p2 = getNonNegativeIntegerInput("Enter the price of item 2: ")
	q2 = getNonNegativeIntegerInput("Enter the quantity of item 1: ")
	p3 = getNonNegativeIntegerInput("Enter the price of item 3: ")
	q3 = getNonNegativeIntegerInput("Enter the quantity of item 1: ")

	return
}

func calculateSubtotalTaxAndTotal(p1, q1, p2, q2, p3, q3 int) (subtotal int, tax, total float64) {
	subtotal = p1 * q1 + p2 * q2 + p3 * q3
	tax = float64(subtotal) * TaxRate
	total = float64(subtotal) + tax
	return
}

func printResult(subtotal int, tax, total float64) {
	fmt.Printf("Subtotal: $%.2f\n", float64(subtotal))
	fmt.Printf("Tax: $%.2f\n", tax)
	fmt.Printf("Total: $%.2f\n", total)
}

func main() {
	printResult(calculateSubtotalTaxAndTotal(getInput()))
}
