package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"strings"
	"bufio"
)
const Inputfilename = "inputFile.txt"
const Outputfilename = "outputFile.txt"


func main() {
	fileContents, ok := readFile()
	if ok {
		modifiedString := strings.Replace(fileContents, "utilize", "use", -1)
		writeToFile(modifiedString)
	}
}

func writeToFile(s string) {
	file, e := os.OpenFile(Outputfilename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if e != nil {
		fmt.Fprintf(os.Stderr, "Could not open file %s\n", Outputfilename)
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