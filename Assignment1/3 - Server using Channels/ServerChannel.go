package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// METHOD 1 - closure
	logs := make(chan string)
	go logLogs(logs)               // putl logger into separate thread
	handleHello := makeHello(logs) // create page (write to http: ) and generate log message to channel

	// METHOD 2
	passer := &DataPasser{logs: make(chan string)} // Create structure and init it with string channel
	go passer.log()                                // start logger in separate thread

	//
	// Init http server
	//
	http.HandleFunc("/1", handleHello)        // method 1
	http.HandleFunc("/2", passer.handleHello) // method 2
	http.ListenAndServe(":9999", nil)         // and run it forever
}

//
// METHOD 1
///
// Idea is  is to use a function which will return another handler function.
// When the function is returned, it will create a closure around the channel.
//
func makeHello(logger chan string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logger <- r.Host                          // put hostaddr into local channel
		io.WriteString(w, "Hello world! Case 1.") // put data to server itself
	}
}

// Function called via go. Used to receive data from channel and print to console
func logLogs(logger chan string) {
	for item := range logger { // get data from local channel
		fmt.Println("1. Item", item) // and put received data to console
	}
}

// METHOD 2
//
// Use a struct which holds the channel as a member and use pointer receiver
// methods to handle the request...

type DataPasser struct {
	logs chan string
}

// DataParser "delegate" - http request handler
// and performs writing to channel
func (p *DataPasser) handleHello(w http.ResponseWriter, r *http.Request) {
	p.logs <- r.URL.String() // In this method we'll put into channel the page name
	io.WriteString(w, "Hello world. Case 2")
}

// "delegate" which performs actual logging (read from channel)
func (p *DataPasser) log() {
	for item := range p.logs {
		fmt.Println("2. Item", item)
	}
}
