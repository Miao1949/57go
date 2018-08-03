package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {
	numberOfNumbers := getNonNegativeIntegerInput("How many numbers do you want to add? ")

	fmt.Println(numberOfNumbers)
	var sum uint64 = 0
	var i uint64
	for i = 1; i <= numberOfNumbers; i++ {
		fmt.Print("Enter a number: ")
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.TrimSpace(input)
		ii, err := strconv.ParseUint(input, 10, 64)
		if err == nil {
			sum += ii
		}
	}

	fmt.Printf("The total is %d\n", sum)
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
