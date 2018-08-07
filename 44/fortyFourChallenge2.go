package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"encoding/json"
	"bufio"
	"strings"
	"strconv"
)

const Filename = "Products.json"
const Outfile = "outfile.json"

const Products = "Products"
const Name = "Name"
const Price = "Price"
const Quantity = "Quantity"

type ProductsType struct {
	Products []Product
}

type Product struct {
	Name     string
	Price    float64
	Quantity float64
}

func main() {

	productsInInternalFormat := readJsonFile()
	askUserForNameOfProductToFindAndDisplayThatProduct(productsInInternalFormat.Products)
}

func convertToInternalFormat(productsInJsonFormat []interface{}) (productsInInternalFormat []Product) {
	productsInInternalFormat = make([]Product, len(productsInJsonFormat))
	for index, p := range productsInJsonFormat {
		product := p.(map[string]interface{})
		productsInInternalFormat[index] = Product{product[Name].(string), product[Price].(float64), product[Quantity].(float64)}
	}

	return
}

func askUserForNameOfProductToFindAndDisplayThatProduct(products []Product) {
	done := false
	for !done {
		productName := getNonEmptyInputFromUser("What is the product Name? ")
		product, found := findProduct(productName, products)
		if found {
			printProduct(product)
		} else {
			fmt.Println("Sorry, that product was not found in our inventory.")
			if getYesOrNoInput("Would you like to add the product to the inventory? ") {
				productToAdd := createProduct(productName)
				products = append(products, productToAdd)
				writeToJsonFile(products)
			}
		}

		if !getYesOrNoInput("Would you like to perform another search? ") {
			done = true
		}
	}
}

func createProduct(productName string) (createdProduct Product) {
	price := getNonNegativeFloatInput("Please enter a Price: ")
	quantity := getNonNegativeFloatInput("Please enter a Quantity: ")
	createdProduct = Product{productName, price, quantity}
	return
}

func findProduct(productName string, products []Product) (product Product, found bool){
	for _, p := range products {
		if strings.ToUpper(p.Name) == strings.ToUpper(productName) {
			product =  p
			found = true
			return
		}
	}
	return
}

func printProduct(product Product) {
	fmt.Printf("Name: %s\n", product.Name)
	fmt.Printf("Price: $%.2f\n", product.Price)
	fmt.Printf("Quantity: %.0f\n", product.Quantity)
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

func readJsonFile() (products ProductsType){
	fileContent, e := ioutil.ReadFile(Filename)
	if e != nil {
		fmt.Fprint(os.Stderr, "Could not read file %s\n", Filename)
		panic(e)
	}

	if err := json.Unmarshal(fileContent, &products); err != nil {
		fmt.Println("Could not urmarshal fileContent as json!")
		return
	}
	return
}

func writeToJsonFile(products []Product) {
	productsStruct := ProductsType{products}
	marshalledJsonData, e := json.Marshal(productsStruct)
	if e != nil {
		fmt.Fprintf(os.Stderr, "Could not marshall products to json!\n")
		return
	}

	file, e := os.OpenFile(Outfile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if e != nil {
		fmt.Fprintf(os.Stderr, "Could not open file %s\n", Outfile)
		return

	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err := writer.Write(marshalledJsonData)
	if e != nil {
		fmt.Fprintf(os.Stderr, "Could not write to file %v\n", err)
		return

	}
	writer.Flush()
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

// Get yes or no answer:
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

func getNonNegativeFloatInput(msg string) (retFloat float64) {
	done := false
	for ; !done; {
		fmt.Print(msg)
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.TrimSpace(input)
		i, err := strconv.ParseFloat(input, 64)

		if err == nil && i > 0 {
			retFloat= i
			done = true
		}
	}
	return
}

