package server

import (
	"encoding/json"
	"io"
	"log"
	"net"

	"github.com/iAziz786/frenzy/client"
)

// Produce will get the client message produce to the consumers
func Produce(conn net.Conn) {
	for {

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

		var m client.Message
		err = json.Unmarshal(b, &m)

		if err != nil {
			log.Printf("unable to unmarshal %s\n", err)
			continue
		}

		file, err := CreateLogFile(m.Topic)

		if err != nil {
			panic(err)
		}

		defer file.Close()

		if _, err := file.WriteString(string(b) + "\n"); err != nil {
			log.Printf("unable to write event %s", err)
			continue
		}

	}
}
