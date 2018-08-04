package main

import (
	"strings"
	"os"
	"fmt"
	"bufio"
)

const Filename = "contestants.txt"

func main ()  {
	writeContestantsToFile(getNamesOfContestants())
}

func getNamesOfContestants() (names []string) {
	done := false
	names = make([]string, 0)
	for !done {
		name := strings.TrimSpace(getStringUntilEmpty("Enter a name: "))
		if len(name) > 0 {
			names = append(names, name)
		} else {
			done = true
		}
	}

	return
}

func getStringUntilEmpty(msg string) (input string) {
	fmt.Print(msg)
	input, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	return
}

func writeContestantsToFile(contestants []string) (errorOccured bool){
	file, err := os.Create(Filename)
	if err != nil {
		fmt.Println("Could not open file to store data in!")
		errorOccured = true
		return
	}
	defer file.Close()
	fileWriter := bufio.NewWriter(file)
	defer fileWriter.Flush()

	for _, contestant := range contestants {
		_, err := fileWriter.WriteString(contestant + "\n")

		if err != nil {
			fmt.Println("Could not write to file! Error: ", err)
			errorOccured = true
			return
		}
	}

	return
}
