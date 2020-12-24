package main

import (
	"encoding/json"
	"log"
	"net"

	"github.com/iAziz786/frenzy/client"
)

func main() {
	l, err := net.Listen("tcp", ":9022")

	if err != nil {
		log.Fatalf("unable to listen the server %s", err)
	}

	defer l.Close()

	for {
		conn, err := l.Accept()

		if err != nil {
			log.Printf("unable to accept connection %s\n", err)
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	d := json.NewDecoder(conn)

	var msg client.Message

	err := d.Decode(&msg)

	if err != nil {
		log.Printf("unable to decode %s\n", err)
	}
}
