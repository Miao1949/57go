package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"math"
)

func getStringInput(msg string, allowedResponses []string) (retString string) {
	allowedResponsesMap := make(map[string]bool)
	for _, allowedResponse := range allowedResponses {
		allowedResponsesMap[allowedResponse] = true
	}

	return getRestrictedStringInput(msg, allowedResponsesMap)
}

func getRestrictedStringInput(msg string, allowedResponses map[string]bool) (retString string) {
	done := false
	for ; !done; {
		fmt.Print(msg)
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.TrimSpace(input)

		if len(input) > 0 && allowedResponses[input]{
			retString = input
			done = true
		}
	}
	return
}

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

func getHeightAndWeight() (height uint64, weight uint64) {
	height = getNonNegativeIntegerInput("Height:")
	weight = getNonNegativeIntegerInput("Weight: ")
	return
}

func getHeightAndWeightInImperialUnits() (height uint64, weight uint64) {
	feet := getNonNegativeIntegerInput("Feet: ")
	inches := getNonNegativeIntegerInput("Inches: ")
	height = feet * 12 + inches
	weight = getNonNegativeIntegerInput("Weight: ")
	return
}


func calculateBmiUsingImperialUnits(height uint64, weight uint64) (bmi float64) {
	bmi = float64(weight) / math.Pow(float64(height), 2) * 703.0
	return
}

func calculateBmiUsingMetricUnits(height uint64, weight uint64) (bmi float64) {
	bmi = float64(weight) / math.Pow(float64(height) / 100.0, 2)
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
	metricOrImperial := getStringInput("Metric(m) or imperial(i) units? ", []string{"m", "M", "i", "I"})
	if metricOrImperial == "i" || metricOrImperial == "I" {
		printBmiResults(calculateBmiUsingImperialUnits(getHeightAndWeightInImperialUnits()))
	} else {
		printBmiResults(calculateBmiUsingMetricUnits(getHeightAndWeight()))
	}
}
