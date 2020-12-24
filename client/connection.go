package client

import (
	"errors"
	"net"
)

// ConnType holds whether the connection is consumer or producer
type ConnType string

// Connection creates a connection to the Kafka cluster and use that
// connection to give you producer and consumer
type Connection struct {
	brokers string
}

// NewConnection will instantiate the connect to the cluster and returns
// a connection
func NewConnection(brokers string) *Connection {
	return &Connection{
		brokers: brokers,
	}
}

// GetProducer creates a producer which then used to send the events
func (c Connection) GetProducer() (*Producer, error) {
	conn, err := net.Dial("tcp", c.brokers)
	if err != nil {
		return nil, errors.New("unable to dial to the broker")
	}

	conn.Write([]byte(ProducerConn))

	return &Producer{conn: conn}, nil
}

// GetConsumer creates a consumer which then used to read the events
func (c Connection) GetConsumer() (*Consumer, error) {
	conn, err := net.Dial("tcp", c.brokers)

	if err != nil {
		return nil, errors.New("unable to dial to the broker")
	}

	conn.Write([]byte(ConsumerConn))

	return &Consumer{conn: conn}, nil
}
