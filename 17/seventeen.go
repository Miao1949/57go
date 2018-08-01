package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func getNonNegativeIntegerInput(msg string) (retInt uint64) {
	done := false
	for ; !done; {
		fmt.Print(msg)
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.TrimSpace(input)
		i, err := strconv.ParseUint(input, 10, 32)

		if err == nil {
			retInt = i
			done = true
		}
	}
	return
}


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

func getInput() (alcoholConsumed uint64, weight uint64, sex string, hoursSinceLastDrink uint64) {
	alcoholConsumed = getNonNegativeIntegerInput("Alcohol consumed: ")
	weight = getNonNegativeIntegerInput("Weight: ")
	sex = getStringInput("Sex [m|f]",  []string{"m", "f"})
	hoursSinceLastDrink = getNonNegativeIntegerInput("Hours since last drink: ")
	return
}

func calculateBac(alcoholConsumed uint64, weight uint64, sex string, hoursSinceLastDrink uint64) (bac float64) {
	r := 0.73
	if sex == "f" {
		r = 0.66
	}

	bac = float64(alcoholConsumed) * 5.14 / float64(weight) * r - 0.015 * float64(hoursSinceLastDrink)
	return
}


func main() {
	alcoholConsumed, weight, sex, hoursSinceLastDrink := getInput()
	bac := calculateBac(alcoholConsumed, weight, sex, hoursSinceLastDrink)
	fmt.Printf("Your bac is: %.2f\n", bac)

	if bac >= 0.08 {
		fmt.Println("It is not legal for you to drive.")
	} else {
		fmt.Println("It is legal for you to drive.")
	}
}
