package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"time"
	"io"
)

const (
	FirstName      = "firstName"
	LastName       = "lastName"
	Position       = "Position"
	SeparationDate = "Separation Date"
	Name           = "Name"
	Filename       = "persons.txt"
)

func main() {
	choise := getStringInput("Filter on name(n) or position(p) or separation date(s)? ", []string{"n", "p", "s"})
	var filteredData []map[string]string

	data := getData()
	switch choise {
	case "n":
		filteredData = filterOnSearchString(data, []string{FirstName, LastName})
	case "p":
		filteredData = filterOnSearchString(data, []string{Position})
	case "s":
		filteredData = filterDataBasedOnSeparationDate(data)
	}
	printData(filteredData)
}

func filterOnSearchString(data []map[string]string, attributesToFilterOn []string) (filteredData []map[string]string){
	fmt.Print("Enter a search string: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		searchString := scanner.Text()
		filteredData = filterData(data, attributesToFilterOn, searchString)
	}
	return
}

func filterDataBasedOnSeparationDate(data []map[string]string) (filteredData []map[string]string) {
	filteredData = make([]map[string]string, 0)
	// Mon Jan 2 15:04:05 MST 2006
	layout := "2006-01-02"
	now := time.Now()
	sixMonthsAgo := now.AddDate(-2, -6, 0)
	fmt.Println(sixMonthsAgo)
	for _, personMap := range data {
		separationDateAsString := personMap[SeparationDate]

		if len(separationDateAsString) > 0 {
			t, _ := time.Parse(layout, separationDateAsString)
			fmt.Println(t)
			if t.Before(sixMonthsAgo) {
				filteredData = append(filteredData, personMap)
			}
		}
	}

	return
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

func getData() (data []map[string]string) {
	file, err := os.Open(Filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	done := false
	var personMap map[string]string
	for !done {
		line, e := reader.ReadString('\n')

		if e != nil {
			done = true
		}

		if e == nil || e == io.EOF {
			line = strings.TrimSpace(line)
			fields := strings.Split(line, ",")
			if len(fields) > 3 {
				personMap = createMap(strings.TrimSpace(fields[0]), strings.TrimSpace(fields[1]), strings.TrimSpace(fields[2]), strings.TrimSpace(fields[3]))
				data = append(data, personMap)
			} else if len(fields) > 2 {
				personMap = createMap(strings.TrimSpace(fields[0]), strings.TrimSpace(fields[1]), strings.TrimSpace(fields[2]), "")
				data = append(data, personMap)
			}
		}
	}

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

