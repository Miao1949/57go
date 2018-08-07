package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func main() {
	createDirectoriesAndFiles(getInput())
}

func createDirectoriesAndFiles(siteName string, siteAuthor string, createJavaScriptDir bool, createCssDir bool) {
	if createDirectories(siteName, createJavaScriptDir, createCssDir) {
		createIndexHtmlFile(siteName, siteAuthor)
	}
}

func createIndexHtmlFile(siteName string, siteAuthor string) {
	filename := siteName + "/index.html"
	file, e := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if e != nil {
		fmt.Fprint(os.Stderr, "Could not open %s\n", file)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	fmt.Fprintf(writer, "<html>\n")
	fmt.Fprintf(writer, "<head>\n")
	fmt.Fprintf(writer, "<meta>%s</meta>\n", siteAuthor)
	fmt.Fprintf(writer, "<title>%s</title>\n", siteName)
	fmt.Fprintf(writer, "<head>\n")
	fmt.Fprintf(writer, "</head>\n")
	fmt.Fprintf(writer, "</html>\n")
	writer.Flush()
}

func createDirectories(siteName string, createJavaScriptDir bool, createCssDir bool) (ok bool) {
	e := os.MkdirAll(siteName, os.ModePerm)
	if e != nil {
		fmt.Fprintf(os.Stderr, "Could not create dir %s", siteName)
		return  false
	}

	if createJavaScriptDir {
		dirName := siteName + "/js"
		e := os.MkdirAll(dirName, os.ModePerm)
		if e != nil {
			fmt.Fprintf(os.Stderr, "Could not create dir %s", dirName)
			return  false
		}
	}

	if createCssDir{
		dirName := siteName + "/css"
		e := os.MkdirAll(dirName, os.ModePerm)
		if e != nil {
			fmt.Fprintf(os.Stderr, "Could not create dir %s", dirName)
			return  false
		}
	}

	return true
}


func getInput() (siteName string, siteAuthor string, createJavaScriptDir bool, createCssDir bool) {
	siteName = getStringFromStdIn("Site name: ")
	siteAuthor = getStringFromStdIn("Site author: ")
	createJavaScriptDir = getYesOrNoInput("Do you want a folder for JavaScript? ")
	createCssDir = getYesOrNoInput("Do you want a folder for CSS? ")
	return

}

func getStringFromStdIn(msg string) (inputtedString string) {
	scanner := bufio.NewScanner(os.Stdin)
	done := false
	for !done {
		fmt.Print(msg)
		if scanner.Scan() {
			inputtedString = strings.TrimSpace(scanner.Text())

			if len(inputtedString) > 0 {
				done = true
			}
		}
	}

	return
}

func getRestrictedStringInput(msg string, allowedResponses map[string]bool) (retString string) {
	done := false
	for ; !done; {
		fmt.Print(msg)
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.TrimSpace(input)

		if len(input) > 0 && allowedResponses[input]{
			retString = input
			done = true
		}
	}
	return
}

func getYesOrNoInput(msg string) (answer bool) {
	yesOrNoOptions :=  make(map[string]bool)
	yesOrNoOptions["y"] = true
	yesOrNoOptions["Y"] = true
	yesOrNoOptions["n"] = true
	yesOrNoOptions["N"] = true
	retString := getRestrictedStringInput(msg, yesOrNoOptions)
	answer = retString == "y" || retString == "Y"
	return
}
