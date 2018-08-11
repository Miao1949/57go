package main

import (
	"net/http"
	"fmt"
	"time"
	"encoding/json"
	"os"
)
// // Mon Jan 2 15:04:05 MST 2006
const Layout = "2006-01-02 15:04:05"
type ResponsePayload struct {
	CurrentTime string `json:"currentTime"`
}

func main()  {
	http.HandleFunc("/", handleCall)
	http.ListenAndServe("localhost:8080", nil)
}

func handleCall(writer http.ResponseWriter, request *http.Request) {
	payload := ResponsePayload{CurrentTime: time.Now().Format(Layout)}
	payloadAsBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not marshal payload as json! Err: %v", err)
		return
	}
	header := writer.Header()
	header.Set("Content-Type",  "application/json")
	fmt.Fprintf(writer, "%s", string(payloadAsBytes))
}
