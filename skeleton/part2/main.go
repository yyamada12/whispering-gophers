// Solution to part 2 of the Whispering Gophers code lab.
//
// This program extends part 1.
//
// It makes a connection the host and port specified by the -dial flag, reads
// lines from standard input and writes JSON-encoded messages to the network
// connection.
//
// You can test this program by installing and running the dump program:
// 	$ go get github.com/campoy/whispering-gophers/util/dump
// 	$ dump -listen=localhost:8000
// And in another terminal session, run this program:
// 	$ part2 -dial=localhost:8000
// Lines typed in the second terminal should appear as JSON objects in the
// first terminal.
//
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"log"
	"net"
	"os"
)

var dialAddr = flag.String("dial", "localhost:8000", "host:port to dial")

type Message struct {
	Body string
}

func main() {
	// TODO: Parse the flags.
	flag.Parse()

	// TODO: Open a new connection using the value of the "dial" flag.
	// TODO: Don't forget to check the error.
	c, err := net.Dial("tcp", *dialAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	s := bufio.NewScanner(os.Stdin)
	// TODO: Create a json.Encoder writing into the connection you created before.
	e := json.NewEncoder(c)
	for s.Scan() {
		m := Message{Body: s.Text()}
		err := e.Encode(m)
		if err != nil {
			log.Fatal(err)
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}
