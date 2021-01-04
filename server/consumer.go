package server

import (
	"log"
	"net"
)

// Consume consumes all the incoming request
func Consume(conn net.Conn, topic string) {
	msgStream := ReadLogFile(topic)
	for msg := range msgStream {

		_, err := conn.Write([]byte(msg + "\n"))

		if err != nil {
			log.Fatalf("unable to write to consumer %s", err)
		}
	}
}
