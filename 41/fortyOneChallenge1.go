package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
)

const Outfilename = "sortedPersons.txt"

type Person struct {
	firstName string
	lastName string
}
func main() {
	persons := getDataFromUser()
	printPersons(persons)
	writePersonsToFile(sort(persons))
}

func writePersonsToFile(persons []Person) {
	file, e := os.OpenFile(Outfilename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if e != nil {
		fmt.Fprint(os.Stderr, "Could not open file. Error: %v\n", e)
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	fmt.Fprintf(writer,"Total of %d names\n", len(persons))
	fmt.Fprint(writer, strings.Repeat("-", 16) + "\n")
	for _, person := range persons {
		fmt.Fprintf(writer,"%s, %s\n", person.lastName, person.firstName)
	}

	writer.Flush()
}

func sort(persons []Person) (sortedPersons []Person) {
	sortedPersons = make([]Person, len(persons))

	for indexOfPersonToSort, person := range persons {
		insertionPoint := -1
		for index, sortedPerson := range sortedPersons {
			if person.lastName < sortedPerson.lastName {
				insertionPoint = index
			}
		}

		if insertionPoint < 0 {
			sortedPersons[indexOfPersonToSort] = person
		} else {
			copy(sortedPersons[insertionPoint + 1:], sortedPersons[insertionPoint:len(sortedPersons) - 1])
			sortedPersons[insertionPoint] = person
		}
	}

	return
}

func printPersons(persons []Person) {
	fmt.Printf("Total of %d names\n", len(persons))
	fmt.Println(strings.Repeat("-", 16))
	for _, person := range persons {
		fmt.Printf("%s, %s\n", person.lastName, person.firstName)
	}
}

func getDataFromUser() (persons []Person) {
	line, end := getStringInput("Please enter a name on the format 'LastName, FirstName'. End with a blank line: ");
	for !end {
		fields := strings.Split(line, ",")
		if len(fields) >= 2 {
			persons = append(persons, Person{strings.TrimSpace(fields[1]), strings.TrimSpace(fields[0])})
		}
		line, end = getStringInput("Please enter a name on the format 'LastName, FirstName'. End with a blank line: ");
	}
	return
}

// Get input from user. Indicate stop when the user has entered an empty input. Indicate that by setting a flag.
func getStringInput(msg string) (retVal string, end bool) {
	done := false
	end = false
	for !done {
		fmt.Print(msg)
		s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		retVal = strings.TrimSpace(s)

		if len(retVal) == 0 {
			end = true
			done = true
		} else {
			done = true
		}
	}
	return
}