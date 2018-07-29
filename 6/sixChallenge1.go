package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"time"
)

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

func main() {
	age := getIntegerInput("What is your current age? ")
	retirementAge := getIntegerInput("At what age do you want to retire? ")

	yearsUntilRetirement := retirementAge - age

	if yearsUntilRetirement <= 0 {
		fmt.Println("You can retire now!")
	} else {
		fmt.Printf("You have %d years left until you can retire.\n", yearsUntilRetirement)

		now := time.Now()
		currentYear := now.Year()
		fmt.Printf("It's %d so you can retire in %d\n", currentYear, currentYear + yearsUntilRetirement)
	}
}
