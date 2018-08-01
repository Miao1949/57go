package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

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

func fahrenheitToCelsius(degreesFahrenheit uint64) (degreesCelsius float64) {
	degreesCelsius = (float64(degreesFahrenheit) - 32.0) * 5.0 / 9.0
	return
}

func celsiusToFahrenheit(degreesCelsius uint64) (degreesFahrenheit float64) {
	degreesFahrenheit = (float64(degreesCelsius) * 9.0 / 5.0) + 32
	return
}


func processsFahrenheitToCelsius() {
	degreesFahrenheit := getNonNegativeIntegerInput("Please enter the temperature in Fahrenheit: ")
	degreesCelsius := fahrenheitToCelsius(degreesFahrenheit)
	fmt.Printf("The temperature in Celsius is %.2f.\n", degreesCelsius)
}

func processsCelsiusToFahrenheit() {
	degreesCelsius := getNonNegativeIntegerInput("Please enter the temperature in Celsius: ")
	degreesFahrenheit := celsiusToFahrenheit(degreesCelsius)
	fmt.Printf("The temperature in Fahrenheitis %.2f\n", degreesFahrenheit)
}

func main() {
	cOrF := getStringInput("Press C to convert from Fahrenheit to Celsius.\nPress F to convert from Celsius to Fahrenheit.\nYour choise: ", []string{"C", "c", "F", "f"})
	fmt.Println("You chose:", cOrF)
	if cOrF == "C" || cOrF == "c" {
		processsFahrenheitToCelsius()
	} else {
		processsCelsiusToFahrenheit()
	}
}
