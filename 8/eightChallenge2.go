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

func determinePluralization(n int) (pieceOrPieces string) {
	if n == 1 {
		pieceOrPieces = "piece"
	} else {
		pieceOrPieces = "pieces"
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

	pieceOrPiecesForNumberEachPersonGets := determinePluralization(numberOfSlicesPerPerson)
	pieceOrPiecesForLeftOverNumber := determinePluralization(numberOfLeftOverSlices)

	fmt.Printf("Each person gets %d %s of pizza.\n", numberOfSlicesPerPerson, pieceOrPiecesForNumberEachPersonGets)
	fmt.Printf("here are %d leftover %s.\n", numberOfLeftOverSlices, pieceOrPiecesForLeftOverNumber)
}
