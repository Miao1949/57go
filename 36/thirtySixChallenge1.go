package main

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
	"strings"
	"math"
)

func main()  {
	printStatistics(calculateStatistics(getInput()))
}

func printStatistics(average float64, max int, min int, standardDeviation float64) {
	fmt.Printf("The average is %.2f\n", average)
	fmt.Printf("The min is %d\n", min)
	fmt.Printf("The max is %d\n", max)
	fmt.Printf("The standard deviation is %.2f\n", standardDeviation)
}

func calculateStatistics(numbers []int) (average float64, max int, min int, standardDeviation float64) {
	min = calculateMin(numbers)
	max = calculateMax(numbers)
	average = calculateAverage(numbers)
	standardDeviation = calculateStandardDeviation(numbers, average)
	return
}

func calculateMax(numbers []int) (max int) {
	max = numbers[0]
	for _, n := range numbers {
		if n > max {
			max = n
		}
	}
	return
}

func calculateMin(numbers []int) (min int) {
	min = numbers[0]
	for _, n := range numbers {
		if n < min {
			min = n
		}
	}
	return
}

func calculateAverage(numbers []int) (average float64) {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	average = float64(sum) / float64(len(numbers))
	return
}

func calculateStandardDeviation(numbers []int, average float64) (standardDeviation float64) {
	sumOfSquaredDifferencesFromAverage := 0.0
	for _, number := range numbers {
		differanceFromAverage := average - float64(number)
		sumOfSquaredDifferencesFromAverage += differanceFromAverage * differanceFromAverage

	}

	standardDeviation = math.Sqrt(sumOfSquaredDifferencesFromAverage / float64(len(numbers)))
	return
}


func getInput() (numbers []int) {
	done := false
	for !done {
		fmt.Print("Enter a number: ")
		s, e := bufio.NewReader(os.Stdin).ReadString('\n')
		if e == nil {
			s = strings.TrimSpace(s)
			if s == "done" {
				done = true
			} else {
				n, err := strconv.Atoi(s)
				if err == nil {
					numbers = append(numbers, n)
				}
			}
		}
	}
	return
}