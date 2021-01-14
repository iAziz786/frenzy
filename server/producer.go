package server

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net"

	"github.com/iAziz786/frenzy/client"
	"github.com/iAziz786/frenzy/constant"
)

// Produce will get the client message produce to the consumers
func Produce(conn net.Conn) {
	for {
		var buf bytes.Buffer

		_, err := io.Copy(&buf, conn)

		if err != nil {
			if err == io.EOF {
				return
			}
			log.Printf("unable to read all %s\n", err)
			continue
		}

		trimmedBytes := bytes.Trim(buf.Bytes(), string(constant.MsgEndIdent))

		msgJSONBytes := bytes.Split(trimmedBytes, []byte{constant.MsgEndIdent})

		for _, msgBytes := range msgJSONBytes {
			if len(msgBytes) <= 0 {
				continue
			}

			var m client.Message
			err = json.Unmarshal(msgBytes, &m)

			if err != nil {
				log.Printf("unable to unmarshal %s %s\n", msgBytes, err)
				continue
			}

			file, err := CreateLogFile(m.Topic)

			if err != nil {
				panic(err)
			}

			if _, err := file.Write(append(msgBytes, byte(constant.MsgEndIdent))); err != nil {
				log.Printf("unable to write event %s", err)
				if err := file.Close(); err != nil {
					log.Fatalf("unable to close the file")
				}
				continue
			}

		}
	}
}
