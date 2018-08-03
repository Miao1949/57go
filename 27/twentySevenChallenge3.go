package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"unicode"
	"regexp"
)

var nameRegExp = regexp.MustCompile(`[a-zåäöA-ZÅÄÖ][a-zåäöA-ZÅÄÖ]+`)
var employeeIdRegRxp = regexp.MustCompile(`[a-zåäöA-ZÅÄÖ]{2}-\d\d\d\d$`)
var zipCodeRegRxp = regexp.MustCompile(`^\d+$`)


//IIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIII
// VALIDATOR
//IIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIII

type Validator interface {
	validate(value string) (valid bool, validationErrorString string)
}

//---------------------------------------------------------
// FIRST NAME VALIDATOR.
//---------------------------------------------------------
type FirstNameValidator struct {

}

func (fn FirstNameValidator) validate(value string) (valid bool, validationErrorString string) {
	valid, validationErrorString = validateName(value, "first name")
	return
}

//---------------------------------------------------------
// LAST NAME VALIDATOR.
//---------------------------------------------------------
type LastNameValidator struct {

}

func (fn LastNameValidator) validate(lastName string) (valid bool, validationErrorString string){
	valid, validationErrorString = validateName(lastName, "last name")
	return
}

//---------------------------------------------------------
// ZIP CODE VALIDATOR.
//---------------------------------------------------------
type ZipCodeValidator struct {

}

func (fn ZipCodeValidator) validate(zipCode string) (valid bool, validationErrorString string) {
	valid  = validateZipCode(zipCode)
	if !valid {
		validationErrorString = "The ZIP code must be numeric."
	}
	return
}

//---------------------------------------------------------
// EMPLOYEE ID VALIDATOR.
//---------------------------------------------------------
type EmployeeIdValidator struct {

}

func (fn EmployeeIdValidator) validate(employeeId string) (valid bool, validationErrorString string) {
	valid  = validateEmployeeId(employeeId)
	if !valid {
		validationErrorString = fmt.Sprintf("%s is not a valid ID.", employeeId)
	}
	return
}

//---------------------------------------------------------
// MAIN
//---------------------------------------------------------

func main() {
	firstNameValidator := FirstNameValidator{}
	lastNameValidator := LastNameValidator{}
	zipCodeValidator := ZipCodeValidator{}
	employeeIdValidator := EmployeeIdValidator{}

	handleField(firstNameValidator, "Enter the first name:", "first name")
	handleField(lastNameValidator, "Enter the last name:", "last name")
	handleField(zipCodeValidator, "Enter the ZIP code:", "zip code")
	handleField(employeeIdValidator, "Enter an employee ID:", "employee ID")

	fmt.Println("There were no errors found.")
}

//---------------------------------------------------------
// PRIVATE METHODS.
//---------------------------------------------------------

func handleField(validator Validator, inputString string, fieldName string) {
	done := false
	for !done {
		fmt.Print(inputString)
		inputtedValue, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		inputtedValue = strings.TrimSpace(inputtedValue)

		valid, validationErrorString := validator.validate(inputtedValue)
		if valid {
			done = true
		} else {
			fmt.Println(validationErrorString)
		}
	}
}

func validateName(name string, nameOfName string) (valid bool, validationErrorString string) {
	if isStringEmpty(name) {
		valid = false
		validationErrorString = fmt.Sprintf("The %s must be filled in.",nameOfName)
		return
	}

	if !nameRegExp.MatchString(name) {
		valid = false
		validationErrorString = fmt.Sprintf("%s is not a valid %s. It is too short.", name, nameOfName)
		return

	}

	return  true, ""
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

func validateZipCode(zipCode string) bool {
	return zipCodeRegRxp.MatchString(zipCode)
}


func validateEmployeeId(employeeId string) bool {
	return employeeIdRegRxp.MatchString(employeeId)
}

