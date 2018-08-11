package main

import (
	"os"
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"time"
)

// CONSTANTS
const NEW = "new"
const SHOW = "show"
const SEARCH = "search"

const TIMESTAMP = "Timestamp"
const NOTE = "Note"
const DATABASE_SERVICE_URL="https://fiftyone-a1bf4.firebaseio.com/notes2.json"
const TIMESTAMP_LAYOUT = "2006-01-02 15:04:05"

// TYPES
type Note struct {
	Note string
	Timestamp string
	Tag string
}

func main()  {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "USAGE: go run fiftyOne.go command [options]\n")
		return
	}
	command := os.Args[1]
	if command != NEW && command != SHOW && command != SEARCH{
		fmt.Fprintf(os.Stderr, "Unkown command %s\n", command)
		return
	} else if (command == NEW || command == SEARCH) && len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Not enough params to %s\n", command)
		return
	}

	if command == NEW {
		note := strings.Join(os.Args[2:], " ")
		addNote(note)
	} else if command == SHOW {
		showNotes()
	} else if command == SEARCH {
		searchString := strings.Join(os.Args[2:], " ")
		searchNotes(searchString)
	}
}

func searchNotes(searchString  string) {
	notes, ok := retrieveNotes()
	if !ok {
		return
	}

	for _, note := range notes {
		if strings.Contains(note.Note, searchString) {
			printNote(note)
		}
	}
}

func showNotes() {
	notes, ok := retrieveNotes()
	if !ok {
		return
	}

	for _, note := range notes {
		printNote(note)
	}
}

func addNote(note string) {
	timestamp := time.Now().Format(TIMESTAMP_LAYOUT)
	noteToAdd := Note{Note: note, Timestamp: timestamp}
	bytes, e := json.Marshal(noteToAdd)
	if e != nil {
		fmt.Fprintf(os.Stderr, "Could not marshal fileContent as json! Err: %v\n", e)
		return
	}

	reader := strings.NewReader(string(bytes))

	fmt.Println(note)
	client := &http.Client{}
	request, e2 := http.NewRequest("POST", DATABASE_SERVICE_URL, reader)
	if e2 != nil {
		fmt.Fprintf(os.Stderr, "Could not create request! Err: %v\n", e2)
		return
	}

	response, e3 := client.Do(request)
	if e3 != nil {
		fmt.Fprintf(os.Stderr, "Error when making request! Err: %v\n", e3)
		return
	}

	defer request.Body.Close()

	contentAsByteArray, e5 := ioutil.ReadAll(response.Body)
	if e5 != nil {
		fmt.Fprintf(os.Stderr, "Could not read contents. Error: %v\n", e5)
		return
	}

	fmt.Println(string(contentAsByteArray))
}

func printNote(note Note) {
	if len(note.Tag) > 0 {
		fmt.Printf("%s %s Tag: %s\n", note.Timestamp, note.Note, note.Tag)
	} else {
		fmt.Printf("%s %s\n", note.Timestamp, note.Note)
	}
}

func retrieveNotes() (notes []Note, ok bool) {
	resp, err := http.Get(DATABASE_SERVICE_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Got the following error when reading from URL %s : %v\n", DATABASE_SERVICE_URL, err)
		return
	}
	defer resp.Body.Close()
	contentAsByteArray, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		fmt.Fprintf(os.Stderr, "Could not read contents. Error: %v\n", err)
		return
	}

	var dat map[string]interface{}
	if err := json.Unmarshal(contentAsByteArray, &dat); err != nil {
		fmt.Fprintf(os.Stderr, "Could not urmarshal fileContent as json!\n")
		return
	}

	notes = make([]Note, 0)
	for _, attributeMap := range dat {
		note := attributeMap.(map[string]interface{})
		notes = append(notes, Note{Timestamp:note[TIMESTAMP].(string), Note:note[NOTE].(string)})
	}

	ok = true
	return
}