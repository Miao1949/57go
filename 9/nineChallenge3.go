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
	lengthOfLongestSide := getIntegerInput("Length of longest side: ")
	lengthOfShortestSide := getIntegerInput("Length of shortest side: ")
	widthOfLongestSide := getIntegerInput("Width of longest side: ")
	widthOfShortestSide := getIntegerInput("Width of shortest side: ")

	effectiveLenghtOfShortestSide := lengthOfShortestSide - widthOfLongestSide

	areaOfLongestPart := lengthOfLongestSide * widthOfLongestSide
	areaOfShortestPart := effectiveLenghtOfShortestSide * widthOfShortestSide
	totalArea := areaOfLongestPart + areaOfShortestPart

	numberOfGallonsNeeded := int(math.Ceil(float64(totalArea) / SquareFeetCoveredByOneGallon))
	fmt.Printf("You need %d gallons of paint to cover %d square feet.\n", numberOfGallonsNeeded, totalArea)
}
