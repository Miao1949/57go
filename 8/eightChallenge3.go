package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"math"
)

func getIntegerInput(msg string) (retInt int) {
	done := false
	for ; !done; {
		fmt.Print(msg)
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.TrimSpace(input)
		i,err := strconv.Atoi(input)

		if err == nil && i >= 0 {
			retInt = i
			done = true
		}
	}
	return
}


func main() {
	numberOfPeople := getIntegerInput("How many people? ")
	numberOfSlicesPerPizza := getIntegerInput("How many slices per pizza? ")
	desiredNumberOfPiecesPerPerson := getIntegerInput("How many pieces does each person want? ")


	totalNumberOfSlices := desiredNumberOfPiecesPerPerson  * numberOfPeople
	numberOfNeededPizzas := int(math.Ceil(float64(totalNumberOfSlices) / float64(numberOfSlicesPerPizza)))
	numberOfLeftOverSlices := numberOfNeededPizzas * numberOfSlicesPerPizza - desiredNumberOfPiecesPerPerson * numberOfPeople

	fmt.Printf("You need to buy %d pizzas.\n", numberOfNeededPizzas)
	fmt.Printf("There will be %d pices of pizza over.\n", numberOfLeftOverSlices)
}
