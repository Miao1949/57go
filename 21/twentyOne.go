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
		i, err := strconv.ParseUint(input, 10, 64)

		if err == nil {
			retInt = i
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

func displayMonth(monthNumber uint64) {
	monthAsString := ""
	switch monthNumber {
	case 1: monthAsString = "January"
	case 2: monthAsString = "February"
	case 3: monthAsString = "March"
	case 4: monthAsString = "April"
	case 5: monthAsString = "May"
	case 6: monthAsString = "June"
	case 7: monthAsString = "July"
	case 8: monthAsString = "August"
	case 9: monthAsString = "September"
	case 10: monthAsString = "October"
	case 11: monthAsString = "November"
	case 12: monthAsString = "December"
	default:
		panic("Invalid month entered!")
	}

	fmt.Println(monthAsString)
}


func main() {
	displayMonth(getMonthNumber())
}
