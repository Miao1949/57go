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

type errorMessage struct {
	Heading string
	ErrorMessage string
}

//---------------------------------------------------------
// CONSTANTS
//---------------------------------------------------------
const hostName = "localhost"
const port = 8000
var serverAddress = hostName + ":" + strconv.Itoa(port)

const saveTextSnippetPath = "saveTextSnippet"
const displayTextSnippetPath = "displayTextSnippet"
const displayEditTextSnippetPath = "displayEditTextSnippet"
const editTextSnippetPath = "editTextSnippet"

var displayTextSnippetRegRxp = regexp.MustCompile(`/` + displayTextSnippetPath + `/(\w+)$`)
var displayEditTextSnippetRegRxp = regexp.MustCompile(`/` + displayEditTextSnippetPath + `/(\w+)$`)
var editTextSnippetRegRxp = regexp.MustCompile(`/` + editTextSnippetPath + `/(\w+)$`)

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

const displayEditTextSnippetPageHtmlContents = `<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
 <meta http-equiv="Content-Type" content="text/html;charset=UTF-8"> 
</head>
<body>
<h1>Text snippet saver</h1>
<form method=post action="http://{{.Host}}:{{.Port}}/{{.Path}}">
<table>
<tr>
<th align=right>Text to edit:</th>
<td><input type=text name="text" value="{{.Text}}" size=32 /></td>
</tr>
<tr>
<td><input type=submit value="Edit" /></td>
</tr>
</table>
</form>
</body>
</html>`
var displayEditTextSnippetPageTemplate = template.Must(template.New("displayTextSnippet").Parse(displayEditTextSnippetPageHtmlContents))

const textSnippetHasBeenEditedPageHtmlContents = `<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
 <meta http-equiv="Content-Type" content="text/html;charset=UTF-8"> 
</head>
<body>
<h1>Text snippet saver</h1>
 The text has been updated. It can be retrieved at this url:  <a href="http://{{.Host}}:{{.Port}}/{{.Path}}">{{.Host}}:{{.Port}}/{{.Path}}</a>
</body>
</html>
`
var textSnippetHasBeenEditedPageTemplate = template.Must(template.New("textSnippetHasBeenEdited").Parse(textSnippetHasBeenEditedPageHtmlContents))

const errorPageHtmlContents = `<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
 <meta http-equiv="Content-Type" content="text/html;charset=UTF-8"> 
</head>
<body>
<h1>{{.Heading}}</h1>
 {{.ErrorMessage}}
</body>
</html>`
var errorPageTemplate = template.Must(template.New("errorPageTemplate").Parse(errorPageHtmlContents))


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
	//http.HandleFunc("/" + editTextSnippetPath, handleEditTextSnippet)
	log.Fatal(http.ListenAndServe(serverAddress, nil))
}


func handleDefaultRequest(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("Path is: %s\n", request.URL.Path)
	requestMatched := false
	if len(request.URL.Path) == 1 && request.URL.Path[0] == '/' {
		//---------------------------------------------------------
		// ROOT
		//---------------------------------------------------------
		handleRootRequest(writer, request)
		requestMatched = true
	} else if len(request.URL.Path) > 1 && displayTextSnippetRegRxp.MatchString(request.URL.Path)  {
		//---------------------------------------------------------
		// DISPLAY SNIPPET.
		//---------------------------------------------------------
		fmt.Println("Path matches display text snippet!")
		groups := displayTextSnippetRegRxp.FindAllStringSubmatch(request.URL.Path, -1)
		if len(groups) > 0 && len(groups[0]) > 1 {
			textSnippetUrl := strings.TrimSpace(groups[0][1])
			handleDisplaySnippetRequest(writer, request, textSnippetUrl)
			requestMatched = true
		}
	} else if len(request.URL.Path) > 1 && displayEditTextSnippetRegRxp.MatchString(request.URL.Path)  {
		//---------------------------------------------------------
		// DISPLAY EDIT SNIPPET.
		//---------------------------------------------------------
		fmt.Println("Path matches display edit text snippet!")
		groups := displayEditTextSnippetRegRxp.FindAllStringSubmatch(request.URL.Path, -1)
		if len(groups) > 0 && len(groups[0]) > 1 {
			textSnippetUrl := strings.TrimSpace(groups[0][1])
			handleDisplayEditTextSnippetRequest(writer, request, textSnippetUrl)
			requestMatched = true
		}
	} else if len(request.URL.Path) > 1 && editTextSnippetRegRxp.MatchString(request.URL.Path)  {
		//---------------------------------------------------------
		// EDIT SNIPPET.
		//---------------------------------------------------------
		fmt.Println("Path matches edit text snippet!")
		groups := editTextSnippetRegRxp.FindAllStringSubmatch(request.URL.Path, -1)
		if len(groups) > 0 && len(groups[0]) > 1 {
			textSnippetUrl := strings.TrimSpace(groups[0][1])
			handleEditTextSnippetRequest(writer, request, textSnippetUrl)
			requestMatched = true
		}
	}

	if !requestMatched {
		returnErrorPage(writer, http.StatusNotFound)
	}
}

