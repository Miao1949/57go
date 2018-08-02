package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func getIntegerInput(msg string) (retInt int, end bool) {
	done := false
	end = false
	for ; !done; {
		fmt.Print(msg)
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.TrimSpace(input)

		if len(input) == 0 {
			end = true
			done = true
		} else {

			i, err := strconv.Atoi(input)

			if err == nil {
				retInt = i
				done = true
			}
		}
	}
	return
}

func getNumbers() (numbers []int) {
	numbers = make([]int, 5)
	numberMap := make(map[int]bool)

	userIndicatedStop := false
	var integerEnteredByUser int
	i := 0
	for !userIndicatedStop {
		i++
		done := false
		for !done {
			msg := fmt.Sprintf("Enter number %d: ", i)
			integerEnteredByUser, userIndicatedStop = getIntegerInput(msg)
			if userIndicatedStop == true {
				done = true
			}

			if userIndicatedStop == false && numberMap[integerEnteredByUser] == false {
				numbers = append(numbers, integerEnteredByUser)
				numberMap[integerEnteredByUser] = true
				done = true
			}
		}
	}

	return
}

func findLargestNumber(numbers []int) (largestNumber int) {
	largestNumber = numbers[0]

	for _, number := range numbers {
		if number > largestNumber {
			largestNumber = number
		}
	}

	return
}

func main() {
	numbers := getNumbers()
	largestNumber := findLargestNumber(numbers)
	fmt.Printf("The largest number is %d.\n", largestNumber)
}
