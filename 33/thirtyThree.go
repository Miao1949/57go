package main

import (
	"math/rand"
	"time"
	"fmt"
	"bufio"
	"os"
)

func main() {
	answers := []string{"Yes", "No", "Maybe", "Ask again later"}

	randomNumberGenerator := rand.New(rand.NewSource(time.Now().Unix()))
	getStringInput("What's your question? ")
	answer := answers[randomNumberGenerator.Intn(len(answers))]
	fmt.Println(answer)
}

func getStringInput(msg string) (input string) {
	done := false
	for !done {
		fmt.Print(msg)
		s, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err == nil {
			input = s
			done = true
		}
	}

	return
}
