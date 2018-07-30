package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
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
	numberOfPizzas := getIntegerInput("How many pizzas? ")
	numberOfSlicesPerPizza := getIntegerInput("How many slices per pizza? ")

	totalNumberOfSlices := numberOfPizzas  * numberOfSlicesPerPizza
	numberOfSlicesPerPerson := totalNumberOfSlices / numberOfPeople
	numberOfLeftOverSlices := totalNumberOfSlices % numberOfPeople
	fmt.Printf("Each person gets %d pieces of pizza.\n", numberOfSlicesPerPerson)
	fmt.Printf("here are %d leftover pieces.\n", numberOfLeftOverSlices)
}
