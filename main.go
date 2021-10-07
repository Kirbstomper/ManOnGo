package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Puts executing goroutine to sleep for 5 seconds to simulate work
// then logs that a message was received
func logMessage() {
	time.Sleep(5 * time.Second)
	log.Println("Message received!")
}

//Starts up a server to receive request and returns
//result of fizzBuzz
func main() {
	http.HandleFunc("/fizzbuzz", func(rw http.ResponseWriter, r *http.Request) {
		logMessage()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(rw, "Error reading request body")
			return
		}
		s := string(body)
		i, err := strconv.ParseInt(s, 0, 64)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(rw, "Error parsing request, please send only an integer")
			return
		}
		fmt.Fprint(rw, FizzBuzz(i))
	})
	log.Println("Starting Server!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
