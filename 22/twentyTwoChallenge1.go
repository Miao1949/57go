package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
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
	i1 := getIntegerInput("Enter the first number: ")

	done := false
	var i2, i3 int

	for !done {
		i2 = getIntegerInput("Enter the second number: ")
		if i2 != i1 {
			done = true
		}
	}

	done = false
	for !done {
		i3 = getIntegerInput("Enter the third number: ")
		if i3 != i1 && i3 != i2 {
			done = true
		}
	}

	var largestNumber int
	if i1 > i2 {
		if i1 > i3 {
			largestNumber = i1
		} else {
			largestNumber = i3
		}
	} else if i2 > i3 {
		largestNumber = i2
	} else {
		largestNumber = i3
	}

	fmt.Printf("The largest number is %d.\n", largestNumber)
}
