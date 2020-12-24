package client

import (
	"encoding/json"
	"log"
	"net"
)

const (
	// ProducerConn is a connection type where you send the data to broker
	ProducerConn ConnType = "PRODUCER"
)

// Producer helps you to producer events which you can send to the cluster
type Producer struct {
	conn net.Conn
}

// Send sends the provided message to the connection
func (p Producer) Send(msg Message) error {
	b, err := json.Marshal(Message{
		Value: []byte("hello"),
	})

	if err != nil {
		log.Printf("unable to marshal the value %s\n", err)
	}

	_, err = p.conn.Write(b)
	if err != nil {
		return err
	}

	return nil
}

// Close will close the connection to the broker
func (p Producer) Close() error {
	return p.conn.Close()
}
