package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	employees := []string{"John Smith", "Jackie Jackson", "Chris Jones", "Amanda Cullen", "Jeremy Goodwin"}

	printEmployees(employees)
	employeeToRemove := getStringInput("Enter an employee name to remove: ")
	employees = removeFromArray(employees, employeeToRemove)
	printEmployees(employees)
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

func removeFromArray(arr []string, elementToRemove string) (arrAfterRemoval []string) {
	arrAfterRemoval = make([]string, 0)

	for _, element := range arr {
		if element != elementToRemove {
			arrAfterRemoval = append(arrAfterRemoval, element)
		}
	}

	return
}