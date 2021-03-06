package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
	"strings"
	"strconv"
)

const Infile = "dataFile.csv"
const First = "First"
const Last = "Last"
const salary = "Salary"

type Employee struct {
	firstName string
	lastName string
	salary int
}

func main() {
	printData(readDataFromFile())
}


func readDataFromFile() (emploees []Employee){
	file, e := os.Open(Infile)
	if e != nil {
		fmt.Fprintf(os.Stderr, "Could not open infile! Error: %v\n", e)
		panic(e)
	}

	defer file.Close()

	emploees = make([]Employee, 0)
	reader := bufio.NewReader(file)
	done := false
	for !done {
		line, e2 := reader.ReadString('\n')
		if e2 == io.EOF {
			done = true
		}

		if e2 == nil || e2 == io.EOF {
			line = strings.TrimSpace(line)
			fields := strings.Split(line, ",")
			if len(fields)>=3 {
				salary, e3 := strconv.Atoi(fields[2])
				if e3 == nil {
					emploees = append(emploees, Employee{fields[1], fields[0], salary})
				}
			} else {
				fmt.Fprint(os.Stderr, "Line is incorrectrly formatted! %s ", line)
			}
		}
	}

	return
}

func printData(employees []Employee) {
	lenOfLongestFirstName, lenOfLongestLastName, lenOfLongestSalary := findLengthOfLongestNames(employees)
	firstNameColumnSize := max(lenOfLongestFirstName, len(First)) + 1
	lastNameColumnSize := max(lenOfLongestLastName, len(Last)) + 1
	salaryColumnSize := max(lenOfLongestSalary, len(salary)) + 1
	printHeader(lastNameColumnSize, firstNameColumnSize, salaryColumnSize)

	for _, employee := range employees {
		fmt.Print(employee.lastName)
		fmt.Print(strings.Repeat(" ", lastNameColumnSize-len(employee.lastName)))
		fmt.Print(employee.firstName)
		fmt.Print(strings.Repeat(" ", firstNameColumnSize-len(employee.firstName)))
		fmt.Println(employee.salary)
	}
}

func printHeader(lastNameColumnSize int, firstNameColumnSize int, salaryColumnSize int) {
	fmt.Print(Last)
	fmt.Print(strings.Repeat(" ", lastNameColumnSize-len(Last)))
	fmt.Print(First)
	fmt.Print(strings.Repeat(" ", firstNameColumnSize-len(First)))
	fmt.Println(salary)
	fmt.Println(strings.Repeat("-", firstNameColumnSize+lastNameColumnSize+salaryColumnSize))
}

func findLengthOfLongestNames(emploees []Employee) (lenOfLongestFirstName int, lenOfLongestLastName int, lenOfLongestSalary int) {
	if len(emploees) == 0 {
		return
	}

	for _, employee := range emploees {
		if len(employee.firstName) > lenOfLongestFirstName {
			lenOfLongestFirstName = len(employee.firstName)
		}

		if len(employee.lastName) > lenOfLongestLastName {
			lenOfLongestLastName = len(employee.lastName)
		}

		if len(string(employee.salary)) > lenOfLongestSalary {
			lenOfLongestSalary = len(string(employee.salary))
		}
	}
	return
}

func max(i1 int, i2 int) (max int) {
	max = i1
	if i2 > i1 {
		max = i2
	}
	return
}