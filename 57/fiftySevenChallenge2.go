package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
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
const triviaFilename = "triviaChallenge2.json"

const playGameChoice = "1"
const addQuestionChoice = "2"
const editQuestionChoice = "3"
const removeQuestionChoice = "4"
const displayQuestionsChoice = "5"
const quitChoice = "q"
const choises = "12345q"


const menuText = `
*************************
*       MAIN MENU       *
*************************
1) Play game.
2) Add question.
3) Edit question
4) Remove question
5) Display questions
q) Quit
`

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
	Level int
}

//---------------------------------------------------------
// MAIN
//---------------------------------------------------------

func main() {
	done := false

	for !done {
		printMenu()
		choice := getMenuChoiceFromUser()
		done = actOnChoice(choice)
	}
	printMenu()
}

//---------------------------------------------------------
// PRIVATE METHODS.
//---------------------------------------------------------

func playGame(questions questions) {
	runGame(questions.Questions)
}

func printMenu() {
	fmt.Print(menuText)
}

func getMenuChoiceFromUser() (choice string) {
	reader := bufio.NewReader(os.Stdin)
	done := false
	for !done {
		fmt.Print("Choice: ")
		line, err := reader.ReadString('\n')
		if err == nil {
			line = strings.TrimSpace(line)
			if len(line) == 1 && strings.Contains(choises, line) {
				choice = line
				done = true
			}
		}
	}
	return
}

func actOnChoice(choice string) (quit bool){
	questions := readDataFromFile()
	switch choice {
	case playGameChoice: playGame(questions)
	case addQuestionChoice:  addQuestion(questions)
	case editQuestionChoice: editQuestion(questions)
	case removeQuestionChoice:  removeQuestion(questions)
	case displayQuestionsChoice: displayQuestions(questions, true)
	case quitChoice: quit = true
	default:
		panic("Unsupported choice!")
	}

	return
}

func addQuestion(questions questions) {
	q := createQuestionFromUserInput()
	questions.Questions = append(questions.Questions,  q)
	writeToFile(questions)
}

func createQuestionFromUserInput() (q question) {
	questionText := getNonEmptyInputFromUser("Question: ")
	answerText := getNonEmptyInputFromUser("Answer: ")
	distractor1 := getNonEmptyInputFromUser("Distractor 1: ")
	distractor2 := getNonEmptyInputFromUser("Distractor 2: ")
	distractor3 := getNonEmptyInputFromUser("Distractor 3: ")
	level := getIntFromUser("Level: ")
	return question{
		Question: questionText,
		Answer: answerText,
		Distractors: []string{distractor1, distractor2, distractor3},
		Level:level}
}

func editQuestion(questions questions) {
	if len(questions.Questions) == 0 {
		fmt.Println("*** There is no question to edit! ***")
		return
	}

	displayQuestions(questions, false)
	choice := getChoiceFromUser("Enter the number of the question to edit: ", len(questions.Questions)+1)
	choiceAsInt, _ := strconv.Atoi(choice)
	indexOfQuestionToEdit := choiceAsInt - 1
	newQuestion := createQuestionFromUserInput()
	questions.Questions[indexOfQuestionToEdit] = newQuestion
	writeToFile(questions)
}

func removeQuestion(questions questions) {
	if len(questions.Questions) == 0 {
		fmt.Println("*** There is no question to remove! ***")
		return
	}

	displayQuestions(questions, false)
	choice := getChoiceFromUser("Enter the number of the question to remove: ", len(questions.Questions)+1)
	choiceAsInt, _ := strconv.Atoi(choice)
	indexOfQuestionToRemove := choiceAsInt - 1

	copy(questions.Questions[indexOfQuestionToRemove:], questions.Questions[indexOfQuestionToRemove + 1 :])
	questions.Questions = questions.Questions[0 : len(questions.Questions)-1]

	fmt.Println("After removal the questions look like this: ", questions.Questions)


	writeToFile(questions)
}

func displayQuestions(questions questions, displayAdditionalInfo bool) {
	for index, question := range questions.Questions {
		fmt.Printf("%d: %s\n", index+1, question.Question)
		if displayAdditionalInfo {
			fmt.Printf("Answer: %s Level:%d\n", question.Answer, question.Level)
			fmt.Printf("D1: %s D2: %s D3: %s\n\n", question.Distractors[0], question.Distractors[1], question.Distractors[2])
		}
	}
}

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

