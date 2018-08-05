package main

import (
	"bufio"
	"os"
	"strings"
		"fmt"
	"strconv"
	"io"
)

const Filename = "numbers.txt"

func main() {
	evenNumbers, err := processFile()
	if err == nil {
		fmt.Print("The even numbers are ")
		for _, evenNumber := range evenNumbers {
			fmt.Print(evenNumber)
			fmt.Print(" ")
		}
		fmt.Println()
	} else {
		fmt.Println("Error when processing file!")
	}
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

func processFile() (evenNumbers []int, err error){
	evenNumbers = make([]int, 0)
	file, e := os.Open(Filename)
	if e != nil {
		fmt.Println("Could not open file")
		err = e
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	done := false
	for !done {
		s, e := reader.ReadString('\n')
		if e == nil || e == io.EOF {
			i, e2 := strconv.Atoi(strings.TrimSpace(s))
			if e2 == nil && i % 2 == 0 {
				evenNumbers = append(evenNumbers, i)
			}
		}

		if e != nil {
			done = true
		}
	}

	return

}
