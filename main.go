package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"net"
	"strings"

	"github.com/iAziz786/frenzy/client"
)

func main() {
	l, err := net.Listen("tcp", ":9022")

	if err != nil {
		log.Fatalf("unable to listen the server %s", err)
	}

	defer l.Close()

	msgStream := make(chan client.Message)

	defer close(msgStream)

	for {
		conn, err := l.Accept()

		if err != nil {
			log.Printf("unable to accept connection %s\n", err)
			continue
		}

		go handleConn(conn, msgStream)
	}
}

func handleConn(conn net.Conn, msgStream chan client.Message) {
	defer conn.Close()

	// First send the metadata about the connection, like whether the
	// connection is from a producer or a consumer
	reader := bufio.NewReader(conn)
	b, err := reader.ReadString('\n')

	if err != nil {
		log.Fatalf("unable read all the data %s\n", err)
		return
	}

	// Removing the whitespace which is not necessary
	b = strings.Trim(b, "\n")

	switch b {
	case string(client.ProducerConn):
		// do producer stuff
		defer func() {
			for {
				var m client.Message

				b := make([]byte, 1e2)

				n, err := conn.Read(b)

				if err == io.EOF {
					// connection closed
					log.Println("producer closed the connection")
					return
				}

				if err != nil {
					log.Printf("unable to read %s", err)
					continue
				}

				b = b[:n]

				if err != nil {
					log.Printf("unable to read all %s\n", err)
					continue
				}

				err = json.Unmarshal(b, &m)

				if err != nil {
					log.Printf("unable to unmarshal %s\n", err)
					continue
				}

				msgStream <- m
			}
		}()
		break
	case string(client.ConsumerConn):
		// do consumer stuff
		defer func() {
			for msg := range msgStream {

				b, err := json.Marshal(&msg)

				if err != nil {
					log.Fatalf("unable to marshal %s", err)
				}

				_, err = conn.Write(b)

				if err != nil {
					log.Fatalf("unable to write to consumer %s", err)
				}
			}
		}()
		break
	default:
		log.Println("unknown connection type")
	}
}
