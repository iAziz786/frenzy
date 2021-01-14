package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net"

	"github.com/iAziz786/frenzy/constant"
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
	// subscribeTopics should have been array of string but for MVP it is just
	// a string, which is the topic the consumer subscribe to
	subscribedTopics string
}

// Read reads the income message prodcue by the producer if the there is any
// error the message will be nil and error value will contain some value
func (c Consumer) Read() ([]Message, error) {
	// TODO: max 1kb read is supported, in future that will be
	// configurable
	b := make([]byte, 1e6)

	n, err := c.conn.Read(b)

	if err != nil {
		log.Printf("unable to read message %s", err)
		return nil, err
	}

	b = b[:n]

	msgJSONBytes := bytes.Split(
		bytes.Trim(b, string(constant.MsgEndIdent)),
		[]byte{constant.MsgEndIdent},
	)

	messages := []Message{}

	for _, msgBytes := range msgJSONBytes {
		if len(msgBytes) < 0 {
			continue
		}
		var m Message
		err = json.Unmarshal(msgBytes, &m)

		if err != nil {
			log.Printf("error while unmarshal JSON %s %s", string(msgBytes), err)
			continue
		}

		messages = append(messages, m)
	}

	return messages, nil
}

// Subscribe will take array of topics and if any event in those topic the
// consumer will be notified about them
func (c Consumer) Subscribe(topics string) error {
	c.subscribedTopics = topics

	msg := Message{
		Topic: c.subscribedTopics,
	}
	b, err := json.Marshal(msg)
	if err != nil {
		return errors.New("error while converting client message")
	}

	if _, err := c.conn.Write(b); err != nil {
		return errors.New("unable to send bootstrap information")
	}

	return nil
}

// Close will close the connection to the broker
func (c Consumer) Close() error {
	return c.conn.Close()
}
