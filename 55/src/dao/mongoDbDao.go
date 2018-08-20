package dao

import (
	"strconv"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"os"
)

const host = "localhost"
const port = 27017
var url = host + ":" + strconv.Itoa(port)
const databaseName = "exercise55"
const textSnippetsCollectionName = "textSnippets"

//---------------------------------------------------------
// TYPES
//---------------------------------------------------------
type TextSnippet struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
	Url string `bson:"url"`
	Text string `bson:"text"`
}

//---------------------------------------------------------
// PUBLIC FUNCTIONS.
//---------------------------------------------------------

func GetAllTextSnippets() (textSnippets []TextSnippet, err error){
	session := getSession()
	defer session.Close()

	textSnippetCollection := getRefToTextSnippetCollection(session)
	err = textSnippetCollection.Find(bson.M{}).All(&textSnippets)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when fetching text snippets. Error: %v\n", err)
		return
	}

	return
}

func GetAllUrls() (urls []string, err error) {
	session := getSession()
	defer session.Close()

	textSnippets := make([]TextSnippet ,0)
	textSnippetCollection := getRefToTextSnippetCollection(session)
	err = textSnippetCollection.Find(bson.M{}).All(&textSnippets)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when fetching text snippets. Error: %v\n", err)
		return
	}

	for _, textSnippet := range textSnippets {
		urls = append(urls, textSnippet.Url)
	}

	return
}

func UrlExist(url string) (exists bool, err error) {
	session := getSession()
	defer session.Close()

	var textSnippet TextSnippet
	textSnippetCollection := getRefToTextSnippetCollection(session)
	err = textSnippetCollection.Find(bson.M{"url": url}).One(&textSnippet)

	exists = err == nil
	return
}

func GetTextForUrl(url string) (text string, err error) {
	session := getSession()
	defer session.Close()

	var textSnippet TextSnippet
	textSnippetCollection := getRefToTextSnippetCollection(session)
	err = textSnippetCollection.Find(bson.M{"url": url}).One(&textSnippet)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when fetching text snippet. Error: %v\n", err)
		return
	}

	text = textSnippet.Text

	return
}

func SetTextForUrl(url string, text string) (err error) {
	session := getSession()
	defer session.Close()

	var textSnippet TextSnippet
	textSnippetCollection := getRefToTextSnippetCollection(session)
	err = textSnippetCollection.Find(bson.M{"url": url}).One(&textSnippet)

	query := bson.M{"url": url}
	change := bson.M{"$set": bson.M{"text": text}}
	err = textSnippetCollection.Update(query, change)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when setting text to %s for url %s. Error: %v\n", text, url, err)
	}

	return
}

func AddTextSnippet(url string, text string) (err error) {
	session := getSession()
	defer session.Close()

	textSnippetCollection := getRefToTextSnippetCollection(session)
	err = textSnippetCollection.Insert(TextSnippet{Url: url, Text: text})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when inserting text snippet. Url: %s text: %s. Error: %v\n", url, text, err)
	}

	return
}

//---------------------------------------------------------
// PRIVATE FUNCTIONS.
//---------------------------------------------------------

func getRefToTextSnippetCollection(session *mgo.Session) *mgo.Collection {
	return  session.DB(databaseName).C(textSnippetsCollectionName)
}

func getSession() (session *mgo.Session) {
	// Open session to DB.
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	return
}