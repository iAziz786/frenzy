package main

import (
	"fmt"
	"io/ioutil"
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

	// First send the metadata about the connection, like whether the
	// connection is from a producer or a consumer
	b, err := ioutil.ReadAll(conn)

	if err != nil {
		log.Fatal("unable read all the data")
		return
	}

	switch string(b) {
	case string(client.ProducerConn):
		// do producer stuff
		fmt.Println(client.ProducerConn)
		break
	case string(client.ConsumerConn):
		// do consumer stuff
		fmt.Println(client.ConsumerConn)
		break
	default:
		log.Fatalln("unknown connection type")
	}
}
