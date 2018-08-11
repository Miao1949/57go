package main

import (
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
	"firebase.google.com/go"
	"log"
	"golang.org/x/net/context"
	"fmt"
	"os"
	"bufio"
	"io"
	"strings"
	"strconv"
)

const DatabaseUrl = "https://fiftyone-a1bf4.firebaseio.com"
const Collection = "todos"
const FirebaseApiKey = "./fiftyone-a1bf4-firebase-adminsdk-plp8l-d4e881d092.json"

type Todo struct {
	Ordinal int `json:"ordinal"`
	Task string `json:"todo"`
}

func main() {
	client, err := initializeFirebaseClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get ref to firebase client")
		return
	}
	todosCollection := getRefToTodosCollection(client)

	todoTexts := getTodosFromUser()
	todos, ok := readCurrentTodosFromFirebase(todosCollection)
	if !ok {
		fmt.Println("Could not read todos from DB 1")
		return
	}

	highestOrdinal := getHighestOrdinalFromTodos(todos)
	addNewTodosToFirebase(todoTexts, highestOrdinal, todosCollection)

	todos, ok = readCurrentTodosFromFirebase(todosCollection)
	if !ok {
		fmt.Println("Could not read todos from DB 1")
		return
	}

	printTodos(todos)
	if len(todos) > 0 {
		askUserToDeleteATodo(todos, todosCollection)
	}
}

func getTodosFromUser() (todoTexts []string) {
	reader := bufio.NewReader(os.Stdin)
	done := false
	for !done {
		fmt.Print("Please enter TODO: ")
		s, err := reader.ReadString('\n')
		if err == nil || err == io.EOF {
			if len(strings.TrimSpace(s)) > 0 {
				todoTexts = append(todoTexts, strings.TrimSpace(s))
			} else {
				done = true
			}
		}

		if err != nil {
			done = true
		}
	}

	return
}


func initializeFirebaseClient() (client *db.Client, err error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile(FirebaseApiKey)
	config := firebase.Config{DatabaseURL: DatabaseUrl}
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

func getRefToTodosCollection(client *db.Client)  (ref *db.Ref){
	ref = client.NewRef(Collection)
	return ref
}

func readCurrentTodosFromFirebase(todosCollection *db.Ref) (todos []Todo, ok bool){
	queryResults, err := todosCollection.OrderByKey().GetOrdered(context.Background())


	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not read data. Error: %v\n", err)
		return
	}
	for _, result := range queryResults {
		var todo Todo
		if err := result.Unmarshal(&todo); err != nil {
			fmt.Fprintf(os.Stderr, "Could not unmarshal data. Error: %v\n", err)
			return
		}
		todos = append(todos, todo)
	}

	ok = true
	return
}

func deleteTodoFromFirebase(ordinalOfTodoToDelete int, todosCollection *db.Ref) (ok bool){
	queryResults, err := todosCollection.OrderByKey().GetOrdered(context.Background())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not read data. Error: %v\n", err)
		return
	}
	for _, queryNode := range queryResults {
		var todo Todo
		if err := queryNode.Unmarshal(&todo); err != nil {
			fmt.Fprintf(os.Stderr, "Could not unmarshal data Error: %v\n", err)
			continue
		}

		if todo.Ordinal == ordinalOfTodoToDelete {
			key := queryNode.Key()
			todosCollection.Child(key).Delete(context.Background())
			return true
		}
	}

	return false
}


func getHighestOrdinalFromTodos(todos []Todo) (highestOrdinal int)  {
	if len(todos) == 0 {
		return 0
	}

	for _, todo := range todos {
		if todo.Ordinal > highestOrdinal {
			highestOrdinal = todo.Ordinal
		}
	}

	return
}

func addNewTodosToFirebase(todoTexts []string, currentlyHighestOrdinal int, todosCollection *db.Ref) {
	nextOrdinal := currentlyHighestOrdinal + 1
	for _, todoText := range todoTexts {
		addNewTodoToFirebase(Todo{Ordinal:nextOrdinal, Task: todoText}, todosCollection)
		nextOrdinal++
	}
}

func addNewTodoToFirebase(todo Todo, todosCollection *db.Ref) {
	_, e := todosCollection.Push(context.Background(), todo)
	if e != nil {
		fmt.Fprintf(os.Stderr, "Could add todo to firebase! Err: %v\n", e)
		return
	}
}

func printTodos(todos []Todo) {
	fmt.Println("*** TODO: ")
	for _,todo := range todos {
		fmt.Printf("[%d] %s\n", todo.Ordinal, todo.Task)
	}
}

func intInArray(i int, ia []int) bool {
	for _, anInt := range ia {
		if anInt == i {
			return true
		}
	}
	fmt.Printf("%s is not in %v", i, ia)
	return false
}

func askUserToDeleteATodo(todos []Todo, todosCollection *db.Ref) {
	existingOrdinals := extractExistingOrdinals(todos)
	done := false
	reader := bufio.NewReader(os.Stdin)
	for !done {
		fmt.Print("Enter the number of the TODO entry to remove: ")
		s, err := reader.ReadString('\n')
		if err == nil || err == io.EOF {
			i, err := strconv.Atoi(strings.TrimSpace(s))
			if err == nil && intInArray(i, existingOrdinals){
				deleteTodoFromFirebase(i, todosCollection)
				done = true
			}
		}
	}
}

func extractExistingOrdinals(todos []Todo) (ordinals []int) {
	for _, todo := range todos {
		ordinals = append(ordinals, todo.Ordinal)
	}
	return
}

