package main

import "fmt"

const QUOTE = "quote"
const AUTHOR = "author"

func main() {
	quotesAndAuthor := []map[string]string {
		{QUOTE : "These aren't the droids you're looking for.", AUTHOR: "Obi-Wan Kenobi",},
		{QUOTE : "Veni, vici, vidi.", AUTHOR: "Ceasar",},
		{QUOTE : "Alea jacta est.", AUTHOR: "Ceasar",},
		{QUOTE : "It is a bridge, and you know what we do to bridges, don't you?", AUTHOR: "Hedge",},
		{QUOTE : "Now get your hands off me, or I will loose my temper.", AUTHOR: "Mael",},
		{QUOTE : "The rage of an elder god unleashed.", AUTHOR: "Mael",},
	}

	for _, v := range quotesAndAuthor {
		output := v[AUTHOR] + " says " + "\"" + v[QUOTE]+ "\""
		fmt.Println(output)
	}
}
