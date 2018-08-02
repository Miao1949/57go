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
	monthNumberToStringMap := make(map[uint64]string)
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

	fmt.Println(monthNumberToStringMap[monthNumber])
}


func main() {
	displayMonth(getMonthNumber())
}
