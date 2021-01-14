package server

import (
	"bufio"
	"io"
	"log"
	"net"
	"time"

	"github.com/iAziz786/frenzy/constant"
)

// Poll preodically reads data from a reader for MaxPollWait or till max buffer
// data. Which ever happen first, it will return all the read data at that point.
func Poll(r io.Reader) ([]byte, error) {
	reader := []byte{}
	afterTimer := time.After(time.Second)
	scanner := bufio.NewReader(r)
loop:
	for {
		select {
		case <-afterTimer:
			break loop
		default:
			if text, err := scanner.ReadBytes(constant.MsgEndIdent); err != io.EOF {
				reader = append(reader, text...)
				if len(reader) >= constant.MaxBufData {
					break loop
				}
			}
		}
	}

	return reader, nil
}

// Consume consumes all the incoming request
func Consume(conn net.Conn, topic string) {
	msgStream := ReadLogFile(topic)
	for msg := range msgStream {
		_, err := conn.Write(msg)

		if err != nil {
			log.Fatalf("unable to write to consumer %s", err)
		}
	}
}
