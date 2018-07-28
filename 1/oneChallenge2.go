package main

import "fmt"

func main() {
	fmt.Print("What is your name: ")
	var name string
	fmt.Scanf("%s", &name)

	if name == "Lars" {
		fmt.Println("Hello, ", name, " nice to meet you! Hey! That's my name too!")
	} else if name == "Frida" || name == "Anne-Frid" {
		fmt.Println("Hello, ", name, " nice to meet you! That's what my wife is called too!")
	} else {
		fmt.Println("Hello, ", name, " nice to meet you!")
	}
}
