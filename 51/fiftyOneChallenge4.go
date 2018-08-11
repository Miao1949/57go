package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"strings"
	"bufio"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
	"firebase.google.com/go"
	"log"
	"golang.org/x/net/context"
)

const Outfilename = "sortedPersons.txt"
const DATABASE_URL = "https://fiftyone-a1bf4.firebaseio.com"
const COLLECTION = "fortyOneUnsorted"
const FIREBASE_SDK_API_KEY = "./fiftyone-a1bf4-firebase-adminsdk-plp8l-d4e881d092.json"

type Person struct {
	FirstName string
	LastName string
}

func initializeFirebaseClient() (client *db.Client, err error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile(FIREBASE_SDK_API_KEY)
	config := firebase.Config{DatabaseURL:DATABASE_URL}
	app, err := firebase.NewApp(ctx, &config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err = app.Database(ctx)
	if err != nil {
		log.Fatalf("Could not create database client!: %v\n", err)
	}

	return
}

func getRefToPersonsCollection(client *db.Client)  (ref *db.Ref){
	ref = client.NewRef(COLLECTION)
	return ref
}


func main() {
	databaseClient, err := initializeFirebaseClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not initialize firebase client!\n")
		return
	}

	personCollection := getRefToPersonsCollection(databaseClient)

	persons, ok := retrieveNotes2(personCollection)
	if !ok {
		return
	}

	writePersonsToFile(sort(persons))
}

func retrieveNotes2(personCollection *db.Ref) (persons []Person, ok bool){
	queryResults, err := personCollection.OrderByKey().GetOrdered(context.Background())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not read data. Error: %v\n", err)
		return
	}
	persons = make([]Person, 0)
	for _, result := range queryResults {
		var note Person
		if err := result.Unmarshal(&note); err != nil {
			fmt.Fprintf(os.Stderr, "Could not unmarshal data. Error: %v\n", err)
			return
		}
		persons = append(persons, note)
	}

	ok = true
	return
}

func writePersonsToFile(persons []Person) {
	file, e := os.OpenFile(Outfilename, os.O_CREATE|os.O_WRONLY, 0644)
	if e != nil {
		fmt.Fprint(os.Stderr, "Could not open file. Error: %v\n", e)
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	fmt.Fprintf(writer,"Total of %d names\n", len(persons))
	fmt.Fprint(writer, strings.Repeat("-", 16) + "\n")
	for _, person := range persons {
		fmt.Fprintf(writer,"%s, %s\n", person.LastName, person.FirstName)
	}

	writer.Flush()
}

func sort(persons []Person) (sortedPersons []Person) {
	sortedPersons = make([]Person, len(persons))

	for indexOfPersonToSort, person := range persons {
		insertionPoint := -1
		for index, sortedPerson := range sortedPersons {
			if person.LastName < sortedPerson.LastName {
				insertionPoint = index
			}
		}

		if insertionPoint < 0 {
			sortedPersons[indexOfPersonToSort] = person
		} else {
			copy(sortedPersons[insertionPoint + 1:], sortedPersons[insertionPoint:len(sortedPersons) - 1])
			sortedPersons[insertionPoint] = person
		}
	}

	return
}

func printPersons(persons []Person) {
	fmt.Printf("Total of %d names\n", len(persons))
	fmt.Println(strings.Repeat("-", 16))
	for _, person := range persons {
		fmt.Printf("%s, %s\n", person.LastName, person.FirstName)
	}
}

func readDataFromFile() (persons []Person) {
	fileData, err := ioutil.ReadFile("persons.txt")
	if err != nil {
		fmt.Fprint(os.Stderr, "Could not open file. Error: %v\n", err)
		return
	}

	for _, line := range strings.Split(string(fileData), "\n") {
		fields := strings.Split(line, ",")
		if len(fields) >= 2{
			persons = append(persons, Person{strings.TrimSpace(fields[1]), strings.TrimSpace(fields[0])})
		}
	}

	return
}