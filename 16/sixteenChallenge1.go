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
		i, err := strconv.ParseUint(input, 10, 64)

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
	age := getInput()
	outputMessage := ""
	if age >= 16 {
		outputMessage = "You are old enough to legally drive."
	} else {
		outputMessage = "You are not old enough to legally drive."
	}

	fmt.Println(outputMessage)
}
