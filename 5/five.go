package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

// Get integer input.
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
	i1 := getIntegerInput("What is the first number? ")
	i2 := getIntegerInput("What is the second number? ")
	fmt.Printf("%d + %d = %d\n", i1, i2, i1 + i2)
	fmt.Printf("%d - %d = %d\n", i1, i2, i1 - i2)
	fmt.Printf("%d * %d = %d\n", i1, i2, i1 * i2)
	fmt.Printf("%d / %d = %d\n", i1, i2, i1 / i2)
}