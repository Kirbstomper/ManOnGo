package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// Returns fizz if the number is evenly divisible by 3
// buzz if the number is evenly divisible by 5
// fizzbuzz if the number is evenly divisible by both
func fizzBuzz(i int64) string {
	s := ""
	if i%3 == 0 {
		s = s + "fizz"
	}
	if i%5 == 0 {
		s = s + "buzz"
	}
	return s
}

//Starts up a server to receive request and returns
//result of fizzBuzz
func main() {
	http.HandleFunc("/fizzbuzz", func(rw http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
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
		fmt.Fprint(rw, fizzBuzz(i))
	})
	log.Println("Starting Server!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
