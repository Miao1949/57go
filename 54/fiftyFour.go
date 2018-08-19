package main

import (
	"net"
	"fmt"
	"os"
	"regexp"
	"bufio"
	"io"
	"strings"
	"html/template"
	"strconv"
	"dao"
	"urlHandler"
)
const Host = "localhost"
const Port = 9090
var Address = Host + ":" + strconv.Itoa(Port)
const Get = "GET"
const Post = "POST"
const GenerateShortUrlPath = "generateShortUrl"
const ContentLengthHeader = "Content-Length"

type Location struct {
	Host string
	Port int
	Path string
}

type shortUrlStatistics struct {
	ShortUrl string
	LongUrl string
	NumberOfInvocations int
}

var requestRegRxp = regexp.MustCompile(`(\w+) ([\w|/]+) HTTP/1.1`)
var urlToGenerateShortUrlForRegExp = regexp.MustCompile( `url=(.+)`)
var statisticsPageRegRxp = regexp.MustCompile(`/?(\w+)/stats/?`)


var rootResponseText = `<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">
<head />
<body>
<h1>URL shortener</h1>
<form method=post action="http://{{.Host}}:{{.Port}}/{{.Path}}">
<table>
<tr>
<th align=right>URL to shorten:</th>
<td><input type=text name="url" size=32 /></td>
</tr>
<tr>
<td><input type=submit value="Submit" /></td>
</tr>
</table>
</form>
</body>
</html>`

var rootResponseTemplate = template.Must(template.New("rootRequest").Parse(rootResponseText))

var generateShortUrlResponseText =  `<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">
<head><title>Mapping added</title></head>
<body>
 The short URL of  {{.LongUrl}} is: <a href="{{.ShortUrl}}">{{.ShortUrl}}</a>
</body>
</html>`

var generateShortUrlResponseTemplate = template.Must(template.New("generateShortUrlRequest").Parse(generateShortUrlResponseText))

var generateStatisticsPageResponseText =  `<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">
<head />
<body>
<h1>URL Shortener Statistics </h1>
Short URL: {{.ShortUrl}}
<br />
Long URL: {{.LongUrl}}
<br />
Number of invocations: {{.NumberOfInvocations}}
</body>
</html>
`
var generateStatisticsPageResponseTemplate = template.Must(template.New("generateStatisticsPageRequest").Parse(generateStatisticsPageResponseText))


