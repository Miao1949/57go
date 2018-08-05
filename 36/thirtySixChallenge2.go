package main

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
	"strings"
	"math"
	"io"
)

func main()  {
	numbers, err := getNumbers()
	if  err == nil {
		printStatistics(calculateStatistics(numbers))
	} else {
		fmt.Println(err)
	}
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


func getNumbers() (numbers []int, error error) {
	file, err := os.Open("numbers.txt")
	if err != nil {
		fmt.Println("Could not open file!")
		error = err
		return
	}
	defer file.Close()

	done := false
	reader := bufio.NewReader(file)
	for !done {
		s, e := reader.ReadString('\n')
		if e == nil || e == io.EOF{
			s = strings.TrimSpace(s)

			n, err := strconv.Atoi(s)
			if err == nil {
				numbers = append(numbers, n)
			}
		}

		if e == io.EOF {
			done = true
		} else if e != nil {
			fmt.Println("Error occured when reading file!", e)
			error = e
			return
		}
	}
	return
}