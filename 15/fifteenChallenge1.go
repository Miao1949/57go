package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

// Cannot do this right now as the connection to the Internet is really bad in Löre.

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

func getInput() (username string, password string) {
	username = getNonEmptyInput("What is the username? ")
	password = getNonEmptyInput("What is the password? ")
	return
}

func main() {
	username, password := getInput()

	if username == "lars" {
		if password == "pappaärbäst" {
			fmt.Println("Welcome!")
		} else {
			fmt.Println("I don't know you.")
		}

	} else {
		fmt.Println("I don't know you.")
	}
}
