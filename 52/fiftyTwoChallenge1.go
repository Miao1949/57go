package main

import (
	"fmt"
	"os"
	"net"
	"io"
	"math/rand"
	"time"
)
// // Mon Jan 2 15:04:05 MST 2006
const Layout = "2006-01-02 15:04:05"
const Address = "localhost:8080"

var Quotes = []string{"Science is the belief in the ignorance of experts.", "The thing that doesn't fit is the thing that is most interesting.", "The first principle is that you must not fool yourself and you are the easiest person to fool", "Study hard what interests you the most in the most undisciplined, irreverent and original manner possible"}

func main()  {

	listener, err := net.Listen("tcp", Address)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could start listening to %s! Err: %v", Address, err)
		return
	}

	randomNumberGenerator := rand.New(rand.NewSource(time.Now().Unix()))

	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not handle accept! Err: %v", err)
		}
		go handleConnection(connection, randomNumberGenerator)
	}
}

func handleConnection(connection net.Conn, randomNumberGenerator *rand.Rand) {
	defer connection.Close()
	quote := Quotes[randomNumberGenerator.Intn(len(Quotes))]

	_, err := io.WriteString(connection, quote)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Got error when writing to connection! Err: %v", err)
	}

}
