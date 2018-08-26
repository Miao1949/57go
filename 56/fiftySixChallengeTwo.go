package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

//---------------------------------------------------------
// CONSTANTS
//---------------------------------------------------------
const outfilename = "items.csv"

const enterItemChoice = "1"
const printReportInCsvChoice = "2"
const printReportInHtmlChoice = "3"
const printReportInTextChoice = "4"
const searchChoice = "5"
const quitChoice = "q"
const choises = "12345q"

const menuText=`
*************************
*         MENU          *
*************************
1) Enter item.
2) Print report in CSV format.
3) Print report in HTML format.
4) Print report in text format.
5) Search
q) Quit
`
const nameColumnTitle = "Name"
const serialNumberColumnTitle = "SERIAL NUMBER"
const valueColumnTitle= "VALUE"


//---------------------------------------------------------
// TYPES
//---------------------------------------------------------

type item struct {
	Name string
	SerialNumber string
	Value float64
}

//---------------------------------------------------------
// PRIVATE FUNCTIONS.
//---------------------------------------------------------

func main() {
	done := false
	items := readItemsFromFile()
	for !done {
		printMenu()
		choice := getMenuChoiceFromUser()
		items, done = actOnChoice(choice, items)
	}
}

func readItemsFromFile() (items []item) {
	file, e := os.OpenFile(outfilename, os.O_RDONLY, 0644)
	if e != nil {
		fmt.Fprint(os.Stderr, "Could not open file. Error: %v\n", e)
		return
	}
	defer file.Close()

	items = make([]item, 0)
	reader := bufio.NewReader(file)
	done := false
	for !done {
		line, err := reader.ReadString('\n')
		if err == nil || err == io.EOF {
			fields := strings.Split(line, ";")
			if len(fields) == 3 {
				value, e := strconv.ParseFloat(strings.TrimSpace(fields[2]), 64)
				if e == nil {
					items = append(items, item{Name: fields[0], SerialNumber: fields[1], Value: value})
				}
			}
		}

		if err != nil {
			done = true
		}
	}

	return
}

