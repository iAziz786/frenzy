package client

import (
	"errors"
	"fmt"
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

	// TODO: This is way to inform the broker that this connection is
	// producer type but this approach is most likely to change. The
	// broker does not need to know whether the connection type. It should
	// write it's logic by read from or to the connection.
	_, err = conn.Write([]byte(fmt.Sprintln(ProducerConn)))

	if err != nil {
		return nil, err
	}

	return &Producer{conn: conn}, nil
}

// GetConsumer creates a consumer which then used to read the events
func (c Connection) GetConsumer() (*Consumer, error) {
	conn, err := net.Dial("tcp", c.brokers)

	if err != nil {
		return nil, errors.New("unable to dial to the broker")
	}

	// TODO: This is way to inform the broker that this connection is
	// consumer type but this approach is most likely to change. The
	// broker does not need to know whether the connection type. It should
	// write it's logic by read from or to the connection.
	_, err = conn.Write([]byte(fmt.Sprintln(ConsumerConn)))

	if err != nil {
		return nil, err
	}

	return &Consumer{conn: conn}, nil
}
