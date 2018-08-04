package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	restingPulse := getPositiveIntergerInput("Please enter your resting pulse: ")
	age := getPositiveIntergerInput("Please enter your age: ")

	fmt.Printf("Resting Pulse: %d\t Age: %d\n", restingPulse, age)
	fmt.Println("Intensity       | Rate")
	fmt.Println("----------------|--------")
	for intensity := 55; intensity <= 95; intensity += 5 {
		targetHeartRate := (220 - age - restingPulse) * intensity / 100  + restingPulse
		fmt.Printf("%d%%\t\t| %d bpm\n", intensity, targetHeartRate)
	}
}


func getPositiveIntergerInput(msg string) (retVal int) {
	done := false
	for !done {
		fmt.Print(msg)
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err == nil {
			i, err2 := strconv.Atoi(strings.TrimSpace(input))
			if err2 == nil {
				retVal = i
				done = true
			}
		}
	}

	return
}