func writeItemToFile(itemToWrite item) {
	if strings.Contains(itemToWrite.Name, ";") || strings.Contains(itemToWrite.Name, ";")  {
		fmt.Fprintln(os.Stderr, "Cannot write item to file as it contains a semicolon!") // Enhancement: substitute ; in string.
		return
	}

	file, e := os.OpenFile(outfilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if e != nil {
		fmt.Fprint(os.Stderr, "Could not open file. Error: %v\n", e)
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	fmt.Fprintf(writer, "%s;%s;%.2f\n", itemToWrite.Name, itemToWrite.SerialNumber, itemToWrite.Value)
	writer.Flush()
	return
}

func printMenu() {
	fmt.Print(menuText)
}

func getMenuChoiceFromUser() (choice string) {
	reader := bufio.NewReader(os.Stdin)
	done := false
	for !done {
		fmt.Print("Choice: ")
		line, err := reader.ReadString('\n')
		if err == nil {
			line = strings.TrimSpace(line)
			if len(line) == 1 && strings.Contains(choises, line) {
				choice = line
				done = true
			}
		}
	}
	return
}

func actOnChoice(choice string, inItems []item) (outItems []item, quit bool){
	outItems = inItems

	switch choice {
	case enterItemChoice: item := letUserEnterItem(); writeItemToFile(item); outItems = append(inItems, item)
	case printReportInCsvChoice: printInCsvFormat(inItems)
	case printReportInHtmlChoice: printInHtmlFormat(inItems)
	case printReportInTextChoice: printInTextFormat(inItems)
	case searchChoice: filterAndPrint(inItems)
	case quitChoice: quit = true
	default:
		panic("Unsupported choice!")
	}

	return
}

func letUserEnterItem() (outItem item){
	name := getStringFromUser("Please enter item name: ")
	serialNumber := getStringFromUser("Please enter serial number: ")
	value := getFloatFromUser("Please enter (numeric) value: ")

	outItem = item{Name: name, SerialNumber: serialNumber, Value: value}
	return
}

func getStringFromUser(msg string) (enteredString string) {
	done := false
	reader := bufio.NewReader(os.Stdin)
	for !done {
		fmt.Print(msg)
		line, err := reader.ReadString('\n')
		if err == nil {
			line = strings.TrimSpace(line)
			if len(line) > 0 {
				enteredString = line
				done = true
			}
		}
	}

	return
}

func getFloatFromUser(msg string) (enteredFloat float64) {
	done := false
	reader := bufio.NewReader(os.Stdin)
	for !done {
		fmt.Print(msg)
		line, err := reader.ReadString('\n')
		if err == nil {
			line = strings.TrimSpace(line)
			if len(line) > 0 {
				f, err := strconv.ParseFloat(line, 64)
				if err == nil {
					enteredFloat = f
					done = true
				}
			}
		}
	}

	return
}

func filterAndPrint(items []item) {
	textToSearchFor := getStringFromUser("Please enter text to search for: ")
	matchingItems := make([]item, 0)
	for _, item := range items {
		if strings.Contains(item.Name, textToSearchFor) ||
			strings.Contains(item.SerialNumber, textToSearchFor) ||
			strings.Contains(fmt.Sprintf("%.2f", item.Value), textToSearchFor) {
				matchingItems = append(matchingItems, item)
		}
	}

	printInTextFormat(matchingItems)
}


func printInCsvFormat(items []item) {
	for _, item := range items {
		fmt.Printf("%s,%s,%.2f\n", item.Name, item.SerialNumber, item.Value)
	}
}

func printInHtmlFormat(items []item) {
	fmt.Println("<html>")
	fmt.Println("<body>")
	fmt.Println("<table>")
	fmt.Printf("<tr><th>%s</th><th>%s</th><th>%s</th></tr>\n", nameColumnTitle, serialNumberColumnTitle, valueColumnTitle)
	for _, item := range items {
		fmt.Printf("<tr><td>%s</td><td>%s</td><td>%.2f</td></tr>\n", item.Name, item.SerialNumber, item.Value)
	}
	fmt.Println("</table>")
	fmt.Println("</body>")
	fmt.Println("</html>")

}
func printInTextFormat(items []item) {
	nameColumnWidth := max(findLengthOfLongestName(items), len(nameColumnTitle)) + 1
	serialNumberColumnWidth := max(findLengthOfLongestSerialNumber(items), len(serialNumberColumnTitle)) + 1
	valueColumnWidth := max(findLengthOfLongestValue(items), len(valueColumnTitle)) + 1

	fmt.Println(strings.Repeat("-", nameColumnWidth + serialNumberColumnWidth + valueColumnWidth + 3))
	fmt.Print(nameColumnTitle)
	fmt.Print(strings.Repeat(" ", nameColumnWidth - len(nameColumnTitle)))
	fmt.Print("| ")
	fmt.Print(serialNumberColumnTitle)
	fmt.Print(strings.Repeat(" ", serialNumberColumnWidth - len(serialNumberColumnTitle)))
	fmt.Print("| ")
	fmt.Println(valueColumnTitle)
	fmt.Println(strings.Repeat("-", nameColumnWidth + serialNumberColumnWidth + valueColumnWidth + 3))
	for _, item := range items {
		fmt.Print(item.Name)
		fmt.Print(strings.Repeat(" ", nameColumnWidth - len(item.Name)))
		fmt.Print("| ")
		fmt.Print(item.SerialNumber)
		fmt.Print(strings.Repeat(" ", serialNumberColumnWidth - len(item.SerialNumber)))
		fmt.Print("| ")
		fmt.Printf("%.2f\n", item.Value)
	}
}

func max(i1, i2 int) (largestInt int) {
	if i1 > i2 {
		return i1
	}

	return i2
}

func findLengthOfLongestName(items []item) (lengthOfLongestName int) {
	for _, item := range items {
		if len(item.Name) > lengthOfLongestName {
			lengthOfLongestName = len(item.Name)
		}
	}

	return
}

func findLengthOfLongestSerialNumber(items []item) (lengthOfLongestSerialNumber int) {
	for _, item := range items {
		if len(item.SerialNumber) > lengthOfLongestSerialNumber {
			lengthOfLongestSerialNumber = len(item.SerialNumber)
		}
	}

	return
}

func findLengthOfLongestValue(items []item) (lengthOfLongestValue int) {
	for _, item := range items {
		valueAsString := fmt.Sprintf("%.2f", item.Value)

		if len(valueAsString) > lengthOfLongestValue {
			lengthOfLongestValue = len(valueAsString)
		}
	}

	return
}
