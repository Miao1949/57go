package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

const EN = "EN"
const SE = "SE"

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

func getMonthNumber() (monthNumber uint64) {
	done := false
	for !done {
		monthNumber = getNonNegativeIntegerInput("Please enter the number of the month: ")
		if monthNumber < 1 || monthNumber > 12 {
			fmt.Println("Please provide a month number in the [1, 12] range.")
		} else {
			done = true
		}
	}

	return
}

func getInput() (monthNumber uint64, language string) {
	language = getStringInput("Please enter the language to use (EN|SE): ", []string{EN, SE})
	monthNumber = getMonthNumber()
	return
}

func createMonthMap(language string) (monthNumberToStringMap map[uint64]string) {
	monthNumberToStringMap = make(map[uint64]string)

	if language == SE {
		monthNumberToStringMap[1] = "Januari"
		monthNumberToStringMap[2] = "Februari"
		monthNumberToStringMap[3] = "Mars"
		monthNumberToStringMap[4] = "April"
		monthNumberToStringMap[5] = "Maj"
		monthNumberToStringMap[6] = "Juni"
		monthNumberToStringMap[7] = "Juli"
		monthNumberToStringMap[8] = "Augusti"
		monthNumberToStringMap[9] = "September"
		monthNumberToStringMap[10] = "Oktober"
		monthNumberToStringMap[11] = "November"
		monthNumberToStringMap[12] = "December"

	} else {
		monthNumberToStringMap[1] = "January"
		monthNumberToStringMap[2] = "February"
		monthNumberToStringMap[3] = "March"
		monthNumberToStringMap[4] = "April"
		monthNumberToStringMap[5] = "May"
		monthNumberToStringMap[6] = "June"
		monthNumberToStringMap[7] = "July"
		monthNumberToStringMap[8] = "August"
		monthNumberToStringMap[9] = "September"
		monthNumberToStringMap[10] = "October"
		monthNumberToStringMap[11] = "November"
		monthNumberToStringMap[12] = "December"

	}

	return
}

func displayMonth(monthNumber uint64, monthNumberToStringMap map[uint64]string) {
	fmt.Println(monthNumberToStringMap[monthNumber])
}


func main() {
	monthNumber, language := getInput()
	monthNumberToStringMap := createMonthMap(language)
	displayMonth(monthNumber, monthNumberToStringMap)
}