func main() {
	//shortUrls, err := dao.GetAllShortUrls()
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Error when retrieving short URLs: %v\n", err)
	//	os.Exit(1)
	//}
	//
	//for _, shortUrl := range shortUrls {
	//	fmt.Printf("Short URL: %s\n", shortUrl)
	//}
	//
	//const DnLongUrl = "www.dn.se"
	//shortUrl, err := dao.GetShortUrlForLongUrl(DnLongUrl)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Error when retrieving short URL for long URL %v\n", err)
	//	os.Exit(1)
	//}
	//fmt.Printf("Short URL: %s for long URL: %s\n", shortUrl, DnLongUrl)
	//
	////err = dao.AddMapping("i", "www.vah.se")
	////if err != nil {
	////	fmt.Fprintf(os.Stderr, "Error when adding mapping %v\n", err)
	////	os.Exit(1)
	////}
	//
	//
	//const ShortUrl = "i"
	//longUrl, err := dao.GetLongUrlForShortUrl(ShortUrl)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Error when retrieving long URL for short URL %v\n", err)
	//	os.Exit(1)
	//}
	//fmt.Printf("Long URL: %s for short URL: %s\n", longUrl, ShortUrl,)
	//
	//shortUrl = "d"
	//numberOfInvocations, err := dao.GetNumberOfUrlInvocations(shortUrl)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Error when retrieving the number of invocations for short URL %v\n", err)
	//	os.Exit(1)
	//}
	//
	//fmt.Printf("Number of invocations of %s is %d\n",  shortUrl, numberOfInvocations)
	//
	//fmt.Printf("Will invoke inc for short URL %s\n", shortUrl)
	//
	//dao.ShortUrlInvoked(shortUrl)
	//
	//numberOfInvocations, err = dao.GetNumberOfUrlInvocations(shortUrl)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Error when retrieving the number of invocations for short URL %v\n", err)
	//	os.Exit(1)
	//}
	//fmt.Printf("Number of invocations of %s is %d\n",  shortUrl, numberOfInvocations)


	//const ShortUrl = "DoesNotExist"
	//found, err := dao.MappingExists(ShortUrl)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Error when retrieving long URL for short URL:%s found: %b Error: %v\n", ShortUrl, found, err)
	//	os.Exit(1)
	//}
	//if found {
	//	fmt.Printf("Long URL exists for short URL: %s\n", ShortUrl, )
	//} else {
	//	fmt.Printf("Long URL does NOT exist for short URL: %s\n", ShortUrl, )
	//}
	//
	//sortedSlice := urlHandler.SortSliceByLength([]string{"aa", "a", "aaa", "b", "bbbb", "cc", "c"})
	//fmt.Printf("Sorted slice: %v\n", sortedSlice)
	//
	//currentIdentifier := "a"
	//fmt.Printf("Next identifier after %s is %s\n", currentIdentifier, urlHandler.NextIdentifierAfter(currentIdentifier))
	//currentIdentifier = "b"
	//fmt.Printf("Next identifier after %s is %s\n", currentIdentifier, urlHandler.NextIdentifierAfter(currentIdentifier))
	//currentIdentifier = "z"
	//fmt.Printf("Next identifier after %s is %s\n", currentIdentifier, urlHandler.NextIdentifierAfter(currentIdentifier))
	//currentIdentifier = "az"
	//fmt.Printf("Next identifier after %s is %s\n", currentIdentifier, urlHandler.NextIdentifierAfter(currentIdentifier))
	//currentIdentifier = "zz"
	//fmt.Printf("Next identifier after %s is %s\n", currentIdentifier, urlHandler.NextIdentifierAfter(currentIdentifier))
	//currentIdentifier = "zzzzzz"
	//fmt.Printf("Next identifier after %s is %s\n", currentIdentifier, urlHandler.NextIdentifierAfter(currentIdentifier))
	//currentIdentifier = "zaz"
	//fmt.Printf("Next identifier after %s is %s\n", currentIdentifier, urlHandler.NextIdentifierAfter(currentIdentifier))
	//
	//const LongUrl = "www.newsweek.com"
	//shortUrl, err := urlHandler.GenerateAndAdd(LongUrl)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Error when adding  short URL for long URL:%s Error: %v\n", LongUrl, err)
	//	os.Exit(1)
	//}
	//fmt.Printf("*** Short URL: %s for long URL: %s\n", shortUrl, LongUrl)

	//printAllUrlMappings()
	startServer()
}

func startServer() {
	listener, err := net.Listen("tcp", Address)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not start listening! Err: %v", err)
		os.Exit(1)
	}

	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not accept connection! Err: %v", err)
			continue
		}
		go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	defer connection.Close()

	method, path, _, body, ok := parseRequest(connection)
	if !ok {
		fmt.Fprintf(connection, "Bad request\n")
		return
	}

	method = strings.TrimSpace(method)
	path = strings.TrimSpace(path)
	if method == Get {
		handleGet(connection, path)
	} else if method == Post {
		handlePost(connection, path, body)
	} else {
		fmt.Fprintf(connection, "Unsupported Method\n")
	}
}

