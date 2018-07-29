package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

// Get integer input.
func getNonNegativeIntegerInput(msg string) (retInt int) {
	done := false
	for ; !done; {
		fmt.Print(msg)
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.TrimSpace(input)
		i,err := strconv.Atoi(input)

		if err == nil && i >= 0 {
			retInt = i
			done = true
		}
	}
	return
}

func main() {
	i1 := getNonNegativeIntegerInput("What is the first number? ")
	i2 := getNonNegativeIntegerInput("What is the second number? ")
	fmt.Printf("%d + %d = %d\n", i1, i2, i1 + i2)
	fmt.Printf("%d - %d = %d\n", i1, i2, i1 - i2)
	fmt.Printf("%d * %d = %d\n", i1, i2, i1 * i2)
	fmt.Printf("%d / %d = %d\n", i1, i2, i1 / i2)
}