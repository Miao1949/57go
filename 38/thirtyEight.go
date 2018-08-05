package main

import (
	"bufio"
	"os"
	"strings"
		"fmt"
	"strconv"
)

func main() {
	s := getInput()
	evenNumbers := filterEvenNumbers(s)
	fmt.Print("The even numbers are ")
	for _, evenNumber := range evenNumbers {
		fmt.Print(evenNumber)
		fmt.Print(" ")
	}
	fmt.Println()
}

func getInput() (s string) {
	fmt.Print("Enter a list of numbers, separated by spaces: ")
	readString, e := bufio.NewReader(os.Stdin).ReadString('\n')
	if e == nil {
		s = readString
	}
	return
}

func filterEvenNumbers(inputtedNumber string) (evenNumbers []int) {
	evenNumbers = make([]int, 0)
	ss := strings.Split(inputtedNumber, " ")
	for _, s := range ss {
		i, e := strconv.Atoi(strings.TrimSpace(s))
		if e == nil && i % 2 == 0 {
			evenNumbers = append(evenNumbers, i)
		}
	}

	return
}