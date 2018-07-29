package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

//What is the length of the room in feet? 15
//What is the width of the room in feet? 20
//You entered dimensions of 15 feet by 20 feet.
//The area is
//300 square feet
//27.871 square meters


const AREA_IN_FT2_TO_M2_CONVERSION_FACTOR = 0.09290304

func getIntegerInput(msg string) (retInt int) {
	done := false
	for ; !done; {
		fmt.Print(msg)
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.TrimSpace(input)
		i,err := strconv.Atoi(input)

		if err == nil {
			retInt = i
			done = true
		}
	}
	return
}

func calculateArea(length, width int) int {
	return length * width
}

func main() {
	lengthInFeet := getIntegerInput("What is the length of the room in feet? ")
	widthInFeet := getIntegerInput("What is the width of the room in feet? ")
	fmt.Printf("You entered dimensions of %d feet by %d feet.\n", lengthInFeet, widthInFeet)
	fmt.Println("The area is:")
	areaInSquareFeet := calculateArea(lengthInFeet, widthInFeet)
	fmt.Printf("%d square feet\n", areaInSquareFeet)
	areaInSquareMeters := AREA_IN_FT2_TO_M2_CONVERSION_FACTOR * float64(areaInSquareFeet)
	fmt.Printf("%f square meters\n", areaInSquareMeters)
}