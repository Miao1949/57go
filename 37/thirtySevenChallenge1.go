package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
	"math/rand"
	"time"
)

var Digits = []rune("0123456789")
var Characters = []rune("qwertyuiopåasdfghjklläzxcvbnmQWERTYUIOPÅASDFGHJKLÖÄZXCVBNM")
var SpecialCharacters = []rune("'!#¤%&/()+@£$€¥{[]}")

func main() {
	minLen, numberOfSpecialCharacters, numberOfDigits := getInput()

	if minLen < numberOfSpecialCharacters + numberOfDigits {
		fmt.Println("Min len must be >= numberOfSpecialCharacters + numberOfDigits ")
		os.Exit(1)
	}

	randomNumberGenerator := rand.New(rand.NewSource(time.Now().Unix()))

	var base = make([]rune, 0)
	for i := 0; i < numberOfDigits; i++ {
		base = append(base, Digits[randomNumberGenerator.Intn(len(Digits))])
	}

	for i := 0; i < numberOfSpecialCharacters; i++ {
		base = append(base, SpecialCharacters[randomNumberGenerator.Intn(len(SpecialCharacters))])
	}

	for i := 0; i < minLen - numberOfSpecialCharacters -numberOfDigits; i++ {
		base = append(base, possiblyTranslateLetter(Characters[randomNumberGenerator.Intn(len(Characters))], randomNumberGenerator))
	}

	var randomizedRuneArray []rune
	var r rune
	for len(base) > 0 {
		r, base = pop(base, randomNumberGenerator)
		randomizedRuneArray = append(randomizedRuneArray, r)
	}

	fmt.Println("Your password is")
	fmt.Println(string(randomizedRuneArray))
}

func possiblyTranslateLetter(r rune, randomNumberGenerator *rand.Rand) (possiblyTranslatedRune rune) {
	possiblyTranslatedRune = r
	if randomNumberGenerator.Intn(2) % 2 == 0 {
		if r == 'A' {
			//fmt.Println("Translating A to 4")
			possiblyTranslatedRune = '4'
		} else if r == 'E' {
			//fmt.Println("Translating E to 3")
			possiblyTranslatedRune = '3'
		} else if r == 'o' {
			//fmt.Println("Translating O to 0")
			possiblyTranslatedRune = '0'
		}
	}
	return
}

func pop(runes []rune, randomNumberGenerator *rand.Rand) (selectedRune rune, runesArrayAfterPopping []rune) {
	runesArrayAfterPopping = make([]rune, 0)
	selectedIndex := randomNumberGenerator.Intn(len(runes))
	for index, rune := range runes {
		if index == selectedIndex {
			selectedRune = rune
		} else {
			runesArrayAfterPopping = append(runesArrayAfterPopping, rune)
		}
	}
	return
}

func getInput() (minLen int, numberOfSpecialCharacters int, numberOfDigits int) {
	minLen = getPositiveNumber("What's the minimum length? ")
	numberOfDigits = getPositiveNumber("How many numbers? " )
	numberOfSpecialCharacters = getPositiveNumber("How many special characters? ")
	return
}

func getPositiveNumber(msg string) (number int) {
	done := false
	reader := bufio.NewReader(os.Stdin)
	for !done {
		fmt.Print(msg)
		s, e := reader.ReadString('\n')
		if e == nil {
			s = strings.TrimSpace(s)
			n, e := strconv.Atoi(s)
			if e == nil {
				if n > 0 {
					number = n
					done = true
				}
			}
		}
	}
	return
}
