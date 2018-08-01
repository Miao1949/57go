package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func getNonNegativeIntegerInput(msg string, errorMsg string) (retInt uint64) {
	done := false
	for ; !done; {
		fmt.Print(msg)
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.TrimSpace(input)
		i, err := strconv.ParseUint(input, 10, 32)

		if err == nil {
			retInt = i
			done = true
		} else {
			fmt.Println(errorMsg)
		}
	}
	return
}

func getInput() (age uint64) {
	age = getNonNegativeIntegerInput("What is you age? ", "Please enter a valid age!")
	return
}

func main() {
	countryToDrivingAgeMap := make(map[string]uint64)
	countryToDrivingAgeMap["US"] = 16
	countryToDrivingAgeMap["SWE"] = 18


	age := getInput()
	for country, ageForCountry := range countryToDrivingAgeMap {
		if ageForCountry <= age {
			fmt.Println("You may drive in ", country)
		}
	}
}
