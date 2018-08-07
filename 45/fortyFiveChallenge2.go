package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"strings"
	"bufio"
	"io"
)
const InputFilename = "inputFile.txt"
const ReplacementFilename = "replacementWords.csv"

type Replacement struct {
	WordToReplace string
	ReplacementWord string
}

func main() {
	replacements, ok := readReplacementsFromFile()
	if !ok {
		return
	}

	fileContents, ok := readFile()
	if !ok {
		return
	}

	done := false
	numberOfModifications := 0
	//numberOfReplacements := 0
	var modifiedStringBefore string
	modifiedStringBefore = fileContents
	for !done {
		//modifiedStringAfter := strings.Replace(modifiedStringBefore, "utilize", "use", 1)
		modifiedStringAfter, numberOfReplacements := replace(modifiedStringBefore, replacements)

		if modifiedStringBefore == modifiedStringAfter {
			done = true
		} else {
			modifiedStringBefore = modifiedStringAfter
			numberOfModifications += numberOfReplacements
		}
	}
	writeToFile(modifiedStringBefore)
	fmt.Printf("Number of modifications: %d\n", numberOfModifications)

}

func replace(s string, replacements []Replacement) (outstring string, numberOfReplacements int) {
	stringUnderProcessing := s
	for _, replacement := range replacements {
		stringAfter := strings.Replace(stringUnderProcessing, replacement.WordToReplace, replacement.ReplacementWord, 1)
		if stringAfter != stringUnderProcessing {
			numberOfReplacements++
		}
		stringUnderProcessing = stringAfter
	}

	outstring = stringUnderProcessing
	return
}

func readReplacementsFromFile() (replacements []Replacement, ok bool) {
	file, e := os.Open(ReplacementFilename)
	if e != nil {
		fmt.Fprintf(os.Stderr, "Could not open replacement file!\n")
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	done := false
	for !done {
		line, e := reader.ReadString('\n')
		if e != nil {
			done = true
		}

		if e == nil || e == io.EOF {
			fields := strings.Split(line, ",")
			if len(fields) >= 2 {
				replacements = append(replacements, Replacement{strings.TrimSpace(fields[0]), strings.TrimSpace(fields[1])})
			}
		}
	}

	ok = true
	return
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
	data, e := ioutil.ReadFile(InputFilename)
	if e != nil {
		fmt.Fprintf(os.Stderr, "Could not open file %s!\n", InputFilename)
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
