package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"io"
)

const FilenameIn = "employees.txt"

func main() {
	employees, err := readEmployeesFromFile()
	if err == true {
		fmt.Println("Could not read file")
		os.Exit(1)
	}

	printEmployees(employees)
	employeeToRemove := getStringInput("Enter an employee name to remove: ")
	if !checkIfElementIsInArray(employees, employeeToRemove) {
		fmt.Printf("%s is not an employee\n", employeeToRemove)
	} else {
		employees = removeFromArray(employees, employeeToRemove)
		printEmployees(employees)
	}
}

func readEmployeesFromFile() (employees []string, errorOccured bool) {

	file, err := os.Open(FilenameIn)
	if err != nil {
		fmt.Println("Could not open file!")
		errorOccured = true
		return
	}
	defer file.Close()
	fileReader := bufio.NewReader(file)
	done := false
	for !done {
		employee, err := fileReader.ReadString('\n')
		if err == io.EOF {
			// Handle the case where there isn't a new line after the last employee.
			if len(strings.TrimSpace(employee)) > 0 {
				employees = append(employees, strings.TrimSpace(employee))
			}
			done = true
		} else if err != nil {
			fmt.Println("Error occurred!", err)
			errorOccured = true
			return
		} else {
			employees = append(employees, strings.TrimSpace(employee))
		}
	}

	return
}

func printEmployees(employees []string) {
	fmt.Printf("There are %d employees\n", len(employees))
	for _, employee := range employees {
		fmt.Println(employee)
	}
}

func getStringInput(msg string) (input string) {
	done := false
	for !done {
		fmt.Print(msg)
		s, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err == nil {
			input = strings.TrimSpace(s)
			done = true
		}
	}

	return
}

func checkIfElementIsInArray(arr []string, elementToCheck string) bool {
	for _, element := range arr {
		if element == elementToCheck {
			return true
		}
	}

	return false
}

func removeFromArray(arr []string, elementToRemove string) (arrAfterRemoval []string) {
	arrAfterRemoval = make([]string, 0)

	for _, element := range arr {
		if element != elementToRemove {
			arrAfterRemoval = append(arrAfterRemoval, element)
		}
	}

	return
}