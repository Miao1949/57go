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

// Get input from user. Only allow certain responses.
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

func proceedUsingMeter() {
	lengthInMeter := getIntegerInput("What is the length of the room in meter? ")
	widthInMeter := getIntegerInput("What is the width of the room in meter? ")
	fmt.Printf("You entered dimensions of %d meter by %d meter.\n", lengthInMeter, widthInMeter)
	fmt.Println("The area is:")
	areaInSquareMeters := calculateArea(lengthInMeter, widthInMeter)
	fmt.Printf("%d square meters\n", areaInSquareMeters)
	areaInSquareFeet := float64(areaInSquareMeters) / AREA_IN_FT2_TO_M2_CONVERSION_FACTOR
	fmt.Printf("%f square feet\n", areaInSquareFeet)
}

func proceedUsingFeet() {
	lengthInFeet := getIntegerInput("What is the length of the room in feet? ")
	widthInFeet := getIntegerInput("What is the width of the room in feet? ")
	fmt.Printf("You entered dimensions of %d feet by %d feet.\n", lengthInFeet, widthInFeet)
	fmt.Println("The area is:")
	areaInSquareFeet := calculateArea(lengthInFeet, widthInFeet)
	fmt.Printf("%d square feet\n", areaInSquareFeet)
	areaInSquareMeters := AREA_IN_FT2_TO_M2_CONVERSION_FACTOR * float64(areaInSquareFeet)
	fmt.Printf("%f square meters\n", areaInSquareMeters)
}


func main() {
	meterOrFeet := getStringInput("Use meter(m) or feet(f)",  []string{"m", "f"})
	if meterOrFeet == "m" {
		proceedUsingMeter()
	} else {
		proceedUsingFeet()
	}
}