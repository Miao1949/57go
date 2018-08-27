package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

//---------------------------------------------------------
// CONSTANTS
//---------------------------------------------------------
const triviaFilename = "trivia.json"

//---------------------------------------------------------
// TYPES
//---------------------------------------------------------
type questions struct {
	Questions []question
}
type question struct {
	Question string
	Answer string
	Distractors []string
}

//---------------------------------------------------------
// MAIN
//---------------------------------------------------------

func main() {
	questions := readDataFromFile()
	runGame(questions.Questions)
}

//---------------------------------------------------------
// PRIVATE METHODS.
//---------------------------------------------------------

func readDataFromFile() (questions questions) {
	fileContent, e := ioutil.ReadFile(triviaFilename)
	if e != nil {
		fmt.Fprint(os.Stderr, "Could not read file %s\n", triviaFilename)
		panic(e)
	}

	if err := json.Unmarshal(fileContent, &questions); err != nil {
		fmt.Println("Could not urmarshal file content as json!")
		return
	}

	return
}

func runGame(questions []question) {
	randomNumberGenerator := rand.New(rand.NewSource(time.Now().Unix()))

	gameOver := false
	userWon := false
	numberOfCorrectlyAnsweredQuestions := 0
	userAnsweredCorrectly := false
	for !gameOver {
		userAnsweredCorrectly, questions = displayRandomQuestionsAndLetUserAnswerIt(questions, randomNumberGenerator)
		if !userAnsweredCorrectly {
			gameOver = true
			fmt.Println("That was unfortunately not correct :(")
		} else {
			numberOfCorrectlyAnsweredQuestions++
		}

		if len(questions) == 0 {
			if !gameOver { // Did the user give a correct answer to the final question?
				userWon = true
			}
			gameOver = true
		}
	}

	if userWon {
		fmt.Println("You won!")
	} else {
		fmt.Println("You lost!")
	}
	fmt.Printf("Number of correct answers: %d\n", numberOfCorrectlyAnsweredQuestions)
}

func displayRandomQuestionsAndLetUserAnswerIt(questions []question, randomNumberGenerator *rand.Rand) (userAnsweredCorrectly bool, remainingQuestions []question) {
	// Pick a question.
	indexOfQuestion := randomNumberGenerator.Intn(len(questions))
	questionToAnswer := questions[indexOfQuestion]

	// Generated the alternatives.
	alternatives := generateAlternativesToQuestion(questionToAnswer, randomNumberGenerator)

	// Print the question.
	fmt.Printf("\nQuestion: %s\n", questionToAnswer.Question)

	// Print the alternatives.
	for index, alternative := range alternatives {
		fmt.Printf("%d: %s\n", index + 1, alternative)
	}

	// Get input from user.
	answer := getStringInput("Your choice: ", len(alternatives))

	// Convert answer to an index to be able to fetch actual answer from the alternatives slices.
	answer = strings.TrimSpace(answer)
	i, _:= strconv.Atoi(answer)
	answerIndex := i - 1

	// Check if the selected answer is correctly.
	if alternatives[answerIndex] == questionToAnswer.Answer {
		userAnsweredCorrectly = true
	}

	// Remove the question that the user just answered.
	if len(questions) > 1 {
		copy(questions[indexOfQuestion:], questions[indexOfQuestion+1:])
		questions = questions[0 : len(questions)-1]
	} else {
		questions = nil // To indicate that no questions remain.
	}

	// Return flag indicating if the user answered correctly or not, and the slice of remaining questions.
	return userAnsweredCorrectly, questions
}

func generateAlternativesToQuestion(question question, randomNumberGenerator *rand.Rand) (alternatives []string){
	alternatives = make([]string, len(question.Distractors) + 1)
	srcAlternatives := make([]string, len(question.Distractors) + 1)
	copy(srcAlternatives, question.Distractors)
	srcAlternatives[len(srcAlternatives) - 1] = question.Answer

	indexToInsertInto := 0
	for len(srcAlternatives) > 0 {
		indexToPickFrom := randomNumberGenerator.Intn(len(srcAlternatives))
		alternatives[indexToInsertInto] = srcAlternatives[indexToPickFrom]
		indexToInsertInto++

		if len(srcAlternatives) > 1 {
			copy(srcAlternatives[indexToPickFrom:], srcAlternatives[indexToPickFrom + 1:])
			srcAlternatives = srcAlternatives[: len(srcAlternatives) - 1]
		} else {
			srcAlternatives = nil
		}

	}

	return
}

// Get input from user. Only allow responses in the range [1, maxChoice]
func getStringInput(msg string, maxChoice int) (retString string) {
	allowedResponsesMap := make(map[string]bool)
	i := 1
	for i <= maxChoice {
		allowedResponsesMap[strconv.Itoa(i)] = true
		i++
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
