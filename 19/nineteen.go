package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"math"
)

func getNonNegativeIntegerInput(msg string) (retInt uint64) {
	done := false
	for ; !done; {
		fmt.Print(msg)
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.TrimSpace(input)
		i, err := strconv.ParseUint(input, 10, 64)

		if err == nil {
			retInt = i
			done = true
		}
	}
	return
}

func getInput() (height uint64, weight uint64) {
	height = getNonNegativeIntegerInput("Height:")
	weight = getNonNegativeIntegerInput("Weight: ")
	return
}


func calculateBmi(height uint64, weight uint64) (bmi float64) {
	bmi = float64(weight) / math.Pow(float64(height), 2) * 703.0
	return
}

func printBmiResults(bmi float64) {
	fmt.Printf("Your BMI is %.1f.\n", bmi)
	if bmi < 18.5 {
		fmt.Println("You are underweight")
	} else if bmi > 25.0 {
		fmt.Println("You are overweight. You should see your doctor.")
	} else {
		fmt.Println("You are within the ideal weight range.")
	}
}


func main() {
	printBmiResults(calculateBmi(getInput()))
}
