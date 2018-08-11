package main

import (
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
	"time"
)

///The current time is 15:06:26 UTC January 4 2050.
// // Mon Jan 2 15:04:05 MST 2006
const PayloadLayout = "2006-01-02 15:04:05"
const DisplayLayout = "15:04:05 MST January 2 2006"

type ResponsePayload struct {
	CurrentTime string `json:"currentTime"`
}

func main() {
	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		fmt.Fprint(os.Stderr, "Error when GET:ing request. Err: %v", err)
		return
	}

	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	var r ResponsePayload
	json.Unmarshal(contents, &r)
	currentTime, e := time.Parse(PayloadLayout, r.CurrentTime)
	if e != nil {
		fmt.Fprint(os.Stderr, "Could not parse the time retrieved from the server. Err: %v", err)
		return
	}


	fmt.Printf("The current time is %s.\n", currentTime.Format(DisplayLayout))
}
