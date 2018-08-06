package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	FirstName      = "firstName"
	LastName       = "lastName"
	Position       = "Position"
	SeparationDate = "Separation Date"
	Name           = "Name"
)

func main() {
	unsertedData := getDataFromDatabase()
	choise := getStringInput("Sort on separation date(s), position(p) or last name(l)? ", []string{"s", "p", "l"})
	var sortedData []map[string]string
	switch choise {
	case "s":
		sortedData = sortData(unsertedData, SeparationDate)
	case "p":
		sortedData = sortData(unsertedData, Position)
	case "l":
		sortedData = sortData(unsertedData, LastName)
	}

	printData(sortedData)
}

func getStringInput(msg string, allowedResponses []string) (retString string) {
	allowedResponsesMap := make(map[string]bool)
	for _, allowedResponse := range allowedResponses {
		allowedResponsesMap[allowedResponse] = true
	}

	return getRestrictedStringInput(msg, allowedResponsesMap)
}

func getRestrictedStringInput(msg string, allowedResponses map[string]bool) (retString string) {
	done := false
	for ; !done; {
		fmt.Print(msg)
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.TrimSpace(input)

		if len(input) > 0 && allowedResponses[input] {
			retString = input
			done = true
		}
	}
	return
}

// Sort using insertion sort.
func sortData(data []map[string]string, attributeToSortOn string) (sortedData []map[string]string) {
	sortedData = make([]map[string]string, len(data))

	for indexOfPersonInOriginalSlice, personMapToInsert := range data {
		insertionIndex := -1
		for index, personMapAlreadyInserted := range sortedData {
			if personMapAlreadyInserted == nil {
				// We've reached the end of the already inserted elements.
				break
			} else if personMapToInsert[attributeToSortOn] < personMapAlreadyInserted[attributeToSortOn] {
				insertionIndex = index
				break
			}
		}

		if insertionIndex < 0 {
			// Easy, just insert last.
			sortedData[indexOfPersonInOriginalSlice] = personMapToInsert
		} else {
			// Insert in the first or in the middle.
			// Move the elements that should lie after one step backwatds (discard the last element, which will be empty as the sortedData slice is correctly sized from the beginning.
			copy(sortedData[insertionIndex+1:], sortedData[insertionIndex:len(sortedData)-1])
			// Then just set the element at the insertion point to the element we're currently sorting.
			sortedData[insertionIndex] = personMapToInsert
		}
	}

	return
}

func printData(data []map[string]string) {
	lengthOfLongestName := findLengthOfLongestName(data)
	lengthOfLongestPosition := findLengthOfLongestAttributeValue(Position, data)
	lengthOfLongestSeparationDate := findLengthOfLongestAttributeValue(SeparationDate, data)

	nameColumnLen := max(lengthOfLongestName, len(Name)) + 1
	positionColumnLen := max(lengthOfLongestPosition, len(Position)) + 2
	separationColumnLen := max(lengthOfLongestSeparationDate, len(SeparationDate)) + 1

	printHeader(nameColumnLen, positionColumnLen, separationColumnLen)

	for _, personMap := range data {
		name := personMap[FirstName] + " " + personMap[LastName]
		position := personMap[Position]
		separationDate := personMap[SeparationDate]

		fmt.Print(name)
		fmt.Print(strings.Repeat(" ", nameColumnLen-len(name)-1))
		fmt.Print(" | ")

		fmt.Print(position)
		fmt.Print(strings.Repeat(" ", positionColumnLen-len(position)-2))
		fmt.Print(" | ")

		fmt.Println(separationDate)
	}
}

func printHeader(nameColumnLen int, positionColumnLen int, separationColumnLen int) {
	fmt.Println(Name, strings.Repeat(" ", nameColumnLen-len(Name)-1)+"| "+Position+strings.Repeat(" ", positionColumnLen-len(Position)-1)+"| "+SeparationDate)
	fmt.Print(strings.Repeat("-", nameColumnLen))
	fmt.Print("|")
	fmt.Print(strings.Repeat("-", positionColumnLen))
	fmt.Print("|")
	fmt.Print(strings.Repeat("-", separationColumnLen))
	fmt.Println("")
}

func max(i1 int, i2 int) int {
	if i1 > i2 {
		return i1
	}

	return i2
}

func findLengthOfLongestName(data []map[string]string) (lengthOfLongestName int) {
	lengthOfLongestName = 0

	for _, personMap := range data {
		name := personMap[FirstName] + " " + personMap[LastName]
		if len(name) > lengthOfLongestName {
			lengthOfLongestName = len(name)
		}
	}
	return
}

func findLengthOfLongestAttributeValue(attribute string, data []map[string]string) (lengthOfLongestAttributeValue int) {
	lengthOfLongestAttributeValue = 0

	for _, personMap := range data {
		attributeValue := personMap[attribute]
		if len(attributeValue) > lengthOfLongestAttributeValue {
			lengthOfLongestAttributeValue = len(attributeValue)
		}
	}
	return
}

type Employee struct { // NOTE: THE FIELDS MUST START WITH A CAPITAL LETTER! Otherwise the mapping won't work.
	ID bson.ObjectId `bson:"_id" json:"id"`
	FirstName string        `bson:"firstName" json:"firstName"`
	LastName  string        `bson:"lastName" json:"lastName"`
	Position string        `bson:"position" json:"position"`
	SeparationDate string        `bson:"separationDate" json:"separationDate"`
}

func getDataFromDatabase() (data []map[string]string) {
	// Open session to DB.
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Fetch collection.
	employeeCollection := session.DB("exercisesForProgrammers").C("employee")

	// Fetch data from DB.
	var employees []*Employee
	err = employeeCollection.Find(bson.M{}).All(&employees)
	if err != nil {
		panic(err)
	}

	// Copy to desired datatype.
	data = make([]map[string]string, 0)
	for _, employee := range employees {
		m := make(map[string]string)
		m[FirstName] = employee.FirstName
		m[LastName] = employee.LastName
		m[Position] = employee.Position
		m[SeparationDate] = employee.SeparationDate
		data = append(data, m)
	}

	return
}