func handleDisplaySnippetRequest(writer http.ResponseWriter, request *http.Request, textSnippetUrl string) {
	if request.Method != http.MethodGet {
		returnErrorPage(writer, http.StatusMethodNotAllowed)
		return
	}

	text, err := urlHandler.GetTextForUrl(textSnippetUrl)

	if err != nil {
		returnErrorPage(writer, http.StatusNotFound)
		fmt.Fprintf(os.Stderr, "Could not find text for url: %s! Err: %v\n", textSnippetUrl, err)
		return
	}

	pathToEditPage := displayEditTextSnippetPath + "/" + textSnippetUrl
	if err := displayTextSnippetResponseTemplate.Execute(writer, textAndLocation{Text: text, Host: hostName, Port: port, Path: pathToEditPage}); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(os.Stderr, "Could not execute template! Err: %v\n", err)
		return
	}

}

func handleDisplayEditTextSnippetRequest(writer http.ResponseWriter, request *http.Request, textSnippetUrl string) {
	if request.Method != http.MethodGet {
		returnErrorPage(writer, http.StatusMethodNotAllowed)
		return
	}

	text, err := urlHandler.GetTextForUrl(textSnippetUrl)

	if err != nil {
		returnErrorPage(writer, http.StatusNotFound)
		fmt.Fprintf(os.Stderr, "Could not find text for url: %s! Err: %v\n", textSnippetUrl, err)
		return
	}

	pathToEditThisText := editTextSnippetPath + "/" + textSnippetUrl
	if err := displayEditTextSnippetPageTemplate.Execute(writer, textAndLocation{Text: text, Host: hostName, Port: port, Path: pathToEditThisText}); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(os.Stderr, "Could not execute template! Err: %v\n", err)
		return
	}
}

func handleEditTextSnippetRequest(writer http.ResponseWriter, request *http.Request, textSnippetUrl string) {
	if request.Method != http.MethodPost {
		returnErrorPage(writer, http.StatusMethodNotAllowed)
		return
	}

	textToStore := extractTextToStoreFromRequest(request)

	if len(textToStore) == 0 {
		returnErrorPage(writer, http.StatusBadRequest)
		return
	}

	err := urlHandler.SetTextForUrl(textSnippetUrl, textToStore)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not execute store edited text!! Err: %v\n", err)
		returnErrorPage(writer, http.StatusInternalServerError)
		return
	}

	pathToDisplayTextSnippet := displayTextSnippetPath + "/" + textSnippetUrl
	if err := textSnippetHasBeenEditedPageTemplate.Execute(writer, location{Host: hostName, Port: port, Path: pathToDisplayTextSnippet}); err != nil {
		returnErrorPage(writer, http.StatusInternalServerError)
		fmt.Fprintf(os.Stderr, "Could not execute template! Err: %v\n", err)
		return
	}

}

func returnErrorPage(writer http.ResponseWriter, statusCode int) {
	var heading string
	var errorMsg string

	switch statusCode {
	case 400: heading = "400 Bad request"; errorMsg = "That was a bad request!"
	case 404: heading = "404 Not Found"; errorMsg = "Could not find that page :("
	case http.StatusMethodNotAllowed: heading = "405"; errorMsg = "Method not allowed"
	case http.StatusInternalServerError: heading = "500"; errorMsg = "Something went horribly wrong!"
	default:
		heading = "500"; errorMsg = "Something went even more horribly wrong!"
	}

	if err := errorPageTemplate.Execute(writer, errorMessage{Heading: heading, ErrorMessage: errorMsg}); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(os.Stderr, "Could not execute template! Err: %v\n", err)
		return
	}
	writer.WriteHeader(statusCode)
}



func handleRootRequest(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		returnErrorPage(writer, http.StatusMethodNotAllowed)
		return
	}

	if err := rootResponseTemplate.Execute(writer, location{Host: hostName, Port: port, Path: saveTextSnippetPath}); err != nil {
		writer.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(os.Stderr, "Could not execute template! Err: %v\n", err)
		return
	}
}

func handleSaveTextSnippet(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	textToStore := extractTextToStoreFromRequest(request)

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
		fmt.Fprintf(os.Stderr, "Could not execute template! Err: %v\n", err)
		return
	}
}

func extractTextToStoreFromRequest(request *http.Request) string {
	var textToStore string
	request.ParseForm()
	for formKey, formValue := range request.Form {
		fmt.Printf("Key: %s, value: %s type: %s\n", formKey, formValue, reflect.TypeOf(formValue))
		if formKey == "text" && len(formValue) == 1 && len(formValue[0]) > 0 {
			textToStore = formValue[0]
			break
		}
	}
	return textToStore
}
