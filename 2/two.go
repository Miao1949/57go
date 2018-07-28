package main

import "fmt"

func main() {
	var input string
	fmt.Print("What is the input string? ")
	fmt.Scanf("%s", &input)
	fmt.Println(input, "has", len(input), "characters.")
}
