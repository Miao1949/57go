package main

import ("fmt";
	"net/http"
	"io/ioutil"
	"encoding/json"
	"os"
	"bufio"
	"strings"
	"strconv"
)

const openExchangeAppId = "5e432fbaf6544dc7b084e8190dfd79fd"
const urlToExchangeService = "https://openexchangerates.org/api/latest.json?app_id=" + openExchangeAppId

func readRatesFromOpenExchange() (ratesMap map[string]interface{}, errorToReturn error) {
	fmt.Printf("Fetching exchange rates from %s\n", urlToExchangeService)
	resp, err := http.Get(urlToExchangeService)
	if err != nil {
		fmt.Println("Could not read exchange rates from URL! Error: ", err)
		return
	} else {
		// Make sure the connection is closed.
		defer resp.Body.Close()

		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Could not read from response! Error: ", err)
			errorToReturn = err
			return
		} else {
			fmt.Println(string(contents))
			var dat map[string]interface{}
			if err := json.Unmarshal(contents, &dat); err != nil {
				fmt.Println("Could not urmarshal contents as json!")
				errorToReturn = err
				return
			}

			ratesMap = dat["rates"].(map[string]interface{})
		}
	}
	return
}

// Get integer input.
func getNonNegativeIntegerInput(msg string) (retInt int) {
	done := false
	for ; !done; {
		fmt.Print(msg)
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.TrimSpace(input)
		i,err := strconv.Atoi(input)

		if err == nil && i > 0 {
			retInt = i
			done = true
		}
	}
	return
}

// Get non-empty input.
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

func getInput() (amount int, country string) {
	amount = getNonNegativeIntegerInput("Amount: ")
	country = getNonEmptyInput("What currency : ")
	return
}

func calculateAmountOfDollars(amountInOtherCurrency int, exchangeRateToDollars float64) (amountDollars float64) {
	amountDollars = float64(amountInOtherCurrency) * exchangeRateToDollars / 100.0
	return
}

func printOutput(amountInOtherCurrency int, exchangeRateToDollars float64, amountDollars float64) {
	fmt.Printf("%d at an exchange rate of %.2f is %.2f U.S. dollars.\n", amountInOtherCurrency, exchangeRateToDollars, amountDollars)
}

func main() {
	ratesMap, err := readRatesFromOpenExchange()
	if err != nil {
		fmt.Println("Could not read from open exchange! Error: ", err)
		os.Exit(1)
	}


	amount, currency := getInput()

	if ratesMap[currency] != nil {
		exchangeRate := ratesMap[currency].(float64)
		amountDollars := calculateAmountOfDollars(amount, exchangeRate)
		printOutput(amount, exchangeRate, amountDollars)
	} else {
		fmt.Printf("Does not have an exchange rate for currency %s.\n", currency)
	}
}