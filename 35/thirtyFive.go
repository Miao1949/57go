package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"math/rand"
	"time"
)

func main() {
	printWinner(pickWinner(getNames()))
}

func printWinner(winner string) {
	fmt.Printf("The winner is... %s.\n", winner)
}

func pickWinner(names []string) (winner string) {
	randomNumberGenerator := rand.New(rand.NewSource(time.Now().Unix()))
	winner = names[randomNumberGenerator.Intn(len(names))]
	return
}

func getNames() (names []string) {
	done := false
	names = make([]string, 0)
	for !done {
		name := strings.TrimSpace(getStringUtnilEmpty("Enter a name: "))
		if len(name) > 0 {
			names = append(names, name)
		} else {
			done = true
		}
	}

	return
}

func getStringUtnilEmpty(msg string) (input string) {
	fmt.Print(msg)
	input, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	return
}
