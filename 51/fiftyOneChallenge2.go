package main

import (
	"os"
	"fmt"
	"strings"
	"time"
	"firebase.google.com/go"
	"log"
	"google.golang.org/api/option"
	"golang.org/x/net/context"
	"firebase.google.com/go/db"
)

// CONSTANTS
const NEW = "new"
const SHOW = "show"
const SEARCH = "search"

const TAG_OPTION = "-t"

const TIMESTAMP = "Timestamp"
const NOTE = "Note"
const TAG = "Tag"
const DATABASE_SERVICE_URL="https://fiftyone-a1bf4.firebaseio.com/notes2.json"
const DATABASE_URL = "https://fiftyone-a1bf4.firebaseio.com"
const COLLECTION = "notes2"
const TIMESTAMP_LAYOUT = "2006-01-02 15:04:05"

const FIREBASE_SDK_API_KEY = "./fiftyone-a1bf4-firebase-adminsdk-plp8l-d4e881d092.json"

// TYPES
type Note struct {
	Note string
	Timestamp string
	Tag string
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

func getRefToNotesCollection(client *db.Client)  (ref *db.Ref){
	ref = client.NewRef(COLLECTION)
	return ref
}

func main()  {
	databaseClient, err := initializeFirebaseClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not initialize firebase client!\n")
		return
	}

	notesCollection := getRefToNotesCollection(databaseClient)

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "USAGE: go run fiftyOne.go command [options]\n")
		return
	}

	command := os.Args[1]
	params := os.Args[2:]
	if command != NEW && command != SHOW && command != SEARCH{
		fmt.Fprintf(os.Stderr, "Unkown command %s\n", command)
		return
	} else if command == NEW && (len(params) < 1 || (params[0] == TAG_OPTION && len(params) < 3)){
		fmt.Fprintf(os.Stderr, "Not enough params to %s\n", command)
		return
	} else if command == SEARCH && len(params) < 1 {
		fmt.Fprintf(os.Stderr, "Not enough params to %s\n", command)
		return
	}

	if command == NEW {
		tag := ""
		var noteTextToAdd string
		if params[0] == TAG_OPTION {
			tag = params[1]
			noteTextToAdd = strings.Join(params[2:], " ")
		} else {
			noteTextToAdd = strings.Join(params, " ")
		}

		addNote2(noteTextToAdd, tag, notesCollection)
	} else if command == SHOW {
		showNotes(notesCollection)
	} else if command == SEARCH {
		searchString := strings.Join(os.Args[2:], " ")
		searchNotes(searchString, notesCollection)
	}
}

func searchNotes(searchString  string, notesCollection *db.Ref) {
	notes, ok := retrieveNotes2(notesCollection)
	if !ok {
		return
	}

	for _, note := range notes {
		if strings.Contains(note.Note, searchString) || strings.Contains(note.Tag, searchString){
			printNote(note)
		}
	}
}

func retrieveNotes2(notesCollection *db.Ref) (notes []Note, ok bool){
	queryResults, err := notesCollection.OrderByKey().GetOrdered(context.Background())


	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not read data. Error: %v\n", err)
		return
	}
	notes = make([]Note, 0)
	for _, result := range queryResults {
		var note Note
		if err := result.Unmarshal(&note); err != nil {
			fmt.Fprintf(os.Stderr, "Could not unmarshal data. Error: %v\n", err)
			return
		}
		notes = append(notes, note)
	}

	ok = true
	return
}

func showNotes(notesCollection *db.Ref) {
	notes, ok := retrieveNotes2(notesCollection)
	if !ok {
		return
	}

	for _, note := range notes {
		printNote(note)
	}
}

func addNote2(note string, tag string, notesCollection *db.Ref) {
	timestamp := time.Now().Format(TIMESTAMP_LAYOUT)
	noteToAdd := Note{Note: note, Timestamp: timestamp, Tag:tag}
	ref, e := notesCollection.Push(context.Background(), noteToAdd)
	if e != nil {
		fmt.Fprintf(os.Stderr, "Could add note to firebase! Err: %v\n", e)
		return
	}
	fmt.Println(ref)
}

func printNote(note Note) {
	if len(note.Tag) > 0 {
		fmt.Printf("%s %s Tag: %s\n", note.Timestamp, note.Note, note.Tag)
	} else {
		fmt.Printf("%s %s\n", note.Timestamp, note.Note)
	}
}
