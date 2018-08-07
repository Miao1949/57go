package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"os"
	"strings"
)

const urlToOpenNotify = "http://api.open-notify.org/astros.json"

type Person struct {
	Craft string
	Name string
}

type AstroInfo struct {
	Message string
	People []Person
	Number int
}

func main() {
	astroData, _:= readAstroData()
	displayAstroInfo(astroData)
}

func readAstroData() (astroInfo AstroInfo, errorToReturn error) {
	fmt.Printf("Fetching data from %s\n", urlToOpenNotify)
	resp, err := http.Get(urlToOpenNotify)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not read data from URL! Error: %v", err)
		return
	} else {
		// Make sure the connection is closed.
		defer resp.Body.Close()

		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr,"Could not read from response! Error: %v", err)
			errorToReturn = err
			return
		} else {
			if err := json.Unmarshal(contents, &astroInfo); err != nil {
				fmt.Fprintf(os.Stderr,"Could not urmarshal contents as json!")
				errorToReturn = err
				return
			}
		}
	}
	return
}

func displayAstroInfo(astroInfo AstroInfo)  {
	nameColumnLength := findLengthOfLongestName(astroInfo.People) + 1

	fmt.Printf("There are %d people in space right now:\n", astroInfo.Number)
	printHeader(nameColumnLength)
	for _, person := range astroInfo.People{
		fmt.Print(person.Name)
		fmt.Print(strings.Repeat(" ", nameColumnLength - len(person.Name) + 1))
		fmt.Print("| ")
		fmt.Println(person.Craft)
	}
}

func printHeader(nameColumnLength int) {
	fmt.Print("Name")
	fmt.Print(strings.Repeat(" ", nameColumnLength - len("Name") + 1))
	fmt.Println("| Craft")
	fmt.Print(strings.Repeat("-", nameColumnLength + 1))
	fmt.Print("|")
	fmt.Println(strings.Repeat("-", 6))

}

func findLengthOfLongestName(persons []Person) (lengthOfLongestName int){
	for _, person := range persons {
		if len(person.Name) > lengthOfLongestName {
			lengthOfLongestName = len(person.Name)
		}
	}
	return
}

