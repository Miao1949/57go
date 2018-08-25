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
	"regexp"
	"strings"
)

//---------------------------------------------------------
// TYPES
//---------------------------------------------------------

type location struct {
	Host string // Must be capitalized.
	Port int    // Must be capitalized.
	Path string // Must be capitalized.
}

type textAndLocation struct {
	Text string
	Host string
	Port int
	Path string
}

//---------------------------------------------------------
// CONSTANTS
//---------------------------------------------------------
const hostName = "localhost"
const port = 8000
var serverAddress = hostName + ":" + strconv.Itoa(port)

const saveTextSnippetPath = "saveTextSnippet"
const displayTextSnippetPath = "showTextSnippet"
const displayEditTextSnippetPath = "displayEditTextSnippet"
const editTextSnippetPath = "editTextSnippet"

var showTextSnippetRegRxp = regexp.MustCompile(`/showTextSnippet/(\w+)$`)

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
 The text has been saved. It can be retrieved at this url: <a href="http://{{.Host}}:{{.Port}}/{{.Path}}">{{.Host}}:{{.Port}}/{{.Path}}</a>
</body>
</html>`
var saveTextSnippetResponseTemplate = template.Must(template.New("saveTextSnippetRequest").Parse(saveTextSnippetResponseText))

const displayTextSnippetResponseText = `<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
 <meta http-equiv="Content-Type" content="text/html;charset=UTF-8"> 
</head>
<body>
<h1>Text snippet saver</h1>
The text is:
<br /> 
<strong>{{.Text}}</strong>
<br /> 
You can edit it by clicking on this <a href="http://{{.Host}}:{{.Port}}/{{.Path}}">link</a>
</body>
</html>`
var displayTextSnippetResponseTemplate = template.Must(template.New("displayTextSnippet").Parse(displayTextSnippetResponseText))


const pageNotFoundResponseText = `<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
 <meta http-equiv="Content-Type" content="text/html;charset=UTF-8"> 
</head>
<body>
<h1>404 Not Found</h1>
 Could not find that page :(
</body>
</html>`
var pageNotFoundResponseTemplate = template.Must(template.New("pageNotFound").Parse(pageNotFoundResponseText))


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
	http.HandleFunc("/", handleDefaultRequest)
	http.HandleFunc("/" + saveTextSnippetPath, handleSaveTextSnippet)
	log.Fatal(http.ListenAndServe(serverAddress, nil))

}


func handleDefaultRequest(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	fmt.Printf("Path is: %s\n", request.URL.Path)
	requestMatched := false
	if len(request.URL.Path) == 1 && request.URL.Path[0] == '/' {
		handleRootRequest(writer)
		requestMatched = true
	} else if len(request.URL.Path) > 1 && showTextSnippetRegRxp.MatchString(request.URL.Path)  {
		fmt.Println("Path mathches show text snippet!")
		groups := showTextSnippetRegRxp.FindAllStringSubmatch(request.URL.Path, -1)
		if len(groups) > 0 && len(groups[0]) > 1 {
			textSnippetUrl := strings.TrimSpace(groups[0][1])
			handleShowSnippetRequest(writer, textSnippetUrl)
			requestMatched = true
		}
	}

	if !requestMatched {
		returnPageNotFoundPage(writer)
	}
}

func handleShowSnippetRequest(writer http.ResponseWriter, textSnippetUrl string) {
	text, err := urlHandler.GetTextForUrl(textSnippetUrl)

	if err != nil {
		returnPageNotFoundPage(writer)
		fmt.Fprintf(os.Stderr, "Could not find text for url: %s! Err: %v\n", textSnippetUrl, err)
		return
	}

	pathToEditThisText := editTextSnippetPath + "/" + textSnippetUrl
	if err := displayTextSnippetResponseTemplate.Execute(writer, textAndLocation{Text: text, Host: hostName, Port: port, Path: pathToEditThisText}); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(os.Stderr, "Could not execute templace! Err: %v\n", err)
		return
	}

}

func returnPageNotFoundPage(writer http.ResponseWriter) {
	if err := pageNotFoundResponseTemplate.Execute(writer, nil); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(os.Stderr, "Could not execute templace! Err: %v\n", err)
		return
	}
	writer.WriteHeader(http.StatusNotFound)
}

func handleRootRequest(writer http.ResponseWriter) {
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
		fmt.Printf("Key: %s, value: %s type: %s\n", formKey, formValue, reflect.TypeOf(formValue))
		if formKey == "text" && len(formValue) == 1 && len(formValue[0]) > 0 {
			textToStore = formValue[0]
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

	pathToDisplayTextSnippet := displayTextSnippetPath + "/" + url
	if err := saveTextSnippetResponseTemplate.Execute(writer, location{Host: hostName, Port: port, Path: pathToDisplayTextSnippet}); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(os.Stderr, "Could not execute templace! Err: %v\n", err)
		return
	}
}