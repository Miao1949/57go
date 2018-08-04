package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"math/rand"
	"time"
)

func main() {
	done := false
	for !done {
		playRound()
		playAgain := getStringInput("Play again? ", []string{"y", "Y", "n", "N"})
		if playAgain == "n" || playAgain == "N" {
			fmt.Println("Goodbye!")
			done = true
		}
	}

}

func playRound() {

	fmt.Println("Let's play Guess the Number.")
	difficultyLevelAsString := getStringInput("Pick a difficulty level (1, 2, or 3): ", []string{"1", "2", "3"})
	difficultyLevel, _ := strconv.Atoi(difficultyLevelAsString)
	var highLimit int
	if difficultyLevel == 1 {
		highLimit = 10
	} else if difficultyLevel == 2 {
		highLimit = 100
	} else {
		highLimit = 1000
	}

	numberOfGuesses := 0
	randomNumberGenerator := rand.New(rand.NewSource(time.Now().Unix()))
	correctNumber := 1 + randomNumberGenerator.Intn(highLimit - 1)
	guessesMap := make(map[int]bool)

	guess, numberOfNonIntegerEntries := getPositiveIntergerInput("I have my number. What's your guess? ")
	numberOfGuesses++
	numberOfGuesses += numberOfNonIntegerEntries
	done := false
	for !done {
		if guess == correctNumber {
			fmt.Printf("You got it in %d guesses!\n%s\n", numberOfGuesses, generateTextForNumberOfGuesses(numberOfGuesses))
			done = true
		} else if _, ok := guessesMap[guess]; ok {
			fmt.Printf("You have already guessed %d. Pick another number.\n", guess)
			numberOfGuesses++
			guess, numberOfNonIntegerEntries = getPositiveIntergerInput("Guess again: ")
			numberOfGuesses++
			numberOfGuesses += numberOfNonIntegerEntries
		} else if guess < correctNumber {
			guessesMap[guess] = true
			guess, numberOfNonIntegerEntries = getPositiveIntergerInput("Too low. Guess again: ")
			numberOfGuesses++
			numberOfGuesses += numberOfNonIntegerEntries
		} else {
			guessesMap[guess] = true
			guess, numberOfNonIntegerEntries = getPositiveIntergerInput("Too high. Guess again: ")
			numberOfGuesses++
			numberOfGuesses += numberOfNonIntegerEntries
		}
	}
}

func generateTextForNumberOfGuesses(numberOfGuesses int) (retVal string) {
	if numberOfGuesses == 1 {
		retVal = "You’re a mind reader"
	} else if numberOfGuesses <= 4 {
		retVal = "Most impressive."
	} else if numberOfGuesses <= 6 {
		retVal = "You can do better than that."
	} else {
		retVal = "You can do better than that."
	}

	return
}

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

func getPositiveIntergerInput(msg string) (retVal int, numberOfNonIntegerEntries int) {
	done := false
	numberOfNonIntegerEntries = 0
	for !done {
		fmt.Print(msg)
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err == nil {
			i, err2 := strconv.Atoi(strings.TrimSpace(input))
			if err2 == nil {
				retVal = i
				done = true
			} else {
				numberOfNonIntegerEntries++
			}
		}
	}

	return
}