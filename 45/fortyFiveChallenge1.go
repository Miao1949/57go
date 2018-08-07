package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"strings"
	"bufio"
)
const Inputfilename = "inputFile.txt"

func main() {
	fileContents, ok := readFile()
	if ok {
		done := false
		numberOfModifications := 0
		var modifiedStringBefore string
		modifiedStringBefore = fileContents
		for !done {
			modifiedStringAfter := strings.Replace(modifiedStringBefore, "utilize", "use", 1)
			if modifiedStringBefore == modifiedStringAfter {
				done = true
			} else {
				modifiedStringBefore = modifiedStringAfter
				numberOfModifications++
			}
		}
		writeToFile(modifiedStringBefore)
		fmt.Printf("Number of modifications: %d\n", numberOfModifications)
	}
}

func writeToFile(s string) {
	outfilename := getNonEmptyInput("Enter name of file to write to: ")
	file, e := os.OpenFile(outfilename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if e != nil {
		fmt.Fprintf(os.Stderr, "Could not open file %s\n", outfilename)
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err := writer.WriteString(s)
	if e != nil {
		fmt.Fprintf(os.Stderr, "Could not write to file %v\n", err)
		return

	}
	writer.Flush()
}

func readFile() (fileContents string, ok bool){
	data, e := ioutil.ReadFile(Inputfilename)
	if e != nil {
		fmt.Fprintf(os.Stderr, "Could not open file %s!\n", Inputfilename)
		return
	}

	fileContents = string(data)
	ok = true
	return
}

func getNonEmptyInput(msg string) (input string) {
	done := false
	for ; !done; {
		fmt.Print(msg)
		input, _ = bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.TrimSpace(input)
		if len(input) > 0 {
			done = true
		}
	}

	return input
}
