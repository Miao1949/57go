package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"encoding/json"
	"bufio"
	"strings"
)

const Filename = "products.json"
const Products = "products"
const Name = "name"
const Price = "price"
const Quantity = "quantity"

func main() {
	jsonData := readJsonFile()
	products := jsonData[Products].([]interface{})
	askUserForNameOfProductToFindAndDisplayThatProduct(products)
}

func askUserForNameOfProductToFindAndDisplayThatProduct(products []interface{}) {
	done := false
	for !done {
		productName := getNonEmptyInputFromUser("What is the product name? ")
		product := findProduct(productName, products)
		if product != nil {
			printProduct(product)
			done = true
		} else {
			fmt.Println("Sorry, that product was not found in our inventory.")
		}

	}
}

func findProduct(productName string, products []interface{}) (product map[string]interface{}){
	for _, p := range products {
		pp := p.(map[string]interface{})
		if strings.ToUpper(pp[Name].(string)) == strings.ToUpper(productName) {
			product =  pp
			return
		}
	}
	return
}

func printProduct(product map[string]interface{}) {
	fmt.Printf("Name: %s\n", product[Name].(string))
	fmt.Printf("Price: $%.2f\n", product[Price].(float64))
	fmt.Printf("Quantity: %.0f\n", product[Quantity].(float64))
}

func getNonEmptyInputFromUser(msg string) (inputtedString string) {
	done := false
	scanner := bufio.NewScanner(os.Stdin)
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

func readJsonFile() (dat map[string]interface{}){
	fileContent, e := ioutil.ReadFile(Filename)
	if e != nil {
		fmt.Fprint(os.Stderr, "Could not read file %s\n", Filename)
		panic(e)
	}

	if err := json.Unmarshal(fileContent, &dat); err != nil {
		fmt.Println("Could not urmarshal fileContent as json!")
		return
	}
	return
}