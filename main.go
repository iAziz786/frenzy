package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net"
	"strings"

	"github.com/iAziz786/frenzy/client"
	"github.com/iAziz786/frenzy/constant"
	"github.com/iAziz786/frenzy/server"
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

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	// First send the metadata about the connection, like whether the
	// connection is from a producer or a consumer
	reader := bufio.NewReader(conn)
	b, err := reader.ReadString(constant.MsgEndIdent)

	if err != nil {
		log.Fatalf("unable read all the data %s\n", err)
		return
	}

	// Removing the whitespace which is not necessary
	b = strings.Trim(b, string(constant.MsgEndIdent))

	switch b {
	case string(client.ProducerConn):
		// do producer stuff
		defer server.Produce(conn)
		break
	case string(client.ConsumerConn):
		// do consumer stuff
		msg, err := readClientMessage(conn)

		if err != nil {
			log.Printf("unable to read initial boostrap data %s", err)
			return
		}

		defer server.Consume(conn, msg.Topic)
		break
	default:
		log.Println("unknown connection type")
	}
}

func readClientMessage(reader io.Reader) (*client.Message, error) {
	b := make([]byte, 1e2)
	n, err := reader.Read(b)
	if err != nil {
		return nil, err
	}
	b = b[:n]

	var msg client.Message
	if err := json.Unmarshal(b, &msg); err != nil {
		return nil, errors.New("unable to covert client message")
	}
	return &msg, nil
}
