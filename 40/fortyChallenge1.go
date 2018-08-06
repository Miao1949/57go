package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

const (
	FirstName      = "firstName"
	LastName       = "lastName"
	Position       = "Position"
	SeparationDate = "Separation Date"
	Name           = "Name"
)

func main() {
	fmt.Print("Enter a search string: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		searchString := scanner.Text()
		printData(filterData(createData(), []string{FirstName, LastName}, searchString))
	}

}

func filterData(data []map[string]string, attributesToFilterOn []string, stringToFilterOn string) (filteredData []map[string]string) {
	filteredData = make([]map[string]string, 0)
	for _, personMap := range data {
		for _, attributeToFilterOn := range attributesToFilterOn {
			if strings.Contains(strings.ToUpper(personMap[attributeToFilterOn]), strings.ToUpper(stringToFilterOn)) {
				filteredData = append(filteredData, personMap)
				break
			}
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

	fmt.Println("Results:")
	printHeader(nameColumnLen, positionColumnLen, separationColumnLen)

	for _, personMap := range data {
		name := personMap[FirstName] + " " + personMap[LastName]
		position := personMap[Position]
		separationDate := personMap[SeparationDate]

		fmt.Print(name)
		fmt.Print(strings.Repeat(" ", nameColumnLen-len(name) - 1))
		fmt.Print(" | ")

		fmt.Print(position)
		fmt.Print(strings.Repeat(" ", positionColumnLen-len(position) - 2))
		fmt.Print(" | ")

		fmt.Println(separationDate)
	}
}

func printHeader(nameColumnLen int, positionColumnLen int, separationColumnLen int) {
	fmt.Println(Name, strings.Repeat(" ", nameColumnLen-len(Name) - 1)+"| "+Position+strings.Repeat(" ", positionColumnLen-len(Position) - 1)+"| "+SeparationDate)
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

func createData() (data []map[string]string) {
	data = make([]map[string]string, 0)

	map1 := createMap("John", "Johnson", "Manager", "2016-12-31")
	map2 := createMap("Tou", "Xiong", "Software Engineer", "2016-10-05")
	map3 := createMap("Michaela", "Michaelson", "District Manager", "2015-12-19")
	map4 := createMap("Jake", "Jacobson", "Programmer", "")
	map5 := createMap("Jacquelyn", "Jackson", "DBA", "")
	map6 := createMap("Sally", "Weber", "Web Developer", "2015-12-18")

	data = []map[string]string{map1, map2, map3, map4, map5, map6}
	return
}

func createMap(firstName string, lastName string, position string, separationDate string) (m map[string]string) {
	m = make(map[string]string)
	m[FirstName] = firstName
	m[LastName] = lastName
	m[Position] = position
	m[SeparationDate] = separationDate
	return
}