func parseRequest(connection net.Conn) (method string, path string, headers map[string]string, body string, ok bool) {
	// Read first line (containing, among other things, HTTP method and path)
	reader := bufio.NewReader(connection)
	requestString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not read from connection! Err: %v", err)
		return
	}

	if requestRegRxp.MatchString(requestString) {
		groups := requestRegRxp.FindAllStringSubmatch(requestString, -1)
		if len(groups) > 0 && len(groups[0]) > 2 {
			method = strings.TrimSpace(groups[0][1])
			path = strings.TrimSpace(groups[0][2])
			ok = true
			fmt.Printf("Method: %s, Path: %s\n", method, path)
		}
	}

	// Read headers.
	headers = make(map[string]string)
	for true {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not read line from connection! Err: %v", err)
			return
		}

		// An empty line ends the header section.
		if len(line) == 0 {
			break
		}

		// Split the line into header name and header value.
		fields := strings.Split(line, ":")
		if len(fields) > 1 {
			headers[strings.TrimSpace(fields[0])] = strings.TrimSpace(fields[1])
		}
	}

	// Any body to read?
	if _, okk := headers[ContentLengthHeader]; !okk {
		// No. That's OK
		ok = true
		return
	}

	// Read body.
	contentLength, err := strconv.Atoi(headers[ContentLengthHeader])
	if err == nil && contentLength > 0 {
		bdy := make([]byte, contentLength)

		for i := 0; i < contentLength; i++ {
			b, err := reader.ReadByte()
				if err != nil {
					fmt.Fprintf(os.Stderr, "Could not read bytefrom connection! Err: %v", err)
					return
				}
			bdy[i] = b
		}
		body = string(bdy)
		body = strings.TrimSpace(body)
		fmt.Printf("The body read from connection is: %s\n", body)
	}

	ok = true
	return
}

func handlePost(connection net.Conn, path string, body string) {
	path = removeLeadingSlash(path)

	if path == GenerateShortUrlPath {
		handleGenerateShortUrlRequest(connection, body)
	} else {
		writeResponse(connection, 404, "404 Not Found\n")
	}
}

func handleGenerateShortUrlRequest(connection net.Conn, body string) {
	// Extract URL.
	var urlToGenerateShortUrlFor string
	var ok bool
	if urlToGenerateShortUrlFor, ok = getUrlFromBody(body); !ok {
		writeResponse(connection, 500, "")
		return
	}

	// Generate short code and add it to the persistent storage.
	shortUrl, err := urlHandler.GenerateAndAdd(urlToGenerateShortUrlFor)
	if err != nil {
		writeResponse(connection, 500, "")
		return
	}

	// Generate response.
	writeResponseCodeAndContentType(connection, 200)
	if err := generateShortUrlResponseTemplate.Execute(connection, dao.UrlMapping{ShortUrl: shortUrl, LongUrl: urlToGenerateShortUrlFor}); err != nil {
		fmt.Fprintf(os.Stderr, "Could not execute templace! Err: %v\n", err)
		writeResponse(connection, 500, "")
		return
	}
}

func removeLeadingSlash(path string) string {
	if len(path) > 0 && path[0] == '/' {
		path = path[1:]
	}
	return path
}

func getUrlFromBody(body string) (url string, ok bool) {
	if urlToGenerateShortUrlForRegExp.MatchString(body) {
		groups := urlToGenerateShortUrlForRegExp.FindAllStringSubmatch(body, -1)
		if len(groups) > 0 && len(groups[0]) > 0 {
			url= groups[0][1]
			ok = true
			fmt.Printf("URL: %s\n", url)
		}
	}

	return
}

func handleGet(connection net.Conn, path string) {
	if path == "/" {
		handleGetRoot(connection)
		return
	}

	path = removeLeadingSlash(path)
	if longUrl, err := urlHandler.GetLongUrlForShortUrl(path); err == nil {
		shortUrl := path // To be really clear :)
		urlHandler.ShortUrlInvoked(shortUrl)
		generateRedirectTo(connection, longUrl)
		return
	}

	if statisticsPageRegRxp.MatchString(path) {
		handleGetStatisticsPage(connection, path)
		return
	}

	writeResponse(connection, 404, "Page not found")
}

