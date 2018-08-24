package main

import (
	"net/http"
	"log"
	"html/template"
	"fmt"
	"os"
	"strconv"
	"reflect"
	"urlHandler"
)

//---------------------------------------------------------
// TYPES
//---------------------------------------------------------

type location struct {
	Host string // Must be capitalized.
	Port int    // Must be capitalized.
	Path string // Must be capitalized.
}

//---------------------------------------------------------
// CONSTANTS
//---------------------------------------------------------
const hostName = "localhost"
const port = 8000
var serverAddress = hostName + ":" + strconv.Itoa(port)

const saveTextSnippetPath = "saveTextSnippet"
const showTextSnippetPath = "showTextSnippet"
const displayEditTextSnippetPath = "displayEditTextSnippet"
const editTextSnippetPath = "editTextSnippet"

const rootResponseText = `<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
 <meta http-equiv="Content-Type" content="text/html;charset=UTF-8"> 
</head>
<body>
<h1>Text snippet saver</h1>
<form method=post action="http://{{.Host}}:{{.Port}}/{{.Path}}" accept-charset="utf-8">
<table>
<tr>
<th align=right>Text to save:</th>
<td><input type=text name="text" size=32 /></td>
</tr>
<tr>
<td><input type=submit value="Submit" /></td>
</tr>
</table>
</form>
</body>
</html>
`
var rootResponseTemplate = template.Must(template.New("rootRequest").Parse(rootResponseText))

const saveTextSnippetResponseText = `<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
 <meta http-equiv="Content-Type" content="text/html;charset=UTF-8"> 
</head>
<body>
<h1>Text snippet saver</h1>
 The text has been saved. It can be retrieved at this url: <a href="{{.Host}}:{{.Port}}/{{.Path}}">{{.Host}}:{{.Port}}/{{.Path}}</a>
</body>
</html>`
var saveTextSnippetResponseTemplate = template.Must(template.New("saveTextSnippetRequest").Parse(saveTextSnippetResponseText))

func main() {
	startServer()
	//strings := []string{"aa", "a", "aaa", "b", "bbbb", "cc", "c"}
	//urlHandler.SortSliceByLength(strings)
	//fmt.Printf("Sorted slice: %v\n", strings)

	//textSnippets, err := dao.GetAllTextSnippets()
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Error when retrieving text snippets. Err: %v", err)
	//}
	//
	//for _, textSnippet := range textSnippets {
	//	fmt.Printf("URL: %s textSnippet: %s\n", textSnippet.Url, textSnippet.Text)
	//}
	//
	//fmt.Println("--------------------------------------------------------")
	//urls, err := dao.GetAllUrls()
	//for _, url := range urls {
	//	fmt.Printf("URL: %s\n", url)
	//}
	//
	//url := "d"
	//exists, err := dao.UrlExist(url)
	//fmt.Printf("%s exists: %v\n", url, exists)
	//
	//url = "does not exist"
	//exists, err = dao.UrlExist(url)
	//fmt.Printf("%s exists: %v\n", url, exists)
	//
	//fmt.Println("--------------------------------------------------------")
	//url = "d"
	//text, err := dao.GetTextForUrl(url)
	//fmt.Printf("Text for url: %s is %s\n", url, text)
	//dao.SetTextForUrl(url, "A new text. Old one: " + text)
	//text, err = dao.GetTextForUrl(url)
	//fmt.Printf("Text for url: %s is now %s\n", url, text)
	//
	//
	//fmt.Println("--------------------------------------------------------")
	//url = "aaa"
	//text = "This is the text for url aaal"
	//dao.AddTextSnippet(url, text)
	//text, err = dao.GetTextForUrl(url)
	//fmt.Printf("Text for url: %s is %s after the insertion.\n", url, text)

}

func startServer() {
	http.HandleFunc("/", handleRootRequest)
	http.HandleFunc("/" + saveTextSnippetPath, handleSaveTextSnippet)
	log.Fatal(http.ListenAndServe(serverAddress, nil))

}


func handleRootRequest(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if err := rootResponseTemplate.Execute(writer, location{Host: hostName, Port: port, Path: saveTextSnippetPath}); err != nil {
		writer.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(os.Stderr, "Could not execute templace! Err: %v\n", err)
		return
	}
}

func handleSaveTextSnippet(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var textToStore string
	request.ParseForm()
	for formKey, formValue := range request.Form {
		fmt.Printf("Key: %s, value: %s type: %s", formKey, formValue, reflect.TypeOf(formValue))
		if formKey == "text" && len(formValue) == 1 && len(formValue[0]) > 0 {
			textToStore = formValue[0]
			fmt.Println("Going to store: " , textToStore)
			break
		}
	}

	if len(textToStore) == 0 {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	url, err := urlHandler.GenerateAndAddSnippet(textToStore)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not execute store text!! Err: %v\n", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := saveTextSnippetResponseTemplate.Execute(writer, location{Host: hostName, Port: port, Path: url}); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(os.Stderr, "Could not execute templace! Err: %v\n", err)
		return
	}
}