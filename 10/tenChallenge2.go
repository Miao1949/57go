package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

const TaxRate = 0.055

type PriceAndQuantity struct {
	price, quantity int
}


func getNonNegativeIntegerInput(msg string) (retInt int, end bool) {
	done := false
	end = false
	for ; !done; {
		fmt.Print(msg)
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.TrimSpace(input)

		if len(input) == 0 {
			end = true
			done = true
		} else {

			i, err := strconv.Atoi(input)

			if err == nil && i > 0 {
				retInt = i
				done = true
			}
		}
	}
	return
}

func getInput() (priceAndQuantities []PriceAndQuantity){
	priceAndQuantities = make([]PriceAndQuantity, 1)
	done := false
	counter := 1
	for !done {
		priceMessage := fmt.Sprintf("Enter the price of item %d: ", counter)
		quantityMessage := fmt.Sprintf("Enter the quantity of item %d: ", counter)

		price, end := getNonNegativeIntegerInput(priceMessage)
		if !end {
			quantity, end := getNonNegativeIntegerInput(quantityMessage)
			if !end {
				priceAndQuantities = append(priceAndQuantities, PriceAndQuantity{price, quantity})
			}
		}

		counter++
		done = end
	}

	return
}

func calculateSubtotalTaxAndTotal(priceAndQuantities []PriceAndQuantity) (subtotal int, tax, total float64) {
	for _,p := range priceAndQuantities {
		subtotal += p.price * p.quantity
	}

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
