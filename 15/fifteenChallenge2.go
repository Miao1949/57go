package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"hash/crc32"
)

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

func getSha1HashOfPassword(password string) (hashOfPassword uint32) {
	hashAlg := crc32.NewIEEE()
	hashAlg.Write([]byte(password))
	hashOfPassword = hashAlg.Sum32()
	return
}

func main() {
	// Set up map.
	usernameToPasswordMap := make(map[string]uint32)
	const larsPassword = "pappaärbäst"
	usernameToPasswordMap["lars"] = getSha1HashOfPassword(larsPassword)

	username, password := getInput()


	if usernameToPasswordMap[username] == getSha1HashOfPassword(password) {
		fmt.Println("Welcome!")
	} else {
		fmt.Println("I don't know you.")
	}
}