func writeToFile(questions questions) {
	marshalledJsonData, err := json.Marshal(questions)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not marshal questions to json! Error: %v\n", err)
		return
	}

	file, err := os.OpenFile(triviaFilename, os.O_WRONLY | os.O_TRUNC, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open file %s. Error: %v\n", triviaFilename, err)
		return

	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.Write(marshalledJsonData)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not write to file %v\n", err)
		return

	}
	writer.Flush()
}


func runGame(questions []question) {
	randomNumberGenerator := rand.New(rand.NewSource(time.Now().Unix()))

	gameOver := false
	userWon := false
	numberOfCorrectlyAnsweredQuestions := 0
	userAnsweredCorrectly := false
	level := 1
	for !gameOver {
		userAnsweredCorrectly, questions = displayRandomQuestionsAndLetUserAnswerIt(questions, randomNumberGenerator, level)
		if !userAnsweredCorrectly {
			gameOver = true
			fmt.Println("That was unfortunately not correct :(")
		} else {
			numberOfCorrectlyAnsweredQuestions++
			level++
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

func displayRandomQuestionsAndLetUserAnswerIt(questions []question, randomNumberGenerator *rand.Rand, level int) (userAnsweredCorrectly bool, remainingQuestions []question) {
	// Pick a question.
	questionsToConsider := findAllQuestionsOfLevel(questions, level)
	indexOfQuestion := randomNumberGenerator.Intn(len(questionsToConsider))
	questionToAnswer := questionsToConsider[indexOfQuestion]

	// Generated the alternatives.
	alternatives := generateAlternativesToQuestion(questionToAnswer, randomNumberGenerator)

	// Print the question.
	fmt.Printf("\nQuestion: %s\n", questionToAnswer.Question)

	// Print the alternatives.
	for index, alternative := range alternatives {
		fmt.Printf("%d: %s\n", index + 1, alternative)
	}

	// Get input from user.
	answer := getChoiceFromUser("Your choice: ", len(alternatives))

	// Convert answer to an index to be able to fetch actual answer from the alternatives slices.
	answer = strings.TrimSpace(answer)
	i, _:= strconv.Atoi(answer)
	answerIndex := i - 1

	// Check if the selected answer is correctly.
	if alternatives[answerIndex] == questionToAnswer.Answer {
		userAnsweredCorrectly = true
	}

	// Remove all questions having the same level that the one the user just answered.
	questions = removeQuestionsOfLevelFromSlice(questions, level)

	// Return flag indicating if the user answered correctly or not, and the slice of remaining questions.
	return userAnsweredCorrectly, questions
}

func findAllQuestionsOfLevel(questions []question, level int) (questionsOfSpecifiedLevel []question) {
	questionsOfSpecifiedLevel = make([]question, 0)
	for _, question := range questions {
		if question.Level == level {
			questionsOfSpecifiedLevel = append(questionsOfSpecifiedLevel, question)
		}
	}
	return
}

func removeQuestionsOfLevelFromSlice(questions []question, levelToRemove int) (remainingQuestions []question){
	indicesToRemove  := make([]int, 0)
	for index, question := range questions {
		if question.Level == levelToRemove{
			indicesToRemove = append(indicesToRemove, index)
		}
	}

	for i := len(indicesToRemove) - 1; i >= 0; i-- {
		questions = removeQuestionFromSlice(questions, indicesToRemove[i])
	}

	return questions
}

func removeQuestionFromSlice(questions []question, indexOfQuestionToRemove int) (remainingQuestions []question) {
	if len(questions) > 1 {
		copy(questions[indexOfQuestionToRemove:], questions[indexOfQuestionToRemove+1:])
		questions = questions[0 : len(questions)-1]
	} else {
		questions = nil // To indicate that no questions remain.
	}

	return questions
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

func getIntFromUser(msg string) (inputtedInteger int) {
	done := false
	reader := bufio.NewReader(os.Stdin)
	for !done {
		fmt.Print(msg)
		input, err := reader.ReadString('\n')
		if err == nil {
			input = strings.TrimSpace(input)
			i, err := strconv.Atoi(input)
			if err == nil {
				inputtedInteger = i
				done = true
			}
		}
	}

	return inputtedInteger
}

func getNonEmptyInputFromUser(msg string) (input string) {
	done := false
	reader := bufio.NewReader(os.Stdin)
	for !done {
		fmt.Print(msg)
		s, err := reader.ReadString('\n')
		if err == nil || err == io.EOF {
			s = strings.TrimSpace(s)
			if len(s) > 0 {
				input = s
				done = true
			}
		}
	}
	return
}

// Get input from user. Only allow responses in the range [1, maxChoice]
func getChoiceFromUser(msg string, maxChoice int) (retString string) {
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