func handleGetStatisticsPage(connection net.Conn, path string) {
	//---------------------------------------------------------
	// Extract short URL to display statistics for.
	//---------------------------------------------------------
	groups := statisticsPageRegRxp.FindAllStringSubmatch(path, -1)
	if len(groups) == 0 || len(groups[0]) == 0 {
		fmt.Fprintf(os.Stderr, "Could not extract short URL from display statistics page path!\n")
		writeResponse(connection, 500, "Internal server error. Sorry!")
		return
	}

	//---------------------------------------------------------
	// Fetch its corresponding long URL.
	//---------------------------------------------------------
	shortUrl := groups[0][1]
	var longUrl string
	var err error
	if longUrl, err = urlHandler.GetLongUrlForShortUrl(shortUrl); err != nil {
		fmt.Fprintf(os.Stderr, "Could not get long URL for short URL: %s Err: %v\n", shortUrl, err)
		body := fmt.Sprintf("Does not have a long URL for short URL:%s \n", shortUrl)
		writeResponse(connection, 404, body)
		return
	}

	//---------------------------------------------------------
	// Fetch its number of invocations.
	//---------------------------------------------------------
	var numberOfInvocations int
	if numberOfInvocations, err = urlHandler.GetNumberOfUrlInvocations(shortUrl); err != nil {
		fmt.Fprintf(os.Stderr, "Could not get number of invocations for short URL: %s! Err: %v\n", shortUrl, err)
		numberOfInvocations = 0 // No invocations registered.
	}

	//---------------------------------------------------------
	// Display statistics page.
	//---------------------------------------------------------
	writeResponseCodeAndContentType(connection, 200)
	statistics := shortUrlStatistics{ShortUrl: shortUrl, LongUrl: longUrl, NumberOfInvocations: numberOfInvocations}
	if err := generateStatisticsPageResponseTemplate.Execute(connection, statistics); err != nil {
		fmt.Fprintf(os.Stderr, "Could not execute display statistics templace! Err: %v\n", err)
		return
	}
}

func generateRedirectTo(connection net.Conn, longUrl string) (ok bool) {

	response := "HTTP/1.1 301 Moved Permanently"
	response += "\n"
	response += fmt.Sprintf("Location:http://%s:80/", longUrl)
	response += "\n"

	_, err := fmt.Fprintf(connection, response)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not write redirect! Err: ", err)
		return false
	}

	return true
}

func handleGetRoot(connection net.Conn) {
	ok := writeResponse(connection, 200, "")
	if !ok {
		return
	}

	if err := rootResponseTemplate.Execute(connection, Location{Host: Host, Port:Port, Path:GenerateShortUrlPath}); err != nil {
		fmt.Fprintf(os.Stderr, "Could not execute templace! Err: %v\n", err)
		return
	}
}

func writeResponseCodeAndContentType(connection net.Conn, responseCode int) {
	var responseCodeAsString string
	if responseCode == 200 {
		responseCodeAsString = "200 OK"
	} else if responseCode == 400 {
		responseCodeAsString = "400 Bad Request"
	} else {
		responseCodeAsString = "404 Not Found"
	}

	fmt.Fprintf(connection, "HTTP/1.1 %s\n", responseCodeAsString)
	fmt.Fprintf(connection, "Content-Type: text/html\n\n")
}

func writeResponse(connection net.Conn, responseCode int, body string) (ok bool) {
	s := "HTTP/1.1 "

	if responseCode == 200 {
		s += "200 OK"
	} else if responseCode == 400 {
		s += "400 Bad Request"
	} else {
		s += "404 Not Found"
	}

	s += "\n"
	s += "Content-Type: text/html"
	if len(body) > 0 {
		s += "\n"
		s += ContentLengthHeader + ": " + strconv.Itoa(len(body))
	}
	s += "\n"
	s += "\n"

	_, err := fmt.Fprintf(connection, s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not write headers! Err: ", err)
		return false
	}

	// Write body, if any.
	if len(body) > 0 {
		_, err  = fmt.Fprintf(connection, body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not write body! Err: %v", err)
		}
		return false
	}

	return true
}

func printAllUrlMappings() {
	urlMappings, err := dao.GetAllUrlMappings()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not retrieve data from database! Err: %v", err)
		return
	}

	for _, urlMapping := range urlMappings {
		fmt.Printf("Short URL: %s, Long URL: %s\n", urlMapping.ShortUrl, urlMapping.LongUrl)
	}
}

func readBody(connection net.Conn) (body string) {
	reader := bufio.NewReader(connection)
	done := false
	for !done {
		line, err := reader.ReadString('\n')
		if err == nil || err == io.EOF {
			body += line
		}

		if err != nil {
			done = true
		}
	}

	return body
}