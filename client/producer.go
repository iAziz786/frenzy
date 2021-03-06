package client

import (
	"encoding/json"
	"errors"
	"log"
	"net"

	"github.com/iAziz786/frenzy/constant"
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
	if msg.Topic == "" {
		return errors.New("topic name is missing")
	}

	b, err := json.Marshal(msg)

	if err != nil {
		log.Printf("unable to marshal the value %s\n", err)
	}

	_, err = p.conn.Write(append(b, byte(constant.MsgEndIdent)))

	if err != nil {
		return err
	}

	return nil
}

// Close will close the connection to the broker
func (p Producer) Close() error {
	return p.conn.Close()
}
