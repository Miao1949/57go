package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"math"
)

const SquareFeetCoveredByOneGallon = 350.0

func getIntegerInput(msg string) (retInt int) {
	done := false
	for ; !done; {
		fmt.Print(msg)
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.TrimSpace(input)
		i,err := strconv.Atoi(input)

		if err == nil && i > 0{
			retInt = i
			done = true
		}
	}
	return
}


func main() {
	length := getIntegerInput("Length: ")
	width := getIntegerInput("Width: ")

	area := length * width
	numberOfGallonsNeeded := int(math.Ceil(float64(area) / SquareFeetCoveredByOneGallon))
	fmt.Printf("You need %d gallons of paint to cover %d square feet.\n", numberOfGallonsNeeded, area)
}
