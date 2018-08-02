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

func getNumbers(numberOfNumbers int) (numbers []int) {
	numbers = make([]int, 5)
	numberMap := make(map[int]bool)
	for i := 1; i <= numberOfNumbers; i++ {
		done := false
		for !done {
			msg := fmt.Sprintf("Enter number %d: ", i)
			anInt := getIntegerInput(msg)
			if numberMap[anInt] == false {
				numbers = append(numbers, anInt)
				numberMap[anInt] = true
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
	numbers := getNumbers(3)
	largestNumber := findLargestNumber(numbers)
	fmt.Printf("The largest number is %d.\n", largestNumber)
}
