package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"unicode"
)

func main() {
	fmt.Print("Enter the first name: ")
	firstName, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	firstName = strings.TrimSpace(firstName)

	fmt.Print("Enter the last name: ")
	lastName, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	lastName = strings.TrimSpace(lastName)

	fmt.Print("Enter the ZIP code: ")
	zipCode, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	zipCode = strings.TrimSpace(zipCode)

	fmt.Print("Enter the employeeId:")
	employeeId, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	employeeId = strings.TrimSpace(employeeId)

	valid, validationErrorString := validateInput(firstName, lastName, zipCode, employeeId)
	if valid {
		fmt.Println("There were no errors found.")
	} else {
		fmt.Println(validationErrorString)
	}
}

func validateName(name string, nameOfName string) (valid bool, validationErrorString string) {
	if isStringEmpty(name) {
		valid = false
		validationErrorString = fmt.Sprintf("The %s must be filled in.",nameOfName)
		return
	}

	if !validateNameIsSufficientlyLong(name) {
		valid = false
		validationErrorString = fmt.Sprintf("%s is not a valid %s. It is too short.", name, nameOfName)
		return
	}

	return  true, ""
}

func validateInput(firstName string, lastName string, zipCode string, employeeId string) (valid bool, validationErrorString string) {
	if ok, ves := validateName(firstName, "first name"); !ok {
		valid = false
		validationErrorString = ves
		return
	}

	if ok, ves := validateName(lastName, "last name"); !ok {
		valid = false
		validationErrorString = ves
		return
	}

	if !containsOnlyDigits(zipCode) {
		valid = false
		validationErrorString = "The ZIP code must be numeric."
		return
	}

	if !employeeIdValid(employeeId) {
		valid = false
		validationErrorString = fmt.Sprintf("%s is not a valid ID", employeeId)
		return
	}

	return true, ""
}

func isStringEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func validateNameIsSufficientlyLong(name string) bool {
	return len(strings.TrimSpace(name)) >= 2
}


func containsOnlyDigits(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func employeeIdValid(employeeId string) bool {
	// AA-1234.
	if len(employeeId) != 7 {
		return false
	}

	ss := []rune(employeeId)
	if unicode.IsLetter(ss[0]) && unicode.IsLetter(ss[1]) &&
		ss[2] == '-' &&
		unicode.IsDigit(ss[3]) && unicode.IsDigit(ss[4]) && unicode.IsDigit(ss[5]) && unicode.IsDigit(ss[6]) {
		return true
	}

	return false
}