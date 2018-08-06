package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"strings"
	"bufio"
	"time"
)

const Infilename = "manyPersons.txt"
const Outfilename = "manyPersonsSorted.txt"

type Person struct {
	firstName string
	lastName string
}
func main() {
	start := time.Now()
	writePersonsToFile(sort(readDataFromFile()))
	secs := time.Since(start).Seconds()
	fmt.Printf("The execution took %.2f seconds.\n", secs)
}

func writePersonsToFile(persons []Person) {
	file, e := os.OpenFile(Outfilename, os.O_CREATE|os.O_WRONLY, 0644)
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

func readDataFromFile() (persons []Person) {
	fileData, err := ioutil.ReadFile(Infilename)
	if err != nil {
		fmt.Fprint(os.Stderr, "Could not open file. Error: %v\n", err)
		return
	}

	for _, line := range strings.Split(string(fileData), "\n") {
		fields := strings.Split(line, ",")
		if len(fields) >= 2{
			persons = append(persons, Person{strings.TrimSpace(fields[1]), strings.TrimSpace(fields[0])})
		}
	}

	return
}