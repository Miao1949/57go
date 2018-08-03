package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"unicode"
)

const WeakCharacterLimit = 8

const VeryWeak = "very weak"
const Weak = "weak"
const Average = "average"
const Strong = "strong"
const VeryStrong = "very strong"

const ClassificationIntVeryWeak = 1
const ClassificationIntWeak = 2
const ClassificationIntAverage = 3
const ClassificationIntStrong = 4
const ClassificationIntVeryStrong = 5

type Classification struct {
	classificationString string
	classificationInt int
}

var ClassificationVeryWeak = Classification{VeryWeak, ClassificationIntVeryWeak}
var ClassificationWeak = Classification{Weak, ClassificationIntWeak}
var ClassificationAverage = Classification{Average, ClassificationIntAverage}
var ClassificationStrong = Classification{Strong, ClassificationIntStrong}
var ClassificationVeryStrong = Classification{VeryStrong, ClassificationIntVeryStrong}



func getNonEmptyInput(msg string) (input string) {
	done := false
	for ; !done; {
		fmt.Print(msg)
		input, _ = bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.TrimSpace(input)
		if len(input) > 0 {
			done = true
		}
	}

	return input
}

func containsDigit(s string) bool {
	for _,c := range s {
		if unicode.IsDigit(c) {
			return true
		}
	}

	return false
}

func containsLetter(s string) bool {
	for _,c := range s {
		if unicode.IsLetter(c) {
			return true
		}
	}

	return false
}

func containsSpecialCharater(s string) bool {
	for _,c := range s {
		if unicode.IsSymbol(c) || unicode.IsPunct(c){
			return true
		}
	}

	return false
}

func main() {
	password := getNonEmptyInput("Enter password to check: ")
	classification := passwordValidator(password)
	fmt.Printf("The password '%s' is a %s password.\n", password, classification.classificationString)

}

func passwordValidator(password string) (classification Classification) {
	containsDigit := containsDigit(password)
	containsLetter := containsLetter(password)
	containsSpecialCharacter  := containsSpecialCharater(password)

	if len(password) < WeakCharacterLimit {
		if containsDigit && ! containsLetter && !containsSpecialCharacter {
			classification = ClassificationVeryWeak
		} else if !containsDigit && containsLetter && !containsSpecialCharacter {
			classification = ClassificationWeak
		} else {
			classification = ClassificationAverage
		}
	} else {
		if containsDigit && containsLetter && !containsSpecialCharacter {
			classification = ClassificationStrong
		} else if containsDigit && containsLetter && containsSpecialCharacter {
			classification = ClassificationVeryStrong
		} else {
			classification = ClassificationAverage
		}
	}

	return
}
