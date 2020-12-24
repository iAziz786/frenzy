package client

import (
	"encoding/json"
	"log"
	"net"
)

const (
	// ConsumerConn is a connection type where you receive the data from the
	// broker
	ConsumerConn ConnType = "CONSUMER"
)

// Consumer will listen for the messages produced by an Producer and consumes
// them and you can write your logic based on that
type Consumer struct {
	conn net.Conn
}

// Read reads the income message prodcue by the producer if the there is any
// error the message will be nil and error value will contain some value
func (c Consumer) Read() (*Message, error) {
	var m Message

	b, err := json.Marshal(&m)

	if err != nil {
		log.Printf("unable to marshal message to read %s", err)

		return nil, err
	}

	_, err = c.conn.Read(b)

	if err != nil {
		log.Printf("unable to read message %s", err)

		return nil, err
	}

	return &m, nil
}

// Close will close the connection to the broker
func (c Consumer) Close() error {
	return c.conn.Close()
}
