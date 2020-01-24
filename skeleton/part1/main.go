// Solution to part 1 of the Whispering Gophers code lab.
// This program reads from standard input and writes JSON-encoded messages to
// standard output. For example, this input line:
//	Hello!
// Produces this output:
//	{"Body":"Hello!"}
//
package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

type Message struct {
	Body string
}

func main() {
	// Create a new bufio.Scanner reading from the standard input.
	s := bufio.NewScanner(os.Stdin)
	// Create a new json.Encoder writing into the standard output.
	enc := json.NewEncoder(os.Stdout)

	for /* Iterate over every line in the scanner */ s.Scan() {
		// Create a new message with the read text.
		message := Message{s.Text()}
		// Encode the message, and check for errors!
		err := enc.Encode(message)
		if err != nil {
			log.Fatal(err)
		}
	}
	// Check for a scan error.
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}
