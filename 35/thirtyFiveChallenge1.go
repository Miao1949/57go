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
	done := false
	names := getNames()
	for !done && len(names) > 0 {
		winner := pickWinner(names)
		printWinner(winner)

		if getYesOrNoInput("Play again? (y/n) ") {
			names = removeWinner(names, winner)
		} else {
			done = true
		}
	}
}

func removeWinner(names []string, winner string) (namesWithTheWinnerRemoved []string) {
	for _, name := range names {
		if name != winner {
			namesWithTheWinnerRemoved = append(namesWithTheWinnerRemoved, name)
		}
	}

	return
}

func getYesOrNoInput(msg string) (answer bool) {
	yesOrNoOptions :=  make(map[string]bool)
	yesOrNoOptions["y"] = true
	yesOrNoOptions["Y"] = true
	yesOrNoOptions["n"] = true
	yesOrNoOptions["N"] = true
	retString := getRestrictedStringInput(msg, yesOrNoOptions)
	answer = retString == "y" || retString == "Y"
	return
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
