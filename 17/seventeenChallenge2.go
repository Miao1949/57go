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

// Get non-empty input.
func getNonEmptyInput(msg string) (input string) {
	done := false
	for ; !done; {
		fmt.Print(msg)
		input, _ = bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.TrimSpace(input)
		if len(input) > 0 {
			done = true
		}
	}

	return input
}

func getInput() (numberOfStandardDrinks uint64, weight uint64, sex string, hoursSinceLastDrink uint64, state string) {
	numberOfStandardDrinks = getNonNegativeIntegerInput("Number of standard drinks: ")
	weight = getNonNegativeIntegerInput("Weight: ")
	sex = getStringInput("Sex [m|f]",  []string{"m", "f"})
	hoursSinceLastDrink = getNonNegativeIntegerInput("Hours since last drink: ")
	state = getNonEmptyInput("What state do you live in? ")
	return
}

func calculateBac(numberOfDrinks uint64, weight uint64, sex string, hoursSinceLastDrink uint64) (bac float64) {
	bw := 0.58
	if sex == "f" {
		bw = 0.49
	}

	bac = 0.806 * float64(numberOfDrinks) * 1.2 / (float64(weight) * bw)  - 0.017 * float64(hoursSinceLastDrink)

	return
}


func main() {
	numberOfStandardDrinks, weight, sex, hoursSinceLastDrink, state := getInput()
	bac := calculateBac(numberOfStandardDrinks, weight, sex, hoursSinceLastDrink)
	fmt.Printf("Your bac is: %.2f\n", bac)

	stateToBacMap := make(map[string]float64)
	stateToBacMap["WI"] = 0.05

	bacLimit := 0.08

	if limit, ok := stateToBacMap[state]; ok {
		bacLimit = limit
	}

	if bac >= bacLimit{
		fmt.Println("It is not legal for you to drive.")
	} else {
		fmt.Println("It is legal for you to drive.")
	}
}
